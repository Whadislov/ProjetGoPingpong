package mydb

import (
	"fmt"
	"log"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// SavePlayers saves players in the database.
func (db *Database) SaveUsers(users map[int]*mt.User) error {
	log.Println("(SaveU) User ID of the session set to", userIDOfSession)
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
	log.Println("(SaveP) User ID of the session set to", userIDOfSession)
	for _, player := range players {
		query := `
        INSERT INTO players (id, name, age, ranking, forehand, backhand, blade, user_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            age = EXCLUDED.age,
            ranking = EXCLUDED.ranking,
            forehand = EXCLUDED.forehand,
            backhand = EXCLUDED.backhand,
            blade = EXCLUDED.blade
		WHERE players.user_id = EXCLUDED.user_id;
        `
		_, err := db.Conn.Exec(query, player.ID, player.Name, player.Age, player.Ranking, player.Material[0], player.Material[1], player.Material[2], userIDOfSession)
		if err != nil {
			return fmt.Errorf("failed to save player: %w", err)
		}
	}
	return nil
}

// SaveTeams saves teams in the database.
func (db *Database) SaveTeams(teams map[int]*mt.Team) error {
	log.Println("(SaveT) User ID of the session set to", userIDOfSession)
	for _, team := range teams {
		query := `
        INSERT INTO teams (id, name, user_id)
        VALUES ($1, $2, $3)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name
		WHERE teams.user_id = EXCLUDED.user_id;
        `
		_, err := db.Conn.Exec(query, team.ID, team.Name, userIDOfSession)
		if err != nil {
			return fmt.Errorf("failed to save team: %w", err)
		}
	}
	return nil
}

// SaveClubs saves clubs in the database.
func (db *Database) SaveClubs(clubs map[int]*mt.Club) error {
	log.Println("(SaveC) User ID of the session set to", userIDOfSession)
	for _, club := range clubs {
		query := `
        INSERT INTO clubs (id, name, user_id)
        VALUES ($1, $2, $3)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name
		WHERE clubs.user_id = EXCLUDED.user_id;
        `
		_, err := db.Conn.Exec(query, club.ID, club.Name, userIDOfSession)
		if err != nil {
			return fmt.Errorf("failed to save club: %w", err)
		}
	}
	return nil
}

// SavePlayerClubs saves the player-club relationships in the database.
func (db *Database) SavePlayerClubs(players map[int]*mt.Player) error {
	log.Println("(SavePC) User ID of the session set to", userIDOfSession)
	for _, player := range players {
		for clubID := range player.ClubIDs {
			query := `
            INSERT INTO player_club (player_id, club_id, user_id)
            VALUES ($1, $2, $3)
            ON CONFLICT (player_id, club_id) DO NOTHING
			WHERE player_club.user_id = EXCLUDED.user_id;
            `
			_, err := db.Conn.Exec(query, player.ID, clubID, userIDOfSession)
			if err != nil {
				return fmt.Errorf("failed to save player_club relationship: %w", err)
			}
		}
	}
	return nil
}

// SavePlayerTeams saves the player-team relationships in the database.
func (db *Database) SavePlayerTeams(players map[int]*mt.Player) error {
	log.Println("(SavePT) User ID of the session set to", userIDOfSession)
	for _, player := range players {
		for teamID := range player.TeamIDs {
			query := `
            INSERT INTO player_team (player_id, team_id, user_id)
            VALUES ($1, $2, $3)
            ON CONFLICT (player_id, team_id) DO NOTHING
			WHERE player_team.user_id = EXCLUDED.user_id;
            `
			_, err := db.Conn.Exec(query, player.ID, teamID, userIDOfSession)
			if err != nil {
				return fmt.Errorf("failed to save player_team relationship: %w", err)
			}
		}
	}
	return nil
}

// SaveTeamClubs saves the team-club relationships in the database.
func (db *Database) SaveTeamClubs(teams map[int]*mt.Team) error {
	log.Println("(SaveTC) User ID of the session set to", userIDOfSession)
	for _, team := range teams {
		for clubID := range team.ClubID {
			query := `
            INSERT INTO team_club (team_id, club_id, user_id)
            VALUES ($1, $2, $3)
            ON CONFLICT (team_id, club_id) DO NOTHING
			WHERE team_club.user_id = EXCLUDED.user_id;
            `
			_, err := db.Conn.Exec(query, team.ID, clubID, userIDOfSession)
			if err != nil {
				return fmt.Errorf("failed to save team_club relationship: %w", err)
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
	err = sqlDB.SaveUsers(golangDB.Users)
	if err != nil {
		return err
	}
	err = sqlDB.SavePlayers(golangDB.Players)
	if err != nil {
		return err
	}
	err = sqlDB.SaveTeams(golangDB.Teams)
	if err != nil {
		return err
	}
	err = sqlDB.SaveClubs(golangDB.Clubs)
	if err != nil {
		return err
	}
	err = sqlDB.SavePlayerTeams(golangDB.Players)
	if err != nil {
		return err
	}
	err = sqlDB.SavePlayerClubs(golangDB.Players)
	if err != nil {
		return err
	}
	err = sqlDB.SaveTeamClubs(golangDB.Teams)
	if err != nil {
		return err
	}
	log.Println("Database saved successfully.")
	sqlDB.Close()
	return nil
}
