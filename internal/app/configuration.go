package app

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const configTypeFile = "file"

// LoadConfiguration load configuration based on type and location
func LoadConfiguration(configType, configLocation string) error {
	var loader func(string) error
	switch configType {
	case configTypeFile:
		loader = fromFile
	default:
		return errors.New("unsupported configuration type")
	}
	return loader(configLocation)
}

func fromFile(configLocation string) error {
	rawExt := filepath.Ext(configLocation)
	path := filepath.Dir(configLocation)

	filename := filepath.Base(configLocation)
	filename = strings.TrimSuffix(filename, rawExt)
	ext := strings.TrimPrefix(rawExt, ".")

	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)

	return viper.ReadInConfig()
}
