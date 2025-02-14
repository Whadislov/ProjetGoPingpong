package myapp

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

var appStartOption string
var userOfSession *mt.User
var jsonWebToken string

// For the UserPage to know the current username, email, password
var currentUsername string
var currentEmail string
var currentPassword string
