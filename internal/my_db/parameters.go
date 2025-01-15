package mydb

import (
	"fmt"
)

var sqlDB *Database

// Const for PostgreSQL
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "wn7-00407"
	dbName   = "ttapp_database"
)

var psqlInfo string

func StartOption(s string) {
	if s == "local" {
		// PostgreSQL info
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbName)
	} else if s == "browser" {
		// Neon server
		psqlInfo = "postgresql://ttapp_database_owner:7MopfqD4SIyh@ep-white-unit-a2ap77if.eu-central-1.aws.neon.tech/ttapp_database?sslmode=require"
	}
}

// Query script for table creation
// player_club = table relation for players and clubs
// player_team = table relation for players and teams
// team_club = table relation for teams and clubs

var createTablesQuery string = `
BEGIN;

CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER,
    ranking INTEGER,
    forehand TEXT,
    backhand TEXT,
    blade TEXT
);

CREATE TABLE IF NOT EXISTS teams (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS clubs (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS player_club (
	player_id INTEGER NOT NULL,
	club_id INTEGER NOT NULL,
	PRIMARY KEY (player_id, club_id),
	FOREIGN KEY (player_id) REFERENCES players(id),
	FOREIGN KEY (club_id) REFERENCES clubs(id)
);

CREATE TABLE IF NOT EXISTS player_team (
	player_id INTEGER NOT NULL,
	team_id INTEGER NOT NULL,
	PRIMARY KEY (player_id, team_id),
	FOREIGN KEY (player_id) REFERENCES players(id),
	FOREIGN KEY (team_id) REFERENCES teams(id)
);

CREATE TABLE IF NOT EXISTS team_club (
	team_id INTEGER NOT NULL,
	club_id INTEGER NOT NULL,
	PRIMARY KEY (team_id, club_id),
	FOREIGN KEY (team_id) REFERENCES teams(id),
	FOREIGN KEY (club_id) REFERENCES clubs(id)
);

COMMIT;`

// Query script for table reset because we can't delete elements from the database directly
// Remove "BEGIN;" from the other script
var resetTablesQuery string = `
BEGIN;

DROP TABLE IF EXISTS player_team CASCADE;
DROP TABLE IF EXISTS player_club CASCADE;
DROP TABLE IF EXISTS team_club CASCADE;
DROP TABLE IF EXISTS players CASCADE;
DROP TABLE IF EXISTS teams CASCADE;
DROP TABLE IF EXISTS clubs CASCADE;
` + createTablesQuery[6:]
