package config

import (
	"gopkg.in/yaml.v3"
	"mockswitch/util/interact"
	"os"
	"path/filepath"
	"strings"
)

func (r *Service) ToggleLoad() {
	// * construct toggle map
	r.Toggle.Mock = make(map[string]*ToggleConfig)

	// * iterate through routes
	for _, routeFile := range r.Route.Files {
		for _, routeEndpoint := range routeFile.Endpoints {
			key := r.ToggleKey(*routeEndpoint.Path, *routeEndpoint.Method)
			var responseName *string
			for _, response := range routeEndpoint.Responses {
				responseName = response.Name
			}
			r.Toggle.Mock[key] = &ToggleConfig{
				ResponseName: responseName,
			}
		}
	}

	// * construct toggle path
	configPathDir := filepath.Dir(*r.Config.WorkspaceFile)
	configFileExt := filepath.Ext(*r.Config.WorkspaceFile)
	configFileName := strings.TrimSuffix(filepath.Base(*r.Config.WorkspaceFile), configFileExt)
	configTogglePath := filepath.Join(configPathDir, configFileName+"-toggle"+configFileExt)

	// * check if config file exists
	existingToggleMock := make(map[string]*ToggleConfig)
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
		if _, ok := r.Toggle.Mock[k]; ok {
			// only override if key exists
			r.Toggle.Mock[k] = v
		}
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

func (r *Service) ToggleKey(path string, method string) string {
	return method + " " + path
}
