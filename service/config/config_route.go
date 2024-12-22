package config

import (
	"gopkg.in/yaml.v3"
	"mockswitch/util/interact"
	"os"
	"path/filepath"
	"strings"
)

func (r *Service) RouteLoad() {
	// * walk through routes
	workspaceDir := filepath.Dir(*r.Config.WorkspaceFile)
	err := filepath.Walk(workspaceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mockendpoint.yml") {
			routeFile := new(RouteFile)
			yml, err := os.ReadFile(path)
			if err != nil {
				interact.Fatal("unable to read route file", err)
			}
			if err := yaml.Unmarshal(yml, routeFile); err != nil {
				interact.Fatal("unable to parse route file", err)
			}
			for _, endpoint := range routeFile.Endpoints {
				for key, query := range endpoint.Queries {
					query.Name = &key
				}
				for key, body := range endpoint.Bodies {
					body.Name = &key
				}
				for key, response := range endpoint.Responses {
					response.Name = &key
				}
			}
			r.Route.Files[path] = routeFile
		}
		return nil
	})
	if err != nil {
		interact.Fatal("unable to load route", err)
	}
}
