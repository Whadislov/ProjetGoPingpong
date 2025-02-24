package myapp

import (
	"embed"
	"encoding/json"

	"fyne.io/fyne/v2"

	"github.com/jeandeaual/go-locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func setLanguage(lang string) {
	localizer = i18n.NewLocalizer(bundle, lang)
}

// T sets the right language for the user
func T(messageID string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
}

// Initialize translations
func InitTranslations(t embed.FS) {
	translations = t

	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	translated = []language.Tag{language.Make("en")} // the first item in this list will be the fallback if none match
	err := AddTranslationsFS(translations, "translation")
	if err != nil {
		fyne.LogError("Error occurred loading built-in translations", err)
	}

	setDefaultLanguage()
}

// Initialize translations
func setDefaultLanguage() {
	all, err := locale.GetLocales()
	if err != nil {
		fyne.LogError("Failed to load user locales", err)
		// Default language is system language. If not available, English is the default language
		all = []string{"en"}
	}
	str := closestSupportedLocale(all).LanguageString()

	switch str {
	case "fr-FR":
		currentSelectedLanguage = "Français"
		setLanguage("fr")
		return
	case "de-DE":
		currentSelectedLanguage = "Deutsch"
		setLanguage("de")
		return
	default:
		currentSelectedLanguage = "English"
		setLanguage("en")
		return
	}
}

func closestSupportedLocale(locs []string) fyne.Locale {
	matcher := language.NewMatcher(translated)

	tags := make([]language.Tag, len(locs))
	for i, loc := range locs {
		tag, err := language.Parse(loc)
		if err != nil {
			fyne.LogError("Error parsing user locale", err)
		}
		tags[i] = tag
	}
	best, _, _ := matcher.Match(tags...)
	return localeFromTag(best)
}

func localeFromTag(in language.Tag) fyne.Locale {
	b, _ := in.Base()
	r, _ := in.Region()
	s, _ := in.Script()

	return fyne.Locale(b.String() + "-" + r.String() + "-" + s.String())
}

func AddTranslationsFS(fs embed.FS, dir string) (retErr error) {
	files, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		name := f.Name()

		data, err := fs.ReadFile(dir + "/" + name)
		if err != nil {
			if retErr == nil {
				retErr = err
			}
			continue
		}

		err = addLanguage(data, name)
		if err != nil {
			if retErr == nil {
				retErr = err
			}
			continue
		}
	}

	return retErr
}

func addLanguage(data []byte, name string) error {
	f, err := bundle.ParseMessageFileBytes(data, name)
	if err != nil {
		return err
	}

	translated = append(translated, f.Tag)
	return nil
}
