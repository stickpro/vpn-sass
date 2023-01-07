package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

const (
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		HTTP HTTPConfig
		DB   DBConfig
	}

	HTTPConfig struct {
		Host               string        `env:"HTTP_HOST" env-default:"localhost"`
		Port               string        `env:"HTTP_PORT" env-default:"8080"`
		ReadTimeout        time.Duration `env:"HTTP_READ_TIMEOUT"`
		WriteTimeout       time.Duration `env:"HTTP_WRITE_TIMEOUT"`
		MaxHeaderMegabytes int           `env:"HTTP_MAX_HEADER_MEGABYTES"`
	}
	DBConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
	}
)

func Init() (*Config, error) {
	var cfg Config
	populateDefaults(cfg)
	fmt.Println(cfg)
	err := cleanenv.ReadEnv(&cfg.HTTP)
	if err != nil {
		return nil, err
	}

	log.Println("Parsed Configuration")
	return &cfg, nil
}

func populateDefaults(cfg Config) {
	cfg.HTTP.ReadTimeout = defaultHttpRWTimeout
	cfg.HTTP.WriteTimeout = defaultHttpRWTimeout
	cfg.HTTP.MaxHeaderMegabytes = defaultHttpMaxHeaderMegabytes

}
