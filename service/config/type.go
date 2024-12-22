package config

type Config struct {
	WorkspaceFile *string `yaml:"workspaceFile"`
}

type Workspace struct {
	Version *uint8            `yaml:"version"`
	Id      *string           `yaml:"id"`
	Name    *string           `yaml:"name"`
	Listen  *string           `yaml:"listen"`
	Proxies []*WorkspaceProxy `yaml:"proxies"`
}

type WorkspaceProxy struct {
	Prefix *string `yaml:"prefix"`
	Target *string `yaml:"target"`
}

type Route struct {
	Files map[string]*RouteFile `yaml:"files"`
}

type RouteFile struct {
	Endpoints []*RouteEndpoint `yaml:"endpoints"`
}

type RouteEndpoint struct {
	Name      *string                  `yaml:"name"`
	Method    *string                  `yaml:"method"`
	Queries   []*RouteEndpointQuery    `yaml:"queries"`
	Bodies    []*RouteEndpointQuery    `yaml:"bodies"`
	Responses []*RouteEndpointResponse `yaml:"responses"`
}

type RouteEndpointQuery struct {
	Name     *string `yaml:"name"`
	Type     *string `yaml:"type"`
	Required *bool   `yaml:"required"`
	Validate *string `yaml:"validate"`
}

type RouteEndpointResponse struct {
	Name        *string `yaml:"name"`
	Description *string `yaml:"description"`
	File        *string `yaml:"file"`
}

type Toggle struct {
	Mock map[string]*bool `yaml:"mock"`
}
