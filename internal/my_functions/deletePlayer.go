package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func DeletePlayer(p *mt.Player, c *mt.Club) error {
	if err := c.DeletePlayer(p); err != nil {
		return fmt.Errorf("error when deleting player %s. Reason : %w", p.Name, err)
	}
	return nil
}
