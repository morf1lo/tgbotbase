package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	ENGLISH_LANG = "en"
)

var Localizations map[string]*viper.Viper

func LoadLocalizations() error {
	languages := []string{ENGLISH_LANG}
	Localizations = make(map[string]*viper.Viper)

	for _, lang := range languages {
		v := viper.New()
		v.AddConfigPath("i18n")
		v.SetConfigType("yaml")
		v.SetConfigName(lang)

		if err := v.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to load %s localization: %s", lang, err.Error())
		}

		Localizations[lang] = v
	}

	return nil
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}
