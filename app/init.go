package app

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"time"
)

func (r *App) Init(app *application.App) {
	r.App = app
	go func() {
		time.Sleep(3 * time.Second)
		for _, f := range r.Initialized {
			f()
		}
	}()
}
