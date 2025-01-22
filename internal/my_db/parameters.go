package mydb

import (
	"fmt"
)

var sqlDB *Database
var user_id string

// Const for PostgreSQL
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "p1Sw"
	dbName   = "ttapp_database"
)

var psqlInfo string

func AppStartOption(s string) {
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
CREATE TABLE IF NOT users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    created_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER,
    ranking INTEGER,
    forehand TEXT,
    backhand TEXT,
    blade TEXT,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS teams (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS clubs (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS player_club (
	player_id INTEGER NOT NULL,
	club_id INTEGER NOT NULL,
	PRIMARY KEY (player_id, club_id),
	FOREIGN KEY (player_id) REFERENCES players(id),
	FOREIGN KEY (club_id) REFERENCES clubs(id),
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS player_team (
	player_id INTEGER NOT NULL,
	team_id INTEGER NOT NULL,
	PRIMARY KEY (player_id, team_id),
	FOREIGN KEY (player_id) REFERENCES players(id),
	FOREIGN KEY (team_id) REFERENCES teams(id),
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS team_club (
	team_id INTEGER NOT NULL,
	club_id INTEGER NOT NULL,
	PRIMARY KEY (team_id, club_id),
	FOREIGN KEY (team_id) REFERENCES teams(id),
	FOREIGN KEY (club_id) REFERENCES clubs(id),
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

COMMIT;`

// Query script for table reset because we can't delete elements from the database directly
// Replace "BEGIN;" (createTablesQuery[6:]) from the other script with this part
var resetTablesQuery string = `
BEGIN;

TRUNCATE TABLE users CASCADE;
TRUNCATE TABLE player_team CASCADE;
TRUNCATE TABLE player_club CASCADE;
TRUNCATE TABLE team_club CASCADE;
TRUNCATE TABLE players CASCADE;
TRUNCATE TABLE teams CASCADE;
TRUNCATE TABLE clubs CASCADE;
` + createTablesQuery[6:]
