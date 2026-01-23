package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

const (
	dbDriver = "sqlite3"
	dbSource = "database/app.db"
)

type Config struct {
	DbURl    string `yaml:"db_url"`
	DbDriver string `yaml:"db_driver"`
	UserName string `yaml:"user_name"`
}

const configFile = "bgcliconfig.yaml"
const configDir = ".config/grapevine/"

func Init() (*Config, error) {
	cfgPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	// Create the config file if it doesn't exist
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		err = makeConfigPath()
		if err != nil {
			return nil, err
		}
		cfg := Config{DbURl: dbSource, DbDriver: dbDriver}
		if err := cfg.Write(); err != nil {
			return nil, err
		}
	}
	// Read the config file
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	cfg := Config{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) Write() error {
	cfgPath, err := getConfigPath()
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfgPath, data, 0600); err != nil {
		return err
	}
	return nil
}

func getConfigPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cfgPath := homedir + "/" + configDir + configFile
	return cfgPath, nil
}

func makeConfigPath() error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := homedir + "/" + configDir
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(path + configFile)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return nil
}
