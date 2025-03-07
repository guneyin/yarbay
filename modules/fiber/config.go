package fiber

import (
	"time"
)

type Config struct {
	AppName string        `json:"appName"`
	Port    string        `json:"port"`
	Timeout time.Duration `json:"timeout"`
}

type SwaggerConfig struct {
	HostURL  string
	BasePath string
	FilePath string
	Path     string
	Title    string
}

func NewSwaggerConfig(config ...*SwaggerConfig) *SwaggerConfig {
	cfg := &SwaggerConfig{}

	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.BasePath == "" {
		cfg.BasePath = "/"
	}
	if cfg.FilePath == "" {
		cfg.FilePath = "./docs/swagger.json"
	}
	if cfg.Path == "" {
		cfg.Path = "docs"
	}
	if cfg.Title == "" {
		cfg.Title = "API Documentation"
	}

	return cfg
}
