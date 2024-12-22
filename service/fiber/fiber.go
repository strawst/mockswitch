package fiber

import (
	"context"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"mockswitch/app"
	"mockswitch/service/config"
)

type Fiber struct {
	app    *app.App
	config *config.Service
	Fiber  *fiber.App
}

func Serve(lifecycle fx.Lifecycle, app *app.App, config *config.Service) *Fiber {
	i := &Fiber{
		config: config,
		Fiber:  nil,
	}

	i.Fiber = fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: true,
		Network:       "tcp",
	})

	i.Fiber.Use(i.Handle)

	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			// * shutdown
			_ = i.Fiber.Shutdown()
			return nil
		},
	})

	app.Initialized = append(app.Initialized, func() {
		go func() {
			err := i.Fiber.Listen(*config.Workspace.Listen)
			if err != nil {
				gut.Fatal("unable to listen", err)
			}
		}()
	})

	return i
}
