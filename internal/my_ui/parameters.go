package myapp

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var appStartOption string
var userOfSession *mt.User
var credToken string

// For the UserPage to know the current username, email, password
var currentUsername string
var currentEmail string

var darkTheme myDarkTheme
var lightTheme myLightTheme

// Get translations
var localizer *i18n.Localizer
var currentSelectedLanguage string = "English"
