package main

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// AppConfig ...
type AppConfig struct {
	HTTPAddress     string            `yaml:"httpAddress"`
	UpstreamServers []string          `yaml:"upstreamServers"`
	IPHeaderName    string            `yaml:"ipHeaderName"`
	CORSOrigins     []string          `yaml:"corsOrigins"`
	ProxyLoader     ProxyLoaderConfig `yaml:"proxyLoader"`
}

// ReadYAMLConfig ...
func ReadYAMLConfig(path string, config interface{}) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()

	return yaml.NewDecoder(r).Decode(config)
}
