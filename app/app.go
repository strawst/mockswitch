package app

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

type App struct {
	App         *application.App
	Initialized []func()
}

func New() *App {
	a := &App{
		App: nil,
	}

	return a
}
