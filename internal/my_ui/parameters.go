package myapp

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

var appStartOption string
var userOfSession *mt.User

// For the UserPage to know the current username, email, password
var currentUsername string
var currentEmail string
var currentPassword string
