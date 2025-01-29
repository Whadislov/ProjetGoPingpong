package mydb

import (
	"fmt"
	"log"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// SavePlayers saves players in the database.
func (db *Database) SaveUsers(users map[int]*mt.User) error {
	for _, user := range users {
		query := `
        INSERT INTO users (id, username, email, password_hash, created_at)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (id) DO UPDATE SET
            username = EXCLUDED.username,
            email = EXCLUDED.email,
            password_hash = EXCLUDED.password_hash;
        `
		_, err := db.Conn.Exec(query, user.ID, user.Name, user.Email, user.PasswordHash, user.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to save user: %w", err)
		}
	}
	return nil
}

// SavePlayers saves players in the database.
func (db *Database) SavePlayers(players map[int]*mt.Player) error {
	for _, player := range players {
		if player.ID < 0 {
			// Let postgresql creates its own ID for a new player
			var postgresPlayerID int
			query := `
			INSERT INTO players (id, name, age, ranking, forehand, backhand, blade, user_id)
			VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7)
			RETURNING id;
			`
			err := db.Conn.QueryRow(query, player.Name, player.Age, player.Ranking, player.Material[0], player.Material[1], player.Material[2], userIDOfSession).Scan(&postgresPlayerID)
			if err != nil {
				return fmt.Errorf("failed to save the new player: %w", err)
			}
			// Change the ID for the relationship tables
			player.ID = postgresPlayerID
		} else {
			// Modify the player if it's not new
			query := `
			UPDATE players 
			SET name = $1, age = $2, ranking = $3, forehand = $4, backhand = $5, blade = $6
			WHERE id = $7;
			`
			_, err := db.Conn.Exec(query, player.Name, player.Age, player.Ranking, player.Material[0], player.Material[1], player.Material[2], player.ID)
			if err != nil {
				return fmt.Errorf("failed to save the edited player: %w", err)
			}
		}
	}
	return nil
}

// SaveTeams saves teams in the database.
func (db *Database) SaveTeams(teams map[int]*mt.Team) error {
	for _, team := range teams {
		if team.ID < 0 {
			// Let postgresql creates its own ID for a new team
			var postgresTeamID int
			query := `
			INSERT INTO teams (id, name, user_id)
			VALUES (DEFAULT, $1, $2)
			RETURNING id;
			`
			err := db.Conn.QueryRow(query, team.Name, userIDOfSession).Scan(&postgresTeamID)
			if err != nil {
				return fmt.Errorf("failed to save the new team: %w", err)
			}
			// Change the ID for the relationship tables
			team.ID = postgresTeamID
		} else {
			// Modify the team if it's not new
			query := `
			UPDATE teams 
			SET name = $1
			WHERE id = $2;
			`
			_, err := db.Conn.Exec(query, team.Name, team.ID)
			if err != nil {
				return fmt.Errorf("failed to save the edited team: %w", err)
			}
		}
	}
	log.Println("Leaving SaveT")
	return nil
}

// SaveClubs saves clubs in the database.
func (db *Database) SaveClubs(clubs map[int]*mt.Club) error {
	for _, club := range clubs {
		if club.ID < 0 {
			// Let postgresql creates its own ID for a new club
			var postgresClubID int
			query := `
			INSERT INTO clubs (id, name, user_id)
			VALUES (DEFAULT, $1, $2)
			RETURNING id;
			`
			err := db.Conn.QueryRow(query, club.Name, userIDOfSession).Scan(&postgresClubID)
			if err != nil {
				return fmt.Errorf("failed to save the new club: %w", err)
			}
			// Change the ID for the relationship tables
			club.ID = postgresClubID
		} else {
			// Modify the club if it's not new
			query := `
			UPDATE clubs 
			SET name = $1
			WHERE id = $2;
			`
			_, err := db.Conn.Exec(query, club.Name, club.ID)
			if err != nil {
				return fmt.Errorf("failed to save the edited club: %w", err)
			}
		}
	}
	return nil
}

// SavePlayerClubs saves the player-club relationships in the database.
func (db *Database) SavePlayerClubs(players map[int]*mt.Player) error {
	for _, player := range players {
		for clubID := range player.ClubIDs {
			log.Println("(SavePC) Player ID is = ", player.ID)
			log.Println("(SavePC) Club ID is = ", clubID)
			query := `
			INSERT INTO player_club (player_id, club_id, user_id)
			VALUES ($1, $2, $3)
			ON CONFLICT (player_id, club_id) DO NOTHING;
			`
			_, err := db.Conn.Exec(query, player.ID, clubID, userIDOfSession)
			if err != nil {
				return fmt.Errorf("failed to save the new player_club relationship: %w", err)
			}
		}
	}
	return nil
}

// SavePlayerTeams saves the player-team relationships in the database.
func (db *Database) SavePlayerTeams(players map[int]*mt.Player) error {
	for _, player := range players {
		for teamID := range player.TeamIDs {
			query := `
			INSERT INTO player_team (player_id, team_id, user_id)
			VALUES ($1, $2, $3)
			ON CONFLICT (player_id, team_id) DO NOTHING;
			`
			_, err := db.Conn.Exec(query, player.ID, teamID, userIDOfSession)
			if err != nil {
				return fmt.Errorf("failed to save the new player_team relationship: %w", err)
			}
		}
	}
	return nil
}

// SaveTeamClubs saves the team-club relationships in the database.
func (db *Database) SaveTeamClubs(teams map[int]*mt.Team) error {
	for _, team := range teams {
		for clubID := range team.ClubID {
			query := `
			INSERT INTO team_club (team_id, club_id, user_id)
			VALUES ($1, $2, $3)
			ON CONFLICT (team_id, club_id) DO NOTHING;
			`
			_, err := db.Conn.Exec(query, team.ID, clubID, userIDOfSession)
			if err != nil {
				return fmt.Errorf("failed to save the new team_club relationship: %w", err)
			}
		}
	}
	return nil
}

func (db *Database) ResetTables() error {
	_, err := db.Conn.Exec(resetTablesQuery, userIDOfSession)
	log.Println("(Reset Table) User ID of the session set to", userIDOfSession)
	if err != nil {
		return fmt.Errorf("failed to reset user data: %w", err)
	}
	return nil
}

// SaveDB saves the database (serialize).
func SaveDB(golangDB *mt.Database) error {
	var err error

	sqlDB, err = ConnectToDB()
	if err != nil {
		fmt.Println("Error while connecting to postgres database:", err)
	}
	sqlDB.ResetTables()

	log.Println("(SaveDb) User ID of the session set to", userIDOfSession)
	log.Println("(SaveDb) Entering SaveU")
	err = sqlDB.SaveUsers(golangDB.Users)
	if err != nil {
		return err
	}
	log.Println("(SaveDb) Entering SaveP")
	err = sqlDB.SavePlayers(golangDB.Players)
	if err != nil {
		return err
	}
	log.Println("(SaveDb) Entering SaveT")
	err = sqlDB.SaveTeams(golangDB.Teams)
	if err != nil {
		return err
	}
	log.Println("(SaveDb) Entering SaveC")
	err = sqlDB.SaveClubs(golangDB.Clubs)
	if err != nil {
		log.Println("(SaveDb SaveC) Error", err)
		return err
	}
	log.Println("(SaveDb) Entering SavePT")
	err = sqlDB.SavePlayerTeams(golangDB.Players)
	if err != nil {
		return err
	}
	log.Println("(SaveDb) Entering SavePC")
	err = sqlDB.SavePlayerClubs(golangDB.Players)
	if err != nil {
		return err
	}
	log.Println("(SaveDb) Entering SaveTC")
	err = sqlDB.SaveTeamClubs(golangDB.Teams)
	if err != nil {
		return err
	}
	log.Println("Database saved successfully.")
	sqlDB.Close()
	return nil
}
