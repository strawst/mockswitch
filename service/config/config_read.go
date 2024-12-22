package config

import (
	"fmt"
	"github.com/bsthun/gut"
	"github.com/wailsapp/wails/v3/pkg/application"
	"gopkg.in/yaml.v3"
	"mockswitch/util/interact"
	"os"
	"path/filepath"
)

func (r *Service) Read() {
	// * user's config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		interact.Fatal("unable to get user config directory", err)
		gut.Fatal("unable to get user config directory", err)
	}

	// * path construction
	configPath := filepath.Join(configDir, "Pixcee", "Mockswitch", "default_v1.yml")

	// * check if config file exists
	firsttime := false
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// * prompt to load config first time
		r.ReadPromptDefaultConfig(configPath)
		firsttime = true
	} else if err != nil {
		interact.Fatal("unable check if config file exists", err)
		gut.Fatal("unable check if config file exists", err)
	}

	// * declare struct
	config := new(Config)

	// * read config
	yml, err := os.ReadFile(configPath)
	if err != nil {
		interact.Fatal("unable to read configuration file", err)
	}

	// * parse config
	if err := yaml.Unmarshal(yml, config); err != nil {
		interact.Fatal("unable to parse configuration file", err)
	}

	if !firsttime {
		*config.WorkspaceFile = r.ReadPromptLastConfig(*config.WorkspaceFile, configPath)
	}

	// * validate config
	if err := gut.Validate(config); err != nil {
		interact.Fatal("invalid configuration", err)
	}

	// * read workspace state
	workspace := new(Workspace)
	yml, err = os.ReadFile(*config.WorkspaceFile)
	if err != nil {
		interact.Fatal("unable to read workspace file", err)
	}

	// * parse state
	if err := yaml.Unmarshal(yml, workspace); err != nil {
		interact.Fatal("unable to parse workspace file", err)
	}

	// * assign struct
	r.Config = config
	r.Workspace = workspace
	r.RouteLoad()
	r.ToggleLoad()
}

func (r *Service) ReadPromptDefaultConfig(path string) {
	// * inform user
	dialog := application.QuestionDialog()
	dialog.SetTitle("First-time setup")
	dialog.SetMessage("The workspace is not configured yet. Please create and open Mockswitch workspace file.")
	dialog.AddButtons([]*application.Button{
		{
			Label:     "Cancel",
			IsCancel:  true,
			IsDefault: false,
			Callback: func() {
				r.app.App.Quit()
			},
		},
		{
			Label:     "Open",
			IsCancel:  false,
			IsDefault: true,
			Callback:  nil,
		},
	})
	dialog.Show()
	r.ReadBrowseWorkspace(path)
}

func (r *Service) ReadPromptLastConfig(recent string, path string) string {
	// * inform user
	var choice string
	dialog := application.QuestionDialog()
	dialog.SetTitle("Load latest workspace")
	dialog.SetMessage(fmt.Sprintf("The last workspace located at %s. Do you want to load or browse new one?", recent))
	dialog.AddButtons([]*application.Button{
		{
			Label:     "Load",
			IsCancel:  false,
			IsDefault: true,
			Callback: func() {
				choice = "load"
			},
		},
		{
			Label:     "Browse",
			IsCancel:  false,
			IsDefault: false,
			Callback: func() {
				choice = "browse"
			},
		},
	})
	dialog.Show()

	if choice == "browse" {
		return r.ReadBrowseWorkspace(path)
	}

	return recent
}

func (r *Service) ReadBrowseWorkspace(path string) string {
	// * open file dialog
	fileDialog := application.OpenFileDialog()
	fileDialog.SetTitle("Open Mockswitch workspace file")
	fileDialog.AddFilter("Mockswitch workspace file", "*.yml")
	workspacePath, err := fileDialog.PromptForSingleSelection()
	if err != nil || workspacePath == "" {
		r.app.App.Quit()
	} else {
		// * construct payload
		r.Config = &Config{
			WorkspaceFile: &workspacePath,
		}

		// * encode payload
		out, err := yaml.Marshal(r.Config)
		if err != nil {
			interact.Fatal("unable to parse configuration file", err)
		}

		// * write payload
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			interact.Fatal("unable to create configuration directory", err)
		}
		if err := os.WriteFile(filepath.Join(path), out, 0644); err != nil {
			interact.Fatal("unable to write configuration file", err)
		}
	}

	return workspacePath
}
