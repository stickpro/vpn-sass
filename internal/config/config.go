package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"strconv"
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
		Host               string `env:"HTTP_HOST" env-default:"localhost"`
		Port               string `env:"HTTP_PORT" env-default:"8080"`
		ReadTimeout        string `env:"HTTP_READ_TIMEOUT"`
		WriteTimeout       string `env:"HTTP_WRITE_TIMEOUT"`
		MaxHeaderMegabytes string `env:"HTTP_MAX_HEADER_MEGABYTES"`
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
	populateDefaults()
	err := cleanenv.ReadEnv(&cfg.HTTP)
	if err != nil {
		return nil, err
	}
	log.Println("Parsed Configuration")
	return &cfg, nil
}

func populateDefaults() {
	err := os.Setenv("HTTP_READ_TIMEOUT", strconv.FormatInt(int64(defaultHttpRWTimeout), 10))
	if err != nil {
		return
	}

	err = os.Setenv("HTTP_WRITE_TIMEOUT", strconv.FormatInt(int64(defaultHttpRWTimeout), 10))
	if err != nil {
		return
	}

	err = os.Setenv("HTTP_MAX_HEADER_MEGABYTES", string(defaultHttpMaxHeaderMegabytes))
	if err != nil {
		return
	}
}
