package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Value struct {
	Port       string `json:"port"`
	Dbhost     string `json:"dbhost"`
	Dbport     string `json:"dbport"`
	Dbname     string `json:"dbname"`
	Dbuser     string `json:"dbuser"`
	Dbpassword string `json:"dbpassword"`
	Jwtsecret  string `json:"jwtsecret"`
}

type Option struct {
	Development Value `json:"development"`
	Staging     Value `json:"staging"`
	Production  Value `json:"production"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (o Option) Load() Value {
	var (
		option Option
		value  Value
	)
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Println(err.Error())
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&option)
	switch os.Getenv("ENVIRONMENT") {
	case "development":
		value = option.Development
		break
	case "staging":
		value = option.Staging
		break
	case "production":
		value = option.Production
		break
	}
	return value
}
