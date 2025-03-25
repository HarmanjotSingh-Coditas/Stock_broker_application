package utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	Database DBConfig `yaml:"database"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.ReadFile("./resources/application.yml")
	if err != nil {
		log.Fatalf("Failed to read the Config file: %v", err)
	}
	err = yaml.Unmarshal(file, &AppConfig)
	if err != nil {
		log.Fatalf("Failed to parse the Config file: %v", err)
	}
	fmt.Println("Config Loaded successfully")
}
