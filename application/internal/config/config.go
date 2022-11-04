package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppConfig struct {
		IsDebug       bool   `yaml:"is-debug" env:"IS_DEBUG" env-default:"false"`
		IsDevelopment bool   `yaml:"is-development" env:"IS_DEV" env-default:"false"`
		LogLevel      string `yaml:"log-level" env:"LOG_LEVEL" env-default:"trace"`
		AdminUser     struct {
			Email    string `yaml:"email" env:"ADMIN_EMAIL" env-required:"true"`
			Password string `yaml:"password" env:"ADMIN_PASSWORD" env-required:"true"`
		} `yaml:"admin-user"`
	} `yaml:"app-config"`
	HTTP struct {
		IP   string `yaml:"ip" env:"BIND_IP" env-default:"0.0.0.0"`
		Port string `yaml:"port" env:"PORT" env-default:"8000"`
		CORS struct {
			AllowedMethods     []string `yaml:"allowed-methods" env:"HTTP-CORS-ALLOWED-METHODS"`
			AllowedOrigins     []string `yaml:"allowed-origins" env:"HTTP-CORS-ALLOWED-ORIGINS"`
			AllowCredentials   bool     `yaml:"allow-credentials" env:"HTTP-CORS-ALLOW-CREDENTIALS"`
			AllowedHeaders     []string `yaml:"allowed-headers" env:"HTTP-CORS-ALLOWED-HEADERS"`
			OptionsPassthrough bool     `yaml:"options-passthrough" env:"HTTP-CORS-OPTIONS-PASSTHROUGH"`
			ExposedHeaders     []string `yaml:"exposed-headers" env:"HTTP-CORS-EXPOSED-HEADERS"`
			Debug              bool     `yaml:"debug" env:"HTTP-CORS-DEBUG"`
		} `yaml:"cors"`
	} `yaml:"http"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("initializing config")
		instance = &Config{}

		if err := cleanenv.ReadConfig(CONFIG_PATH, instance); err != nil {
			helpDescription := "Help Config Description"
			help, _ := cleanenv.GetDescription(instance, &helpDescription)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}
