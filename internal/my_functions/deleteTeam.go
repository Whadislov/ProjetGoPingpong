package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func DeleteTeam(t *mt.Team, c *mt.Club) error {
	if err := c.DeleteTeam(t); err != nil {
		return fmt.Errorf("error when deleting team %s. Reason : %w", t.Name, err)
	}
	return nil
}
