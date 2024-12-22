package my_functions

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
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
