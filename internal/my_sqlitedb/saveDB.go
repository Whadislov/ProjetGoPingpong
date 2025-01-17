package mysqlitedb

import (
	"fmt"
	"log"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// SavePlayers saves players in the database.
func (db *Database) SavePlayers(players map[int]*mt.Player) error {
	for _, player := range players {
		query := `
        INSERT INTO players (id, name, age, ranking, forehand, backhand, blade)
        VALUES (?, ?, ?, ?, ?, ?, ?)
        ON CONFLICT(id) DO UPDATE SET
            name=excluded.name,
            age=excluded.age,
            ranking=excluded.ranking,
            forehand=excluded.forehand,
            backhand=excluded.backhand,
            blade=excluded.blade;
        `
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
		query := `
        INSERT INTO teams (id, name)
        VALUES (?, ?)
        ON CONFLICT(id) DO UPDATE SET
            name=excluded.name;
        `
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
		query := `
        INSERT INTO clubs (id, name)
        VALUES (?, ?)
        ON CONFLICT(id) DO UPDATE SET
            name=excluded.name;
        `
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
			query := `
            INSERT OR REPLACE INTO player_club (player_id, club_id)
            VALUES (?, ?);
            `
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
			query := `
            INSERT OR REPLACE INTO player_team (player_id, team_id)
            VALUES (?, ?);
            `
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
			query := `
            INSERT OR REPLACE INTO team_club (team_id, club_id)
            VALUES (?, ?);
            `
			_, err := db.Conn.Exec(query, team.ID, clubID)
			if err != nil {
				return fmt.Errorf("failed to save team_club relationship: %w", err)
			}
		}
	}
	return nil
}

// SaveDB saves the database. (serialize)
func SaveDB(golangDB *mt.Database) error {
	var err error

	// We can save multiple times, but we want this part of code to be executed only once. This bloc sets up a new database
	initOnce.Do(func() {
		log.Println("Database will be cleared and reloaded.")
		DeleteDB(DbPath)
		sqlDB, err = ConnectToDB(DbPath)
		if err != nil {
			fmt.Println("Error while connecting to sql database:", err)
		}
	})

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
	log.Println("Database successfully saved.")
	sqlDB.Close()
	return nil
}
