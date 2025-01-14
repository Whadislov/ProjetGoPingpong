package mysqlitedb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Database represents the SQLite database connection.
type Database struct {
	Conn *sql.DB
}

// ConnectToDB initializes a connection to the SQLite database and creates it if it does not exist.
func ConnectToDB(dbType string) (*Database, error) {
	var conn *sql.DB
	var err error

	if dbType == "sqlite" {

		// Check if the database file exists
		if _, err := os.Stat(sqliteDbPath); os.IsNotExist(err) {
			// Create the database file
			file, err := os.Create(sqliteDbPath)
			if err != nil {
				return nil, fmt.Errorf("failed to create database file: %w", err)
			}
			file.Close()
			log.Println("Database file created successfully.")
		}

		// Connect to sqlite DB
		conn, err = sql.Open("sqlite3", sqliteDbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to sqlite database: %w", err)
		}

	} else if dbType == "postgres" {
		// Connect to postgres DB
		conn, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to postgres database: %w", err)
		}
	}

	// Verify the connection
	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to verify connection to database: %w", err)
	}

	log.Printf("Connected to %v database successfully.", dbType)
	db := &Database{Conn: conn}

	// Create tables if they do not exist
	err = db.CreateTables()
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

// CreateTables creates the necessary tables in the database.
func (db *Database) CreateTables() error {
	createPlayersTable := `
    CREATE TABLE IF NOT EXISTS players (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        age INTEGER,
        ranking INTEGER,
        forehand TEXT,
		backhand TEXT,
		blade TEXT
    );`

	createTeamsTable := `
    CREATE TABLE IF NOT EXISTS teams (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL
    );`

	createClubsTable := `
    CREATE TABLE IF NOT EXISTS clubs (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL
    );`

	// Table relation for players and clubs
	createPlayerClubTable := `
	CREATE TABLE IF NOT EXISTS player_club (
		player_id INTEGER,
		club_id INTEGER,
		FOREIGN KEY (player_id) REFERENCES players(id),
		FOREIGN KEY (club_id) REFERENCES clubs(id),
		PRIMARY KEY (player_id, club_id)
	);`

	// Table relation for players and teams
	createPlayerTeamTable := `
	CREATE TABLE IF NOT EXISTS player_team (
		player_id INTEGER,
		team_id INTEGER,
		FOREIGN KEY (player_id) REFERENCES players(id),
		FOREIGN KEY (team_id) REFERENCES teams(id),
		PRIMARY KEY (player_id, team_id)
	);`

	// Table relation for clubs and teams
	createTeamClubTable := `
	CREATE TABLE IF NOT EXISTS team_club (
		team_id INTEGER,
		club_id INTEGER,
		FOREIGN KEY (team_id) REFERENCES teams(id),
		FOREIGN KEY (club_id) REFERENCES clubs(id),
		PRIMARY KEY (team_id, club_id)
	);`

	_, err := db.Conn.Exec(createPlayersTable)
	if err != nil {
		return fmt.Errorf("failed to create players table: %w", err)
	}

	_, err = db.Conn.Exec(createTeamsTable)
	if err != nil {
		return fmt.Errorf("failed to create teams table: %w", err)
	}

	_, err = db.Conn.Exec(createClubsTable)
	if err != nil {
		return fmt.Errorf("failed to create clubs table: %w", err)
	}

	_, err = db.Conn.Exec(createPlayerClubTable)
	if err != nil {
		return fmt.Errorf("failed to create player_club table: %w", err)
	}

	_, err = db.Conn.Exec(createPlayerTeamTable)
	if err != nil {
		return fmt.Errorf("failed to create player_team table: %w", err)
	}

	_, err = db.Conn.Exec(createTeamClubTable)
	if err != nil {
		return fmt.Errorf("failed to create team_club table: %w", err)
	}

	log.Println("Tables created successfully.")
	return nil
}

// Close closes the database connection.
func (db *Database) Close() error {
	err := db.Conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	log.Println("Database connection closed successfully.")
	return nil
}
