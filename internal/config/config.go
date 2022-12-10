package config

import (
	"fmt"
	"go-grpc-study/internal/environment"
	"go-grpc-study/internal/logger"
	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type Profile string

const (
	Local Profile = "local"
	Dev   Profile = "dev"
	Stage Profile = "stage"
	Prod  Profile = "prod"
)

const (
	_envProfile = "APP_PROFILE"
)

func getProfile() Profile {
	profile := strings.ToLower(environment.GetString(_envProfile, string(Local)))
	return Profile(profile)
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"database-name"`
}

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"database"`
}

var _logger = logger.NewSugar("config")

var Module = fx.Module("config",
	fx.Provide(New),
)

func New() *AppConfig {
	fileName := fmt.Sprintf("./profile/application-%s.yaml", getProfile())
	fullPath, _ := filepath.Abs(fileName)

	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		_logger.Panic("failed to open profile config file")
	}

	var appConfig AppConfig
	if err := yaml.Unmarshal(bytes, &appConfig); err != nil {
		_logger.Panicf("failed to unmarshaling profile config file")
	}

	_logger.Infof("activate '%s' profile config", getProfile())
	return &appConfig
}
