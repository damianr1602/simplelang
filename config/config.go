package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	// LogPath config logs destination
	LogPath string
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(filepath.Dir("../config"))
	viper.ReadInConfig()

	// To add a default value :
	viper.SetDefault("LOG_PATH", "../log/simplelang.log")
	//To get from the toml file or env var
	LogPath = viper.GetString("LOG_PATH")
}
