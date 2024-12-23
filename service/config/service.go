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
		Route: &Route{
			Files: make(map[string]*RouteFile),
		},
		Toggle: &Toggle{
			Mock: make(map[string]*ToggleConfig),
		},
	}

	app.Initialized = append(app.Initialized, func() {
		config.Read()
		config.app.App.EmitEvent("app", map[string]any{
			"config":    config.Config,
			"workspace": config.Workspace,
			"route":     config.Route,
			"toggle":    config.Toggle,
		})
	})

	return config
}
