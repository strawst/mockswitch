package config

import (
	"mockswitch/app"
)

type Service struct {
	app       *app.App
	Config    *Config
	Workspace *Workspace
	Route     *Route
	Toggle    *Toggle
}

func Serve(app *app.App) *Service {
	// * construct config
	config := &Service{
		app:       app,
		Config:    nil,
		Workspace: nil,
	}

	app.Initialized = append(app.Initialized, func() {
		config.Read()
		config.app.App.EmitEvent("app")
	})

	return config
}
