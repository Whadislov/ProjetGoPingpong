package my_types

import (
	"fmt"
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

func (p Player) Show()(err error) {
	if p.IsEmpty() {
		fmt.Println("Error ! Player does not exist")
		return fmt.Errorf("Player does not exist")
	}

	fmt.Printf("%v is %v years old. Ranking : %v. %v plays with %v on the forehand, %v on the backhand and %v.\n", p.Name, p.Age, p.Ranking, p.Name, p.Material[0], p.Material[1], p.Material[2])
	switch len(p.TeamList){
	case 0: {
		fmt.Printf("%v does not have a team.\n", p.Name)
		return nil
	}
	case 1: {
		fmt.Printf("%v plays in %v.\n", p.Name, p.TeamList[0].Name)
		return nil
	}
	case 2: {
		fmt.Printf("%v plays in %v and in %v.\n", p.Name, p.TeamList[0].Name, p.TeamList[1].Name)
		return nil
	}
	default: {
		fmt.Printf("%v plays in more than 2 teams.\n", p.Name)
		return nil
	}
	}
}

func (t Team) Show()(err error) {
	if t.IsEmpty() {
		fmt.Println("Error ! Team does not exist")
		return fmt.Errorf("Team does not exist")
	}
	fmt.Println("Team name:", t.Name)
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

func (c Club) Show()(err error) {
	if c.IsEmpty() {
		fmt.Println("Error ! Club does not exist")
		return fmt.Errorf("Club does not exist")
	}

	fmt.Println("Club name :", c.Name)
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
			fmt.Println(c.PlayerList[i])
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
			break
		}
	}

	if !found {
		team.PlayerList = append(team.PlayerList, player)
	}

	// Add team to the list of team of the player
	found = false
	for _, t := range player.TeamList {
		if t.Name == team.Name {
			found = true
			break
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

func (c *Club) RemovePlayer(player *Player)(error) {
	index, err := c.FindPlayer(player.Name)
	if err != nil {
		return err
	}
	// Move the player at the end of the club list
	c.PlayerList[index] = c.PlayerList[len(c.PlayerList)-1]
	// Remove the player from the list
	c.PlayerList = c.PlayerList[:len(c.PlayerList)-1]


	for _, team := range player.TeamList {
		if err := c.RemovePlayerFromTeam(player, team.Name); err != nil {
			fmt.Printf("Error when deleting player %s from team %s : %s\n", player.Name, team.Name, err)
		}
	}
	// Delete the player, make it empty
	*player = Player{}
	return nil
}

	func (c *Club) RemovePlayerFromTeam(player *Player, teamName string) error {
		// Find team, club view
		teamIndex, err := c.FindTeam(teamName)
		if err != nil {
			return fmt.Errorf("team %s not found: %w", teamName, err)
		}
	
		// Find player in team
		playerIndex := -1
		for i, p := range c.TeamList[teamIndex].PlayerList {
			if p.Name == player.Name {
				playerIndex = i
				break
			}
		}
		if playerIndex == -1 {
			return fmt.Errorf("player %s not found in team %s", player.Name, teamName)
		}
	
		// Delete player in the list
		// ... means we append element per element, the slice is sliced in individual element before being added
		c.TeamList[teamIndex].PlayerList = append(
			c.TeamList[teamIndex].PlayerList[:playerIndex], 
			c.TeamList[teamIndex].PlayerList[playerIndex+1:]..., 
		)
	
		// Delete team from the team list of the player
		teamIndexInPlayer := -1
		for i, t := range player.TeamList {
			if t.Name == teamName {
				teamIndexInPlayer = i
				break
			}
		}
		if teamIndexInPlayer != -1 {
			player.TeamList = append(
				player.TeamList[:teamIndexInPlayer],
				player.TeamList[teamIndexInPlayer+1:]...,
			)
		}
	
		return nil
	}


func (c *Club) RemoveTeam(teamName string)(error) {
	index, err := c.FindTeam(teamName)
	if err != nil {
		return err
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
	teamToRemove = nil
	return nil
}

func (c *Club) GetPlayerList()([]*Player) {
	return c.PlayerList
}

func (p *Player) GetTeam(teamName string, c Club)(*Team) {
	for _, team := range c.TeamList {
		fmt.Println("c.TeamList", c.TeamList)
		if team.Name == teamName {
			return team
		}
	}
	return nil
}

func (c *Club) GetTeamList()([]*Team) {
	return c.TeamList
}
