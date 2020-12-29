package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// InitializeViper Function initializes viper to read config.yml file and environment variables in the application.
func InitializeViper() {
	// Set the file name of the configuration file
	viper.SetConfigFile("config.yml")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading the config file, %s", err)
	}
}
