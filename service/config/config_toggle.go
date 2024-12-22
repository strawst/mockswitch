package config

import (
	"github.com/bsthun/gut"
	"gopkg.in/yaml.v3"
	"mockswitch/util/interact"
	"os"
	"path/filepath"
	"strings"
)

func (r *Service) ToggleLoad() {
	// * construct toggle map
	r.Toggle.Mock = make(map[string]*bool)

	// * iterate through routes
	for _, routeFile := range r.Route.Files {
		for _, routeEndpoint := range routeFile.Endpoints {
			r.Toggle.Mock[*routeEndpoint.Path] = gut.Ptr(false)
		}
	}

	// * construct toggle path
	configPathDir := filepath.Dir(*r.Config.WorkspaceFile)
	configFileExt := filepath.Ext(*r.Config.WorkspaceFile)
	configFileName := strings.TrimSuffix(filepath.Base(*r.Config.WorkspaceFile), configFileExt)
	configTogglePath := filepath.Join(configPathDir, configFileName+"-toggle"+configFileExt)

	// * check if config file exists
	existingToggleMock := make(map[string]*bool)
	existingBytes, err := os.ReadFile(configTogglePath)
	if err != nil {
		if !os.IsNotExist(err) {
			interact.Error("unable to read toggle file", err)
		}
	} else {
		if err := yaml.Unmarshal(existingBytes, &existingToggleMock); err != nil {
			interact.Error("unable to parse toggle file", err)
		}
	}

	// * merge existing toggle mock
	for k, v := range existingToggleMock {
		r.Toggle.Mock[k] = v
	}

	// * write toggle mock
	toggleBytes, err := yaml.Marshal(r.Toggle.Mock)
	if err != nil {
		interact.Error("unable to marshal toggle mock", err)
	}
	if err := os.WriteFile(configTogglePath, toggleBytes, 0644); err != nil {
		interact.Error("unable to write toggle file", err)
	}
}
