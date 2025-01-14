package mysqlitedb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Database represents the SQLite database connection.
type Database struct {
	Conn *sql.DB
}

// ConnectToDB initializes a connection to the SQLite database and creates it if it does not exist.
func ConnectToDB() (*Database, error) {
	var conn *sql.DB
	var err error

	// Connect to postgres DB
	conn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres database: %w", err)
	}

	// Check if the database exists
	query := `SELECT 1 FROM pg_database WHERE datname = $1`
	var exists int
	err = conn.QueryRow(query, dbName).Scan(&exists)
	if err == sql.ErrNoRows {
		// Database does not exist, create it
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		_, err = conn.Exec(createDBQuery)
		if err != nil {
			return nil, fmt.Errorf("failed to create database: %w", err)
		}
		log.Printf("Database %s created successfully.\n", dbName)
	} else if err != nil {
		return nil, fmt.Errorf("error checking database existence: %w", err)
	}
	// Verify the connection
	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to verify connection to database: %w", err)
	}

	log.Printf("Connected to database successfully.")
	db := &Database{Conn: conn}

	// Create tables if they do not exist
	err = db.CreateTables()
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

// CreateTables creates the necessary tables in the database for PostgreSQL.
func (db *Database) CreateTables() error {
	createPlayersTable := `
    CREATE TABLE IF NOT EXISTS players (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        age INTEGER,
        ranking INTEGER,
        forehand TEXT,
        backhand TEXT,
        blade TEXT
    );`

	createTeamsTable := `
    CREATE TABLE IF NOT EXISTS teams (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL
    );`

	createClubsTable := `
    CREATE TABLE IF NOT EXISTS clubs (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL
    );`

	// Table relation for players and clubs
	createPlayerClubTable := `
    CREATE TABLE IF NOT EXISTS player_club (
        player_id INTEGER NOT NULL,
        club_id INTEGER NOT NULL,
        PRIMARY KEY (player_id, club_id),
        FOREIGN KEY (player_id) REFERENCES players(id),
        FOREIGN KEY (club_id) REFERENCES clubs(id)
    );`

	// Table relation for players and teams
	createPlayerTeamTable := `
    CREATE TABLE IF NOT EXISTS player_team (
        player_id INTEGER NOT NULL,
        team_id INTEGER NOT NULL,
        PRIMARY KEY (player_id, team_id),
        FOREIGN KEY (player_id) REFERENCES players(id),
        FOREIGN KEY (team_id) REFERENCES teams(id)
    );`

	// Table relation for clubs and teams
	createTeamClubTable := `
    CREATE TABLE IF NOT EXISTS team_club (
        team_id INTEGER NOT NULL,
        club_id INTEGER NOT NULL,
        PRIMARY KEY (team_id, club_id),
        FOREIGN KEY (team_id) REFERENCES teams(id),
        FOREIGN KEY (club_id) REFERENCES clubs(id)
    );`

	// Execute the table creation queries
	queries := []string{
		createPlayersTable,
		createTeamsTable,
		createClubsTable,
		createPlayerClubTable,
		createPlayerTeamTable,
		createTeamClubTable,
	}

	for _, query := range queries {
		_, err := db.Conn.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to execute query: %w", err)
		}
	}
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
