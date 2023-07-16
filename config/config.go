package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
)

type Config struct {
	App   `yaml:"app"`
	HTTP  `yaml:"http"`
	Log   `yaml:"logger"`
	Mysql `yaml:"mysql"`
}

// App -.
type App struct {
	Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

// HTTP -.
type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

// Log -.
type Log struct {
	Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
}

// Mysql -.
type Mysql struct {
	Host     string `env-required:"true" yaml:"host"`
	Port     string `env-required:"true" yaml:"port"`
	User     string `env-required:"true" yaml:"user"`
	Password string `env-required:"true" yaml:"password"`
	DbName   string `env-required:"true" yaml:"dbName"`
}

var configTarget string = "dev"

var Module = fx.Options(fx.Provide(New))

func New() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigName(configTarget) // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/")  // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatalf("error read config file\n: %v", err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("error unmarshal config file\n: %v", err)
	}

	return cfg, nil
}
