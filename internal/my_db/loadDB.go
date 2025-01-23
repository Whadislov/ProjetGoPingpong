package mydb

import (
	"fmt"
	"log"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// LoadPlayers loads players from the database into the player map.
func (db *Database) LoadPlayers() (map[int]*mt.Player, error) {
	rows, err := db.Conn.Query(fmt.Sprintln("SELECT id, name, age, ranking, forehand, backhand, blade FROM players WHERE user_id = ", userIDOfSession))
	if err != nil {
		return nil, fmt.Errorf("failed to load players: %w", err)
	}
	defer rows.Close()

	var players = make(map[int]*mt.Player)
	for rows.Next() {
		var player mt.Player
		player.Material = []string{"", "", ""}
		player.TeamIDs = make(map[int]string)
		player.ClubIDs = make(map[int]string)

		err := rows.Scan(&player.ID, &player.Name, &player.Age, &player.Ranking, &player.Material[0], &player.Material[1], &player.Material[2])
		if err != nil {
			return nil, fmt.Errorf("failed to scan player: %w", err)
		}
		players[player.ID] = &player
	}

	return players, rows.Err()
}

// LoadTeams loads teams from the database into the team map.
func (db *Database) LoadTeams() (map[int]*mt.Team, error) {
	rows, err := db.Conn.Query(fmt.Sprintln("SELECT id, name FROM teams WHERE user_id = ", userIDOfSession))
	if err != nil {
		return nil, fmt.Errorf("failed to load teams: %w", err)
	}
	defer rows.Close()

	var teams = make(map[int]*mt.Team)
	for rows.Next() {
		var team mt.Team
		team.PlayerIDs = make(map[int]string)
		team.ClubID = make(map[int]string)
		err := rows.Scan(&team.ID, &team.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team: %w", err)
		}
		teams[team.ID] = &team
	}
	return teams, rows.Err()
}

// LoadClubs loads clubs from the database into the club map.
func (db *Database) LoadClubs() (map[int]*mt.Club, error) {
	rows, err := db.Conn.Query(fmt.Sprintln("SELECT id, name FROM clubs WHERE user_id = ", userIDOfSession))
	if err != nil {
		return nil, fmt.Errorf("failed to load clubs: %w", err)
	}
	defer rows.Close()

	var clubs = make(map[int]*mt.Club)
	for rows.Next() {
		var club mt.Club
		club.PlayerIDs = make(map[int]string)
		club.TeamIDs = make(map[int]string)
		err := rows.Scan(&club.ID, &club.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan club: %w", err)
		}
		clubs[club.ID] = &club
	}

	return clubs, rows.Err()
}

// LoadPlayerClubs loads the player-club relationships from the database.
func (db *Database) LoadPlayerClubs(players map[int]*mt.Player, clubs map[int]*mt.Club) error {
	rows, err := db.Conn.Query(fmt.Sprintln("SELECT player_id, club_id FROM player_club WHERE user_id = ", userIDOfSession))
	if err != nil {
		return fmt.Errorf("failed to load player_club relationships: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var playerID, clubID int
		err := rows.Scan(&playerID, &clubID)
		if err != nil {
			return fmt.Errorf("failed to scan player_club relationship: %w", err)
		}
		if player, ok := players[playerID]; ok {
			player.ClubIDs[clubID] = clubs[clubID].Name
		}
		if club, ok := clubs[clubID]; ok {
			club.PlayerIDs[playerID] = players[playerID].Name
		}
	}
	return rows.Err()
}

// LoadPlayerTeams loads the player-team relationships from the database.
func (db *Database) LoadPlayerTeams(players map[int]*mt.Player, teams map[int]*mt.Team) error {
	rows, err := db.Conn.Query(fmt.Sprintln("SELECT player_id, team_id FROM player_team WHERE user_id = ", userIDOfSession))
	if err != nil {
		return fmt.Errorf("failed to load player_team relationships: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var playerID, teamID int
		err := rows.Scan(&playerID, &teamID)
		if err != nil {
			return fmt.Errorf("failed to scan player_team relationship: %w", err)
		}
		if player, ok := players[playerID]; ok {
			player.TeamIDs[teamID] = teams[teamID].Name
		}
		if team, ok := teams[teamID]; ok {
			team.PlayerIDs[playerID] = players[playerID].Name
		}
	}

	return rows.Err()
}

// LoadTeamClubs loads the team-club relationships from the database.
func (db *Database) LoadTeamClubs(teams map[int]*mt.Team, clubs map[int]*mt.Club) error {
	rows, err := db.Conn.Query(fmt.Sprintln("SELECT team_id, club_id FROM team_club WHERE user_id = ", userIDOfSession))
	if err != nil {
		return fmt.Errorf("failed to load team_club relationships: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var teamID, clubID int
		err := rows.Scan(&teamID, &clubID)
		if err != nil {
			return fmt.Errorf("failed to scan team_club relationship: %w", err)
		}
		if team, ok := teams[teamID]; ok {
			team.ClubID[clubID] = clubs[clubID].Name
		}
		if club, ok := clubs[clubID]; ok {
			club.TeamIDs[teamID] = teams[teamID].Name
		}
	}
	return rows.Err()
}

// LoadUsers loads users from the database into the user map.
func (db *Database) LoadUsers() (map[int]*mt.User, error) {
	rows, err := db.Conn.Query("SELECT id, username, email, password_hash FROM user")
	if err != nil {
		return nil, fmt.Errorf("failed to load users: %w", err)
	}
	defer rows.Close()

	var users = make(map[int]*mt.User)
	for rows.Next() {
		var user mt.User
		err := rows.Scan(&user)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users[user.ID] = &user
	}

	return users, rows.Err()
}

// LoadDB loads the database.
func LoadDB() (*mt.Database, error) {
	db, err := ConnectToDB()
	if err != nil {
		fmt.Println("Error while connecting to postgresql database:", err)
		return nil, err
	}

	users, err := db.LoadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Name == usernameOfSession {
			userIDOfSession = user.ID
		}
	}

	players, err := db.LoadPlayers()
	if err != nil {
		return nil, err
	}
	teams, err := db.LoadTeams()
	if err != nil {
		return nil, err
	}
	clubs, err := db.LoadClubs()
	if err != nil {
		return nil, err
	}
	err = db.LoadPlayerTeams(players, teams)
	if err != nil {
		return nil, err
	}
	err = db.LoadPlayerClubs(players, clubs)
	if err != nil {
		return nil, err
	}
	err = db.LoadTeamClubs(teams, clubs)
	if err != nil {
		return nil, err
	}
	golangDB := &mt.Database{
		Players: players,
		Teams:   teams,
		Clubs:   clubs,
	}
	log.Println("Database loaded successfully")
	defer db.Close()
	return golangDB, nil
}

// Choose between Sqlite for local developement or POSTgreSQL for the production
