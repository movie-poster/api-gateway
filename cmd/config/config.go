package config

import (
	"api-gateway/internal/utils"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Config *Configuration

type Configuration struct {
	Server  ServerConfiguration
	Domain  DomainConfiguration
	Bugsnag BugsnagConfiguration
}

type DomainConfiguration struct {
	GrpcMovie string
	GrpcUser  string
}

type ServerConfiguration struct {
	Port     string
	TestPort string
	Secret   string
	Mode     string
}

type BugsnagConfiguration struct {
	ApiKey       string
	ReleaseStage string
}

// SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		utils.LoggerZap.Fatal("Error al leer el archivo de configuración", zap.String("Description", err.Error()))
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		utils.LoggerZap.Fatal("No se puede decodificar en estructura", zap.String("Description", err.Error()))
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
