package mysqlitedb

import (
	"fmt"
	//"sync"
)

// var initOnce sync.Once
var sqlDB *Database

// Const for PostgreSQL
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "wn7-00407"
	dbName   = "ttapp_database"
)

// PostgreSQL info
var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbName)

// Reset script (because we can't delete things from the database)
var resetScript string = `
BEGIN;

DROP TABLE IF EXISTS player_team CASCADE;
DROP TABLE IF EXISTS player_club CASCADE;
DROP TABLE IF EXISTS team_club CASCADE;
DROP TABLE IF EXISTS players CASCADE;
DROP TABLE IF EXISTS teams CASCADE;
DROP TABLE IF EXISTS clubs CASCADE;

CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER,
    ranking INTEGER,
    forehand TEXT,
    backhand TEXT,
    blade TEXT
);

CREATE TABLE teams (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE clubs (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE player_club (
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
