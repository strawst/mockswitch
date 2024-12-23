package config

type Config struct {
	WorkspaceFile *string `yaml:"workspaceFile" json:"workspaceFile"`
}

type Workspace struct {
	Version *uint8            `yaml:"version" json:"version"`
	Id      *string           `yaml:"id" json:"id"`
	Name    *string           `yaml:"name" json:"name"`
	Listen  *string           `yaml:"listen" json:"listen"`
	Proxies []*WorkspaceProxy `yaml:"proxies" json:"proxies"`
}

type WorkspaceProxy struct {
	Prefix *string `yaml:"prefix" json:"prefix"`
	Target *string `yaml:"target" json:"target"`
}

type Route struct {
	Files map[string]*RouteFile `yaml:"files" json:"files"`
}

type RouteFile struct {
	Endpoints []*RouteEndpoint `yaml:"endpoints" json:"endpoints"`
}

type RouteEndpoint struct {
	Name      *string                           `yaml:"name" json:"name"`
	Method    *string                           `yaml:"method" json:"method"`
	Path      *string                           `yaml:"path" json:"path"`
	Queries   map[string]*RouteEndpointQuery    `yaml:"queries" json:"queries"`
	Bodies    map[string]*RouteEndpointQuery    `yaml:"bodies" json:"bodies"`
	Responses map[string]*RouteEndpointResponse `yaml:"responses" json:"responses"`
}

type RouteEndpointQuery struct {
	Name     *string `yaml:"name" json:"name"`
	Type     *string `yaml:"type" json:"type"`
	Required *bool   `yaml:"required" json:"required"`
	Validate *string `yaml:"validate" json:"validate"`
}

type RouteEndpointResponse struct {
	Name        *string `yaml:"name" json:"name"`
	Description *string `yaml:"description" json:"description"`
	File        *string `yaml:"file" json:"file"`
}

type Toggle struct {
	Mock map[string]*ToggleConfig `yaml:"mock" json:"mock"`
}

type ToggleConfig struct {
	ResponseName *string `yaml:"responseName" json:"responseName"`
}
