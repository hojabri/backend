package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Config Variable
var Config *viper.Viper

func init() {
	log.Info("Configs are being initialized")
	
	// Set Configuration File Value
	configEnv := strings.ToLower(os.Getenv("SERVER_MODE"))
	if len(configEnv) == 0 {
		configEnv = "dev"
	}
	
	// Initialize Configurator
	Config = viper.New()
	
	// Set Configurator Configuration
	Config.SetConfigName(configEnv)
	Config.AddConfigPath("./configs")
	Config.AddConfigPath(".")
	Config.AddConfigPath("../../configs")
	Config.AddConfigPath("/configs")
	
	Config.SetConfigType("yaml")

	// Set Configurator Environment
	Config.SetEnvPrefix("BACKEND")
	
	// Set Configurator to Auto Bind Configuration Variables to
	// Environment Variables
	Config.AutomaticEnv()
	
	// Set Configurator to Load Configuration File
	configLoadFile()
	
	if os.Getenv("CONFIG_DEBUG") != "" {
		Config.Debug()
	}
}


// ConfigLoadFile Function to Load Configuration from File
func configLoadFile() {
	// Load Configuration File
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatalf("ConfigLoadFile loading and reading error: %v\n", err)
	}
}

