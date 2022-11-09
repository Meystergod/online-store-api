package config

import (
	"context"
	"flag"
	"os"
	"sync"

	"github.com/Meystergod/online-store-api/pkg/logging"

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
		IP   string `yaml:"http_ip" env:"HTTP_IP" env-default:"0.0.0.0"`
		Port string `yaml:"http_port" env:"HTTP_PORT" env-default:"8000"`
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
	GRPC struct {
		IP   string `yaml:"grpc_ip" env:"GRPC_IP" env-default:"0.0.0.0"`
		Port string `yaml:"grpc_port" env:"GRPC_PORT" env-default:"8001"`
	} `yaml:"grpc"`
	PostgreSQL struct {
		Username string `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
		Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
	} `yaml:"postgre-sql"`
}

const (
	EnvCfgPathName  = "CONFIG-PATH"
	FlagCfgPathName = "config"
)

var cfgPath string
var instance *Config
var once sync.Once

func GetConfig(ctx context.Context) *Config {
	once.Do(func() {
		logging.GetLogger(ctx).Info("initializing config")

		flag.StringVar(&cfgPath, FlagCfgPathName, CONFIG_PATH, "this is application config file")
		flag.Parse()

		if cfgPath == "" {
			cfgPath = os.Getenv(EnvCfgPathName)
		}

		if cfgPath == "" {
			logging.GetLogger(ctx).Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpDescription := "help config description"
			help, _ := cleanenv.GetDescription(instance, &helpDescription)
			logging.GetLogger(ctx).Info(help)
			logging.GetLogger(ctx).Fatal(err)
		}
	})

	return instance
}
