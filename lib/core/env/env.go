package env

import (
	"io/ioutil"

	"github.com/arapov/soil/lib/core/server"
	"gopkg.in/yaml.v2"
)

// Info contains soil application settings.
type Info struct {
	Server server.Info `yaml:"Server"`
}

// LoadConfig reads the configuration file.
func LoadConfig(configFile string) (*Info, error) {
	config := &Info{}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		return nil, err
	}

	return config, err
}
