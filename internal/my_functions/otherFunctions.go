package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"regexp"
)

const DefaultMaterial = "Unknown"

// DefaultPlayerMaterial returns a slice of strings representing the default material for a player.
func DefaultPlayerMaterial() []string {
	return []string{DefaultMaterial, DefaultMaterial, DefaultMaterial}
}

// GetName returns the name of the given entity (Player, Team, or Club).
// Returns an empty string if the entity type is not recognized.
func GetName(x interface{}) string {
	switch v := x.(type) {
	case mt.Player:
		{
			return v.Firstname + v.Lastname
		}
	case *mt.Player:
		{
			return v.Firstname + v.Lastname
		}
	case mt.Team:
		{
			return v.Name
		}
	case *mt.Team:
		{
			return v.Name
		}
	case mt.Club:
		{
			return v.Name
		}
	case *mt.Club:
		{
			return v.Name
		}
	default:
		{
			return ""
		}
	}
}

// isValidName verifies that the name follows some criterias
func IsValidName(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}

	for _, r := range name {
		if r < 'A' || r > 'z' {
			return false, fmt.Errorf("name can only contain letters")
		}
	}
	return true, nil
}

// isValidName verifies that the name follows some criterias
func IsValidUsername(username string) (bool, error) {
	if username == "" {
		return false, fmt.Errorf("username cannot be empty")
	}

	usernameRegex := `^[a-zA-Z0-9_]+$`

	// Compile the regex
	re := regexp.MustCompile(usernameRegex)

	// Verify if the string is a regex
	if re.MatchString(username) {
		return re.MatchString(username), nil
	} else {
		return re.MatchString(username), fmt.Errorf("username must be valid (only letters and figures are allowed, spaces are not allowed)")
	}
}

// isValidEmail verifies that the name follows a valid regex
func IsValidEmail(email string) (bool, error) {
	if email == "" {
		return false, fmt.Errorf("e-mail cannot be empty")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Verify if the string is a regex, true means yes
	if re.MatchString(email) {
		return re.MatchString(email), nil
	} else {
		return re.MatchString(email), fmt.Errorf("e-mail must be valid. Example: abc@def.com")
	}
}

// isValidEmail verifies that the password is not empty and does no contain spaces
func IsValidPassword(password string) (bool, error) {
	if password == "" {
		return false, fmt.Errorf("password cannot be empty")
	}

	for _, char := range password {
		if char == ' ' {
			return false, fmt.Errorf("password must be valid (spaces are not allowed)")
		}
	}

	return true, nil
}
