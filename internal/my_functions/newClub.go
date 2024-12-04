package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func NewClub(name string) (c mt.Club, err error) {
	if name == "" {
		err = fmt.Errorf("club name cannot be empty")
		return mt.Club{}, err
	}

	c.Name = name
	c.PlayerList = []*mt.Player{}
	c.TeamList = []*mt.Team{}
	return c, nil
}
