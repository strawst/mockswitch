package instance

import (
	"embed"
	"github.com/bsthun/gut"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
	"mockswitch/service/config"
)

type App struct {
	frontend embed.FS
	config   *config.Service
	App      *application.App
}

func New(frontend embed.FS, config *config.Service) *App {
	// * extract icon
	icon, err := frontend.ReadFile("frontend/dist/appicon.png")
	if err != nil {
		gut.Fatal("unable to read icon", err)
	}

	a := &App{
		frontend: frontend,
		config:   config,
		App:      nil,
	}

	a.App = application.New(application.Options{
		Name:        "Mockswitch",
		Description: "Mockswitch",
		Icon:        icon,
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		Windows: application.WindowsOptions{},
		Linux:   application.LinuxOptions{},
		Services: []application.Service{
			application.NewService(a.config),
		},
		BindAliases: nil,
		Logger:      nil,
		LogLevel:    0,
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(a.frontend),
		},
		Flags:                       nil,
		PanicHandler:                nil,
		DisableDefaultSignalHandler: false,
		KeyBindings:                 nil,
		OnShutdown:                  nil,
		ShouldQuit:                  nil,
		RawMessageHandler:           nil,
		ErrorHandler:                nil,
	})

	return a
}

func (r *App) Run() {
	r.App.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Mockswitch",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		DevToolsEnabled:  true,
	})

	err := r.App.Run()
	if err != nil {
		log.Fatal(err)
	}
}
