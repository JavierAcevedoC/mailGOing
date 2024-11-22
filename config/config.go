package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{
		SMTPHost:     viper.GetString("SMTP_HOST"),
		SMTPPort:     viper.GetInt("SMTP_PORT"),
		SMTPUser:     viper.GetString("SMTP_USER"),
		SMTPPassword: viper.GetString("SMTP_PASSWORD"), // para esto fue necesario usar el password de aplicaciones de google
	}
}
