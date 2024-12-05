package my_types

import (
	"fmt"
	"log"
)

func (p Player) String() string {
	return fmt.Sprintf("%s (Age: %d, Ranking: %d, Material: %v)",
		p.Name,
		p.Age,
		p.Ranking,
		p.Material,
	)
}

func (p *Player) Show() error {
	if p == nil {
		log.Println("Error! Player does not exist (nil pointer)")
		return fmt.Errorf("player does not exist")
	}

	if p.IsEmpty() {
		log.Println("Error ! Player does not exist")
		return fmt.Errorf("Player does not exist")
	}

	fmt.Println("Showing characteristics of", p.Name)

	// Get team name, we don't want to show the address of the pointer
	teams := []string{}
	for _, team := range p.TeamList {
		teams = append(teams, team.Name)
	}

	fmt.Printf("%s, Age: %d, Ranking: %d, Material: %v, Teams: %v\n",
		p.Name,
		p.Age,
		p.Ranking,
		p.Material,
		teams,
	)
	return nil
}

func TeamNames(t []*Team) []string {
	teamNames := []string{}
	for _, team := range t {
		teamNames = append(teamNames, team.Name)
	}
	return teamNames
}

func (t *Team) Show() error {
	if t == nil {
		log.Println("Error! Team does not exist (nil pointer)")
		return fmt.Errorf("team does not exist")
	}

	if t.IsEmpty() {
		log.Println("Error ! Team does not exist")
		return fmt.Errorf("Team does not exist")
	}
	fmt.Println("Showing the characteristics of", t.Name)
	n := len(t.PlayerList)
	if n >= 1 {
		for i := 0; i < n; i++ {
			fmt.Printf("Player %v: %v\n",
				i+1,
				t.PlayerList[i].Name,
			)
		}

	} else {
		fmt.Printf("There is no player in %v.\n",
			t.Name,
		)
	}
	return nil
}

func (c *Club) Show() error {
	if c == nil {
		log.Println("Error! Club does not exist (nil pointer)")
		return fmt.Errorf("club does not exist")
	}

	if c.IsEmpty() {
		fmt.Println("Error ! Club does not exist")
		return fmt.Errorf("Club does not exist")
	}
	fmt.Println("Characteristics of", c.Name)

	if n := len(c.TeamList); n <= 1 {
		fmt.Printf("%v has %v team.\n", c.Name, n)
	} else {
		fmt.Printf("%v has %v teams.\n", c.Name, n)
	}
	for i := 0; i < len(c.TeamList); i++ {
		m := len(c.TeamList[i].PlayerList)
		if m <= 1 {
			fmt.Printf("Team %v: %v. %v player.\n",
				i+1,
				c.TeamList[i].Name,
				m,
			)
		} else {
			fmt.Printf("Team %v: %v. %v players.\n",
				i+1,
				c.TeamList[i].Name,
				m,
			)
		}
	}

	if n := len(c.PlayerList); n <= 1 {
		fmt.Printf("%v has %v player.\n", c.Name, n)
	} else {
		fmt.Printf("%v has %v players.\n", c.Name, n)
		for i := 0; i < len(c.PlayerList); i++ {
			fmt.Printf("Player %v: %s, Age: %d, Ranking: %d, Material: %v, Team: %v\n",
				i+1,
				c.PlayerList[i].Name,
				c.PlayerList[i].Age,
				c.PlayerList[i].Ranking,
				c.PlayerList[i].Material,
				TeamNames(c.PlayerList[i].TeamList),
			)
		}
	}
	return nil
}

func (d *Database) Show() error {
	fmt.Println("Clubs :")
	for _, club := range d.ClubList {
		err1 := club.Show()
		if err1 != nil {
			return err1
		}
	}
	fmt.Println("Teams :")
	for _, team := range d.TeamList {
		err2 := team.Show()
		if err2 != nil {
			return err2
		}
	}
	fmt.Println("Players :")
	for _, player := range d.PlayerList {
		err3 := player.Show()
		if err3 != nil {
			return err3
		}
	}
	return nil
}
