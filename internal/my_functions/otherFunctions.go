package my_functions

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

const DefaultMaterial = "Unknown"

func DefaultPlayerMaterial() []string {
	return []string{DefaultMaterial, DefaultMaterial, DefaultMaterial}
}

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
