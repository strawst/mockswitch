package config

import (
	"mockswitch/app"
)

type Service struct {
	app     *app.App
	Payload *Payload
	State   *State
}

func Serve(app *app.App) *Service {
	// * construct config
	config := &Service{
		app:     app,
		Payload: nil,
		State:   nil,
	}

	app.Initialized = append(app.Initialized, func() {
		config.Read()
		config.app.App.EmitEvent("app")
	})

	return config
}
