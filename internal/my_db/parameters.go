package mydb

var sqlDB *Database
var userIDOfSession int

func SetUserIDOfSession(id int) {
	userIDOfSession = id
}

var psqlInfo string

func SetPsqlInfo(link string) {
	psqlInfo = link
}

var dbName string

func SetDBName(name string) {
	dbName = name
}

// Query script for table creation (split in two parts users + other tables)
// player_club = table relation for players and clubs
// player_team = table relation for players and teams
// team_club = table relation for teams and clubs

var createUserTableQuery string = `
BEGIN;
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    created_at TIMESTAMPTZ
);
`

// The 3 lines SELECT setval are used to synchronise the autoincrement
var createOtherTablesQuery string = `

CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
	firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
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

SELECT setval('players_id_seq', (SELECT MAX(id) FROM players));
SELECT setval('teams_id_seq', (SELECT MAX(id) FROM teams));
SELECT setval('clubs_id_seq', (SELECT MAX(id) FROM clubs));

COMMIT;`

var createAllTablesQuery string = createUserTableQuery + createOtherTablesQuery

// Query script for table reset because we can't delete elements from the database directly
// $1 is the user_id
var resetTablesQuery string = `
BEGIN;

DELETE FROM players WHERE user_id = $1;
DELETE FROM teams WHERE user_id = $1;
DELETE FROM clubs WHERE user_id = $1;
DELETE FROM player_club WHERE user_id = $1;
DELETE FROM player_team WHERE user_id = $1;
DELETE FROM team_club WHERE user_id = $1;
` + createOtherTablesQuery
