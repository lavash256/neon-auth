package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Lavash95/neon-auth/internal/interface/persistence"

	"gopkg.in/yaml.v2"
)

//Config is main structure of application config
type Config struct {
	Database persistence.PostgresConfig `yaml:"database"`
	RPC      RPCConfig                  `yaml:"rpc"`
}

//RPCConfig is responsible for setting parameters related to rpc
type RPCConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

//LoadConfig ...
func LoadConfig(path string) (Config, error) {
	var config = Config{}

	filename, err := filepath.Abs(path)
	if err != nil {
		return config, err
	}
	yamlConfig, err := ioutil.ReadFile(filepath.Clean(filename))
	yamlConfig = []byte(os.ExpandEnv(string(yamlConfig)))
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(yamlConfig, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

//Validate checks configuration structures for null values
func (config *Config) Validate() error {
	templateError := "%s is empty.Required field"
	if config.Database.Host == "" {
		return fmt.Errorf(templateError, "Database host")
	}
	if config.Database.Port == "" {
		return fmt.Errorf(templateError, "Database port")
	}
	if config.Database.User == "" {
		return fmt.Errorf(templateError, "Database user")
	}
	if config.Database.Name == "" {
		return fmt.Errorf(templateError, "Database name")
	}
	if config.RPC.Port == "" {
		return fmt.Errorf(templateError, "Service port")
	}
	return nil
}
