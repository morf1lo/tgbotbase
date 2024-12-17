package localization

import "github.com/morf1lo/tgbotbase/internal/config"

func GetMessage(lang string, key string) string {
	v, exists := config.Localizations[lang]
	if !exists {
		v = config.Localizations[config.ENGLISH_LANG]
	}

	return v.GetString(key)
}
