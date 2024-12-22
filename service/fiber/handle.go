package fiber

import (
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func (r *Fiber) Handle(c *fiber.Ctx) error {
	path := c.Path()
	method := c.Method()
	key := r.config.ToggleKey(path, method)

	// * mock
	if r.config.Toggle.Mock[key] != nil {
		responseName := r.config.Toggle.Mock[key].ResponseName
		for filepth, file := range r.config.Route.Files {
			filedir := filepath.Dir(filepth)
			filedir = strings.TrimPrefix(filedir, filepath.Dir(*r.config.Config.WorkspaceFile))
			for _, endpoint := range file.Endpoints {
				if r.config.ToggleKey(*endpoint.Path, *endpoint.Method) == key {
					if endpoint.Responses[*responseName] != nil {
						responseFile := filepath.Join(filepath.Dir(*r.config.Config.WorkspaceFile), filedir, *endpoint.Responses[*responseName].File)
						yml, err := os.ReadFile(responseFile)
						if err != nil {
							return gut.Err(false, "unable to read response file", err)
						}

						content := make(map[string]any)
						if err := yaml.Unmarshal(yml, &content); err != nil {
							return gut.Err(false, "unable to parse reponse file", err)
						}

						return c.JSON(content)
					}
				}
			}
		}
		return nil
	}

	// * proxy
	for _, prx := range r.config.Workspace.Proxies {
		if strings.HasPrefix(path, *prx.Prefix) {
			relativePath := strings.TrimPrefix(path, *prx.Prefix)
			targetURL, err := url.JoinPath(*prx.Target, relativePath)
			if err != nil {
				return gut.Err(false, "unable to join target URL", err)
			}
			return proxy.Forward(targetURL)(c)
		}
	}

	return nil
}
