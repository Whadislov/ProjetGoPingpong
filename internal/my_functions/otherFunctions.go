package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
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
			return v.Name
		}
	case *mt.Player:
		{
			return v.Name
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
func isValidName(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("username cannot be empty")
	}

	for _, r := range name {
		if r < 'A' || r > 'z' {
			return false, fmt.Errorf("username can only contain letters")
		}
	}
	return true, nil
}

// isValidName verifies that the name follows some criterias
func isValidUsername(username string) (bool, error) {
	if username == "" {
		return false, fmt.Errorf("username cannot be empty")
	}

	usernameRegex := `^[a-zA-Z0-9_]+$`

	// Compile the regex
	re := regexp.MustCompile(usernameRegex)

	// Verify if the string is a regex
	return re.MatchString(username), nil
}

// isValidEmail verifies that the name follows a valid regex
func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Verify if the string is a regex
	return re.MatchString(email)
}
