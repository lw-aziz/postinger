package localization

import (
	"encoding/json"
	"fmt"
	"postinger/config"
	"postinger/util/logwrapper"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Bundle - to log output on console
var Bundle *i18n.Bundle

// LoadBundle local locales file
func LoadBundle(server config.ServerConfig) *i18n.Bundle {
	// Initialize i18n
	Bundle = i18n.NewBundle(language.Make(server.DefaultLanguage))
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, lang := range server.Languages {
		logwrapper.Logger.Debugln("loading file ", fmt.Sprintf("locales/%v.json", lang))
		Bundle.MustLoadMessageFile(fmt.Sprintf("locales/%v.json", lang))
	}

	return Bundle
}

// GetMessage get message from local file
func GetMessage(lang interface{}, id string, templateData interface{}) string {
	language := config.AppConfig.Server.DefaultLanguage
	if lang != nil {
		language = lang.(string)
	}
	if pos(config.AppConfig.Server.Languages, language) == -1 {
		language = config.AppConfig.Server.DefaultLanguage
	}
	localizer := i18n.NewLocalizer(Bundle, language)

	message, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
		TemplateData: templateData,
	})

	if err != nil || message == "" {
		message = id
	}

	return message
}

func pos(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
