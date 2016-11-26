package main

import (
	"fmt"

	log "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var (
	cfgPath, cfgFile string
	v                *viper.Viper
	logLevel         int
)

// InitializeConfig initializes a config file from either default or provided configuration settings.
func InitializeConfig() (*viper.Viper, error) {

	// Create a new viper instance
	v = viper.New()

	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		if cfgPath == "" {
			v.AddConfigPath(".")
		} else {
			v.AddConfigPath(cfgPath)
		}
	}

	// Initialize with default settings
	setDefaultSettings()

	// Read the configuration file
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return v, fmt.Errorf("Unable to parse Config file : %v", err)
		} else {
			log.WARN.Printf("Unable to read Config file => Using Defaults")
		}
	}

	return v, nil
}

func setDefaultSettings() {
	log.INFO.Println("Setting default")

	// Set default global settings
	v.SetDefault("goslash", map[string]interface{}{
		"loglevel": 6,
	})

	m := map[string]interface{}{"BindPort": 8080}
	v.SetDefault("goslash", map[string]interface{}{
		"HTTP": m,
	})
	/*
		v.SetDefault("goslash.HTTP.BindAddr", "")
		v.SetDefault("goslash.HTTP.BindPort", 8080)
		v.SetDefault("goslash.HTTP.AdminUrl", "/admin")
		v.SetDefault("goslash.HTTP.CheckUrl", "/200")
		v.SetDefault("goslash.HTTP.DumpUrl", "/dump")
		v.SetDefault("goslash.HTTP.ReloadUrl", "/reload")
		v.SetDefault("goslash.HTTP.SetUrl", "/set")
	*/
	// Set default store settings
	v.SetDefault("store", map[string]interface{}{
		"engine":   "file",
		"location": "goslash.db",
	})

}
