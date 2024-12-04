package my_types

import (
	"fmt"
	"log"
)

type Player struct {
	Name 		string
	Age 		int
	Ranking		int
	Material 	[]string
	TeamList	[]*Team  
}

type Team struct {
	Name 		string
	PlayerList 	[]*Player
}

type Club struct {
	Name 		string
	PlayerList	[]*Player
	TeamList	[]*Team //Teams map[string]*Team
}

type PlayerMatch struct{
	League				string
	Player				string
	Adversary 			string
	PlayerMatchOutcome	Outcome			
}

type Outcome struct {
	Victory	string
	Defeat	string
	Draw	string
}

type TeamMatch struct{
	League				string
	TeamComposition 	map[int]string
	
}

type Match struct {
	HomeTeam			Team
	GuestTeam			Team
	TeamMatchOutcome	Outcome
}

func (p *Player) SetPlayerAge(age int) {
	p.Age = age
}

func (p *Player) SetPlayerRanking(ranking int) {
	p.Ranking = ranking
}

func (p *Player) SetPlayerMaterial(forehand string, backhand string, blade string) {
	material := []string{forehand, backhand, blade}
	p.Material = material
}

// ajouter erreur si la value de la clé (joueur) ne fait pas partie de la liste des joueurs
func (t *TeamMatch) SetTeamComposition(teamComposition map[int]string) {
	t.TeamComposition = teamComposition
}

func (p *Player) IsEmpty()(bool) {
	return p.Name == "" &&
	p.Age == 0 &&
	p.Ranking == 0 && 
	p.Material == nil &&
	p.TeamList == nil
}

func (t *Team) IsEmpty()(bool) {
	return t.Name == "" &&
	t.PlayerList == nil
}

func (c *Club) IsEmpty()(bool) {
	return c.Name == "" &&
	c.PlayerList == nil &&
	c.TeamList == nil
}

func (p *Player) Show()(error) {
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

func (t *Team) Show()(error) {
	if t == nil {
        log.Println("Error! Team does not exist (nil pointer)")
        return fmt.Errorf("team does not exist")
    }

	if t.IsEmpty() {
		log.Println("Error ! Team does not exist")
		return fmt.Errorf("Team does not exist")
	}
	fmt.Println("Showing characteristics of", t.Name)
	switch len(t.PlayerList)  {
	case 0: {
		fmt.Printf("%v has 0 player.\n", t.Name)
	}
	case 1: {
		fmt.Printf("%v has 1 player.\n", t.Name)
		fmt.Println("The player is :", t.PlayerList[0])
	}
	default: {
		fmt.Printf("%v has %v players.\n", t.Name, len(t.PlayerList))
		fmt.Println("The players are :")
		for i := 0; i < len(t.PlayerList); i++ {
			fmt.Printf("Player %v : %v\n", i+1, t.PlayerList[i].Name)
		}
		}
	}
	return nil
}

func (p Player) String() string {
	return fmt.Sprintf("%s (Age: %d, Ranking: %d, Material: %v)", 
		p.Name, 
		p.Age, 
		p.Ranking, 
		p.Material,
	)
}

func (c *Club) Show()(error) {
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
				c.PlayerList[i].TeamList,
			)
		}
	}
	return nil
	}


func (c *Club) AddPlayer(player *Player) {
	c.PlayerList = append(c.PlayerList, player)
}
	
func (c *Club) AddTeam(team *Team) {
	c.TeamList = append(c.TeamList, team)
}

func (c *Club) AddPlayerToTeam(player *Player, team *Team)(error) {
	err := c.FindTeam(team)
	if err != nil {
		return err
	}

	// Add player the in the team if not already in
	found := false
	for _, p := range team.PlayerList {
		if p.Name == player.Name {
			found = true

			// Add team to the list of team of the player if not already in (can debug the link between the team and the player)
			found2 := false
			for _, t := range player.TeamList {
				if t.Name == team.Name {
					found2 = true
					return fmt.Errorf("team %s and player %s are already in each other's respective list", team.Name, player.Name)
				}
			}
			if !found2 {
				player.TeamList = append(player.TeamList, team)
			}
			return fmt.Errorf("player %s is already in team %s", player.Name, team.Name)
		}
	}

	if !found {
		team.PlayerList = append(team.PlayerList, player)
	}

	// Add team to the list of team of the player if not already in
	found = false
	for _, t := range player.TeamList {
		if t.Name == team.Name {
			found = true
			return fmt.Errorf("team %s and player %s are already in each other's respective list", team.Name, player.Name)
		}
	}
	if !found {
		player.TeamList = append(player.TeamList, team)
	}
	log.Printf("%s has been successfully added in %s", player.Name, team.Name)
	return nil
}

func (c *Club) FindTeam(team *Team)(error) {
	for i := range c.TeamList {
		if c.TeamList[i].Name == team.Name {
			return nil
		}
	}
	return fmt.Errorf("%v not found in %v", team.Name, c.Name)
}

func (c *Club) FindPlayer(player *Player)(error) {
	for i := range c.PlayerList {
		if c.PlayerList[i].Name == player.Name {
			return nil
		}
	}
	return fmt.Errorf("%v not found in %v", player.Name, c.Name)
}

func (c *Club) DeletePlayer(player *Player)(error) {
	pName := player.Name
	err := c.FindPlayer(player)
	if err != nil {
		return fmt.Errorf("error when deleting player %s : %w", player.Name, err)
	}
	

	// Copy of the slice (working directly on the modified slice cause probleme with index)
	teamListCopy := append([]*Team{}, player.TeamList...)

	for _, team := range teamListCopy {
		if err := c.RemovePlayerFromTeam(player, team); err != nil {
			log.Printf("%v has not been successfully deleted from %v. Reason : %v", player.Name, c.Name, err)
			return err
		}
	}
	
	// Find the index of the player in the club playerlist.
	index := 0
	for _, p := range c.PlayerList {
		if p.Name == player.Name {
			c.PlayerList[index] = c.PlayerList[len(c.PlayerList)-1]
			//index = i
			break
		}
	}

	// Remove the player from the list
	c.PlayerList = c.PlayerList[:len(c.PlayerList)-1]

	// Delete the player, make it empty
	*player = Player{}
	log.Printf("%v has been successfully deleted from %v", pName, c.Name)
	return nil
}

	func (c *Club) RemovePlayerFromTeam(player *Player, team *Team) error {
		// **** Verifications with club lists
		// Find player, club view
		err := c.FindPlayer(player)

		if err != nil {
			log.Printf("%v has not been successfully removed from %v. Reason : %v", player.Name, team.Name, err)
			return err
		}
		
		// Find team, club view
		err2 := c.FindTeam(team)
		if err2 != nil {
			log.Printf("%v has not been successfully removed from %v. Reason : %v", player.Name, team.Name, err2)
			return err2
		}

		// **** Delete player from the team player list 
		// Find player in team playerlist
		playerIndex := -1
		for i, p := range team.PlayerList {
			if p.Name == player.Name {
				playerIndex = i
				break
			}
		}
		// Player is not in team playerlist
		if playerIndex == -1 {
			// Check if the team is in the player's teamlist
			for teamIndexInPlayer := range player.TeamList {
				if player.TeamList[teamIndexInPlayer].Name == team.Name {
					// Found the team in the player's teamlist. Need to remove the team from the list
					copy(player.TeamList[teamIndexInPlayer:], player.TeamList[teamIndexInPlayer+1:])
					player.TeamList[len(player.TeamList)-1] = nil // Clean the last position
					player.TeamList = player.TeamList[:len(player.TeamList)-1]
				}
			}
			err := fmt.Errorf("player %s does not belong to team %s", player.Name, team.Name)
			log.Printf("%v has not been successfully removed from %v. Reason : %v", player.Name, team.Name, err)
			return err
		}

		copy(team.PlayerList[playerIndex:], team.PlayerList[playerIndex+1:])
		team.PlayerList[len(team.PlayerList)-1] = nil // Clean the last position
		team.PlayerList = team.PlayerList[:len(team.PlayerList)-1]
	
		// **** Delete team from the player team list 
		teamIndexInPlayer := -1
		for i, t := range player.TeamList {
			if t.Name == team.Name {
				teamIndexInPlayer = i
				break
			}
		}

		// Find team in player teamlist
		if teamIndexInPlayer == -1 {
			err := fmt.Errorf("team %s is not in player %s's team list", team.Name, player.Name)
			log.Printf("%v has not been successfully removed from %v. Reason : %v", player.Name, team.Name, err)
			return err
		}

		// Team is not in player teamlist
		if teamIndexInPlayer != -1 {
			copy(player.TeamList[teamIndexInPlayer:], player.TeamList[teamIndexInPlayer+1:])
			player.TeamList[len(player.TeamList)-1] = nil // Clean the last position
			player.TeamList = player.TeamList[:len(player.TeamList)-1]
		}
		log.Printf("Player %v has been successfully removed from team %v.", player.Name, team.Name)
		return nil
	}


func (c *Club) DeleteTeam(team *Team)(error) {
	err := c.FindTeam(team)
	if err != nil {
		return err
	}

	tName := team.Name

	// Remove the link between players and the team
	for _, p := range team.PlayerList {
		if err := c.RemovePlayerFromTeam(p, team); err != nil {
			return err
		}
	}
	
	// Create list of teams without the removed team
	newTeamList := []*Team{}

	for _, t := range c.TeamList {
		if t.Name != team.Name {
			newTeamList = append(newTeamList, t)
		} else {
			team.Name = ""
			team.PlayerList = nil
		}
	}
	// Set the list
	c.TeamList = newTeamList

	log.Printf("%v has been successfully deleted from %v", tName, c.Name)
	return nil
}
