package myapp

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

var appStartOption string

func AppStartOption(s string) {
	appStartOption = s
}

var userOfSession *mt.User

func SetUsernameOfSession(u *mt.User) {
	userOfSession = u
}

// For the UserPage to know the current username, email, password
var currentUsername string
var currentEmail string
var currentPassword string
