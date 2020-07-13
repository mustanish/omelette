package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	// Config exported to be used globally
	Config struct {
		Server   serverConfig   `yaml:"server"`
		Database databaseConfig `yaml:"database"`
	}

	serverConfig struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}

	databaseConfig struct {
		DBUrl      string `yaml:"dbUrl"`
		DBName     string `yaml:"dbName"`
		DBUser     string `yaml:"dbUser"`
		DBPassword string `yaml:"dbPassword"`
	}

	profile struct {
		Development Config `yaml:"development"`
		Testing     Config `yaml:"testing"`
		Production  Config `yaml:"production"`
	}
)

// LoadConfig exported to be used globally
func LoadConfig() (*Config, error) {
	var (
		profile       = new(profile)
		config        = new(Config)
		activeProfile = os.Getenv("ENV")
		filePath      = "config.yml"
	)
	if activeProfile == "testing" {
		filePath = "../config.yml"
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("FAILED::could not read config ERROR because of", err.Error())
		return nil, err
	}
	content = []byte(os.ExpandEnv(string(content)))
	if err := yaml.Unmarshal(content, profile); err != nil {
		log.Println("FAILED::could not unmarshal config ERROR because of", err.Error())
		return nil, err
	}
	if len(activeProfile) == 0 {
		activeProfile = "development"
	}
	switch activeProfile {
	case "development":
		config = &profile.Development
		break
	case "testing":
		config = &profile.Testing
		break
	case "production":
		config = &profile.Production
		break
	}
	return config, nil
}
