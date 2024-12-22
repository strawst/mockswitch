package config

import (
	"encoding/json"
	probing "github.com/prometheus-community/pro-bing"
	"net"
)

type Payload struct {
	WorkspaceFile       *string   `yaml:"workspaceFile"`
	RecentWorkspaceFile []*string `yaml:"recentWorkspaceFile"`
}

type State struct {
	Name    *string        `yaml:"name" json:"name"`
	Types   []*StateType   `yaml:"types" json:"types"`
	Devices []*StateDevice `yaml:"devices" json:"devices"`
}

type StateType struct {
	Name  *string          `yaml:"name" json:"name"`
	Label *string          `yaml:"label" json:"label"`
	Icon  *string          `yaml:"icon" json:"icon"`
	Opens []*StateTypeOpen `yaml:"opens" json:"opens"`
}

type StateTypeOpen struct {
	Label   *string           `yaml:"label" json:"label"`
	Ports   []*int            `yaml:"ports" json:"ports"`
	Command *StateTypeCommand `yaml:"command" json:"command"`
}

type StateTypeCommand struct {
	Windows *string `yaml:"windows" json:"windows"`
	MacOS   *string `yaml:"macos" json:"macos"`
}

type StateDevice struct {
	Name       *string                 `yaml:"name" json:"name"`
	Label      *string                 `yaml:"label" json:"label"`
	Type       *string                 `yaml:"type" json:"type"`
	Host       *string                 `yaml:"host" json:"host"`
	Attributes []*StateDeviceAttribute `yaml:"attributes" json:"attributes"`
	Status     *StateDeviceStatus      `yaml:"-" json:"status"`
}

type StateDeviceAttribute struct {
	Name  *string `yaml:"name" json:"name"`
	Value *string `yaml:"value" json:"value"`
}

type StateDeviceStatus struct {
	IpAddress    *net.IP
	Pinger       *probing.Pinger
	LatestPing   *probing.Statistics
	RecentPings  []*probing.Statistics
	CompiledType *StateType
	Port         map[int]bool
}

func (r *StateDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"ipAddress":    r.IpAddress.String(),
		"latestPing":   r.LatestPing,
		"recentPings":  r.RecentPings,
		"compiledType": r.CompiledType,
		"port":         r.Port,
	})
}
