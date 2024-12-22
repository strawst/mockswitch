package main

import (
	"embed"
	"go.uber.org/fx"
	"mockswitch/app"
	"mockswitch/app/instance"
	"mockswitch/service/config"
	"mockswitch/service/fiber"
	"os"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	fx.New(
		fx.Supply(
			frontend,
		),
		fx.Provide(
			config.Serve,
			fiber.Serve,
			instance.New,
			app.New,
		),
		fx.Invoke(
			invoke,
		),
	).Run()
}

func invoke(instance *instance.App, app *app.App, fiber *fiber.Fiber) {
	_ = fiber
	app.Init(instance.App)
	instance.Run()
	os.Exit(0)
}
