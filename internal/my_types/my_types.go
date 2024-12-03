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
	fmt.Println("Showing characteristics of", c.Name)
	switch len(c.TeamList)  {
	case 0: {
		fmt.Printf("%v has 0 team.\n", c.Name)
	}
	case 1: {
		fmt.Printf("%v has 1 team.\n", c.Name)
		fmt.Println("The team is :", c.TeamList[0])
	}
	default: {
		fmt.Printf("%v has %v teams.\n", c.Name, len(c.TeamList))
		fmt.Println("The teams are :")
		for i := 0; i < len(c.TeamList); i++ {
			fmt.Println(c.TeamList[i].Name)

			if len(c.TeamList[i].PlayerList) == 0 {
				fmt.Println("There is no player in this team")
			}
			for j := 0; j < len(c.TeamList[i].PlayerList); j++ {
				fmt.Printf("Player %v : %v\n", j+1, c.TeamList[i].PlayerList[j].Name)
			}
		}
	}
	}
	switch len(c.PlayerList)  {
	case 0: {
		fmt.Printf("%v has 0 player.\n", c.Name)
	}
	case 1: {
		fmt.Printf("%v has 1 player.\n", c.Name)
		fmt.Println("The player is :", c.PlayerList[0])
	}
	default: {
		fmt.Printf("%v has %v players.\n", c.Name, len(c.PlayerList))
		fmt.Println("The players are :")
		for i := 0; i < len(c.PlayerList); i++ {	
			fmt.Printf("Player %v: %s, Age: %d, Ranking: %d, Material: %v\n", 
				i+1, 
				c.PlayerList[i].Name, 
				c.PlayerList[i].Age, 
				c.PlayerList[i].Ranking, 
				c.PlayerList[i].Material,
			)
				
		}
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

func (c *Club) AddPlayerToTeam(player *Player, teamName string)(error) {
	teamIndex, err := c.FindTeam(teamName)
	if err != nil {
		return err
	}
	team := c.TeamList[teamIndex]

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
					return fmt.Errorf("team %s and player %s are already in each other's respective list", teamName, player.Name)
				}
			}
			if !found2 {
				player.TeamList = append(player.TeamList, team)
			}
			return fmt.Errorf("player %s is already in team %s", player.Name, teamName)
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
			return fmt.Errorf("team %s and player %s are already in each other's respective list", teamName, player.Name)
		}
	}
	if !found {
		player.TeamList = append(player.TeamList, team)
	}
	return nil
}

func (c *Club) FindTeam(teamName string)(int, error) {
	for i := range c.TeamList {
		if c.TeamList[i].Name == teamName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%v not found in the club", teamName)
}

func (c *Club) FindPlayer(playerName string)(int, error) {
	for i := range c.PlayerList {
		if c.PlayerList[i].Name == playerName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%v not found in the club", playerName)
}

func (c *Club) DeletePlayer(player *Player)(error) {
	index, err := c.FindPlayer(player.Name)
	if err != nil {
		return fmt.Errorf("error when deleting player %s : %w", player.Name, err)
	}
	

	// Copy of the slice (working directly on the modified slice cause probleme with index)
	teamListCopy := append([]*Team{}, player.TeamList...)

	for _, team := range teamListCopy {
		if err := c.RemovePlayerFromTeam(player, team.Name); err != nil {
			return fmt.Errorf("error when deleting player %s from team %s : %s", player.Name, team.Name, err)
		}
	}
	
	// Move the player at the end of the club list
	c.PlayerList[index] = c.PlayerList[len(c.PlayerList)-1]
	// Remove the player from the list
	c.PlayerList = c.PlayerList[:len(c.PlayerList)-1]

	// Delete the player, make it empty
	*player = Player{}
	return nil
}

	func (c *Club) RemovePlayerFromTeam(player *Player, teamName string) error {
		// Find player, club view
		_, err := c.FindPlayer(player.Name)

		if err != nil {
			return err
		}
		
		// Find team, club view
		teamIndex, err2 := c.FindTeam(teamName)
		if err2 != nil {
			return fmt.Errorf("team %s not found: %w", teamName, err2)
		}

		// Find player in team
		playerIndex := -1
		for i, p := range c.TeamList[teamIndex].PlayerList {
			if p.Name == player.Name {
				playerIndex = i
				break
			}
		}
		// player is not in team playerlist
		if playerIndex == -1 {
			// Check if the team is in the player's teamlist
			for teamIndexInPlayer := range player.TeamList {
				if player.TeamList[teamIndexInPlayer].Name == teamName {
					// Found the team in the player's teamlist. Need to remove the team from the list
					copy(player.TeamList[teamIndexInPlayer:], player.TeamList[teamIndexInPlayer+1:])
					player.TeamList[len(player.TeamList)-1] = nil // Clean the last position
					player.TeamList = player.TeamList[:len(player.TeamList)-1]
				}
			}
			return fmt.Errorf("player %s does not belong to team %s", player.Name, teamName)
		}

		copy(c.TeamList[teamIndex].PlayerList[playerIndex:], c.TeamList[teamIndex].PlayerList[playerIndex+1:])
		c.TeamList[teamIndex].PlayerList[len(c.TeamList[teamIndex].PlayerList)-1] = nil // Clean the last position
		c.TeamList[teamIndex].PlayerList = c.TeamList[teamIndex].PlayerList[:len(c.TeamList[teamIndex].PlayerList)-1]
	
		// Delete team from the team list of the player
		teamIndexInPlayer := -1
		for i, t := range player.TeamList {
			if t.Name == teamName {
				teamIndexInPlayer = i
				break
			}
		}

		// Team not found in player's team list
		if teamIndexInPlayer == -1 {
			return fmt.Errorf("team %s is not in player %s's team list", teamName, player.Name)
		}

		// Team found
		if teamIndexInPlayer != -1 {
			copy(player.TeamList[teamIndexInPlayer:], player.TeamList[teamIndexInPlayer+1:])
			player.TeamList[len(player.TeamList)-1] = nil // Clean the last position
			player.TeamList = player.TeamList[:len(player.TeamList)-1]
		}
		//log.Printf("Player %v has been successfully removed from team %v.", player.Name, teamName)
		return nil
	}


func (c *Club) DeleteTeam(teamName string)(error) {
	index, err := c.FindTeam(teamName)
	if err != nil {
		return fmt.Errorf("error when deleting team %s : %w", teamName, err)
	}
	// Team to remove
	teamToRemove := c.TeamList[index]

	// Remove the link between players and the team
	for _, p := range teamToRemove.PlayerList {
		if err := c.RemovePlayerFromTeam(p, teamName); err != nil {
			return fmt.Errorf("error when deleting player %s from team %s : %w", p.Name, teamName, err)
		}
	}
	
	// Create list of teams without the removed team
	newTeamList := []*Team{}

	for i, t := range c.TeamList {
		if i != index {
			newTeamList = append(newTeamList, t)
		} else {
			c.TeamList[index] = nil
		}
	}
	// Set the list
	c.TeamList = newTeamList
	// Set all values to "" or nil
	teamToRemove.Name = ""
    teamToRemove.PlayerList = nil
	// Set pointer to nil
	//teamToRemove = nil
	//log.Printf("Team %v has been successfully deleted.", teamName)
	return nil
}
