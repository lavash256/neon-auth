package config

import (
	"io/ioutil"
	"neon-auth/app/interface/persistence"
	"os"
	"path/filepath"

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

	filename, _ := filepath.Abs(path)
	yamlConfig, err := ioutil.ReadFile(filename)
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
