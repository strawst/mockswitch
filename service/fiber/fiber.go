package fiber

import (
	"context"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"mockswitch/service/config"
)

type Fiber struct {
	Fiber *fiber.App
}

func Init(lifecycle fx.Lifecycle, config *config.Service) *Fiber {
	i := &Fiber{
		Fiber: nil,
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

	go func() {
		err := i.Fiber.Listen(*config.Workspace.Listen)
		if err != nil {
			gut.Fatal("unable to listen", err)
		}
	}()

	return i
}
