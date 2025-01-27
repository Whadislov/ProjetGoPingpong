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
		query := fmt.Sprintf(`
        INSERT INTO players (id, name, age, ranking, forehand, backhand, blade, user_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, %v)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            age = EXCLUDED.age,
            ranking = EXCLUDED.ranking,
            forehand = EXCLUDED.forehand,
            backhand = EXCLUDED.backhand,
            blade = EXCLUDED.blade;
        `, userOfSession.ID)
		_, err := db.Conn.Exec(query, player.ID, player.Name, player.Age, player.Ranking, player.Material[0], player.Material[1], player.Material[2])
		if err != nil {
			return fmt.Errorf("failed to save player: %w", err)
		}
	}
	return nil
}

// SaveTeams saves teams in the database.
func (db *Database) SaveTeams(teams map[int]*mt.Team) error {
	for _, team := range teams {
		query := fmt.Sprintf(`
        INSERT INTO teams (id, name, user_id)
        VALUES ($1, $2, %v)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name;
        `, userOfSession.ID)
		_, err := db.Conn.Exec(query, team.ID, team.Name)
		if err != nil {
			return fmt.Errorf("failed to save team: %w", err)
		}
	}
	return nil
}

// SaveClubs saves clubs in the database.
func (db *Database) SaveClubs(clubs map[int]*mt.Club) error {
	for _, club := range clubs {
		query := fmt.Sprintf(`
        INSERT INTO clubs (id, name, user_id)
        VALUES ($1, $2, %v)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name;
        `, userOfSession.ID)
		_, err := db.Conn.Exec(query, club.ID, club.Name)
		if err != nil {
			return fmt.Errorf("failed to save club: %w", err)
		}
	}
	return nil
}

// SavePlayerClubs saves the player-club relationships in the database.
func (db *Database) SavePlayerClubs(players map[int]*mt.Player) error {
	for _, player := range players {
		for clubID := range player.ClubIDs {
			query := fmt.Sprintf(`
            INSERT INTO player_club (player_id, club_id, user_id)
            VALUES ($1, $2, %v)
            ON CONFLICT (player_id, club_id) DO NOTHING;
            `, userOfSession.ID)
			_, err := db.Conn.Exec(query, player.ID, clubID)
			if err != nil {
				return fmt.Errorf("failed to save player_club relationship: %w", err)
			}
		}
	}
	return nil
}

// SavePlayerTeams saves the player-team relationships in the database.
func (db *Database) SavePlayerTeams(players map[int]*mt.Player) error {
	for _, player := range players {
		for teamID := range player.TeamIDs {
			query := fmt.Sprintf(`
            INSERT INTO player_team (player_id, team_id, user_id)
            VALUES ($1, $2, %v)
            ON CONFLICT (player_id, team_id) DO NOTHING;
            `, userOfSession.ID)
			_, err := db.Conn.Exec(query, player.ID, teamID)
			if err != nil {
				return fmt.Errorf("failed to save player_team relationship: %w", err)
			}
		}
	}
	return nil
}

// SaveTeamClubs saves the team-club relationships in the database.
func (db *Database) SaveTeamClubs(teams map[int]*mt.Team) error {
	for _, team := range teams {
		for clubID := range team.ClubID {
			query := fmt.Sprintf(`
            INSERT INTO team_club (team_id, club_id, user_id)
            VALUES ($1, $2, %v)
            ON CONFLICT (team_id, club_id) DO NOTHING;
            `, userOfSession.ID)
			_, err := db.Conn.Exec(query, team.ID, clubID)
			if err != nil {
				return fmt.Errorf("failed to save team_club relationship: %w", err)
			}
		}
	}
	return nil
}

func (db *Database) ResetTables() error {
	_, err := db.Conn.Exec(resetTablesQuery)
	if err != nil {
		return fmt.Errorf("failed to reset database: %w", err)
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
