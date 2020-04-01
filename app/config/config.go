package config

import (
	"log"

	"github.com/spf13/viper"
)

type (

	// Config exported to be used globally
	Config struct {
		Server   serverConfig   `json:"server"`
		Database databaseConfig `json:"database"`
	}

	serverConfig struct {
		Port string
		Host string
	}

	databaseConfig struct {
		DBUrl      string
		DBName     string
		DBUser     string
		DBPassword string
	}

	profile struct {
		Development Config `json:"development"`
		Staging     Config `json:"staging"`
		Production  Config `json:"production"`
	}
)

// LoadConfig exported to be used globally
func LoadConfig() (*Config, error) {
	var (
		profile = new(profile)
		config  = new(Config)
	)
	viper.SetConfigName("config")
	viper.AddConfigPath("./app/config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("FAILED::could not read config ERROR")
		return nil, err
	}
	if err := viper.Unmarshal(profile); err != nil {
		log.Println("FAILED::could not read config ERROR")
		return nil, err
	}
	activeProfile := viper.GetString("ENV")
	if len(activeProfile) == 0 {
		activeProfile = "development"
	}
	switch activeProfile {
	case "development":
		config = &profile.Development
		break
	case "staging":
		config = &profile.Staging
		break
	case "production":
		config = &profile.Production
		break
	}
	return config, nil
}
