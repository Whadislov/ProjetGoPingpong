PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE players (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        age INTEGER,
        ranking INTEGER,
        forehand TEXT,
		backhand TEXT,
		blade TEXT
    );
INSERT INTO players VALUES(0,'Julien',27,1632,'Victas V22 Double Extra max','Victas V20 Double Extra 2mm','Koki Niwa Wood');
INSERT INTO players VALUES(1,'Robin',21,1755,'Unknown','Unknown','Unknown');
INSERT INTO players VALUES(2,'Martin',40,1553,'Tibhar Hybrid K3 2mm','Tibhar Hybrid K3 2mm','Tibhar Félix Lebrun Hyper Carbon inner');
INSERT INTO players VALUES(3,'Lasse',20,1189,'Victas V22 Double Extra max','Victas V20 Double Extra max','Koki Niwa Wood');
INSERT INTO players VALUES(4,'Benjamin',40,983,'Unknown','Unknown','Unknown');
INSERT INTO players VALUES(5,'Nicolas',30,747,'Unknown','Unknown','Unknown');
INSERT INTO players VALUES(6,'Arnaud',45,626,'Unknown','Unknown','Unknown');
INSERT INTO players VALUES(7,'Niklas',20,1048,'Tibhar Hybrid MK','Unknown','Unknown');
CREATE TABLE teams (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL
    );
INSERT INTO teams VALUES(0,'Mannschaft 1');
INSERT INTO teams VALUES(1,'Mannschaft 2');
INSERT INTO teams VALUES(2,'Mannschaft 3');
INSERT INTO teams VALUES(3,'Mannschaft 5');
INSERT INTO teams VALUES(4,'D2');
INSERT INTO teams VALUES(5,'D3');
INSERT INTO teams VALUES(6,'D4');
CREATE TABLE clubs (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL
    );
INSERT INTO clubs VALUES(0,'TSG Heilbronn');
INSERT INTO clubs VALUES(1,'SC Turretot');
CREATE TABLE player_club (
		player_id INTEGER,
		club_id INTEGER,
		FOREIGN KEY (player_id) REFERENCES players(id),
		FOREIGN KEY (club_id) REFERENCES clubs(id),
		PRIMARY KEY (player_id, club_id)
	);
INSERT INTO player_club VALUES(6,1);
INSERT INTO player_club VALUES(7,0);
INSERT INTO player_club VALUES(0,1);
INSERT INTO player_club VALUES(0,0);
INSERT INTO player_club VALUES(1,0);
INSERT INTO player_club VALUES(2,0);
INSERT INTO player_club VALUES(3,0);
INSERT INTO player_club VALUES(4,1);
INSERT INTO player_club VALUES(5,1);
CREATE TABLE player_team (
		player_id INTEGER,
		team_id INTEGER,
		FOREIGN KEY (player_id) REFERENCES players(id),
		FOREIGN KEY (team_id) REFERENCES teams(id),
		PRIMARY KEY (player_id, team_id)
	);
INSERT INTO player_team VALUES(6,6);
INSERT INTO player_team VALUES(0,4);
INSERT INTO player_team VALUES(0,1);
INSERT INTO player_team VALUES(0,2);
INSERT INTO player_team VALUES(1,0);
INSERT INTO player_team VALUES(1,1);
INSERT INTO player_team VALUES(2,1);
INSERT INTO player_team VALUES(2,2);
INSERT INTO player_team VALUES(3,3);
INSERT INTO player_team VALUES(4,4);
INSERT INTO player_team VALUES(5,5);
INSERT INTO player_team VALUES(5,4);
CREATE TABLE team_club (
		team_id INTEGER,
		club_id INTEGER,
		FOREIGN KEY (team_id) REFERENCES teams(id),
		FOREIGN KEY (club_id) REFERENCES clubs(id),
		PRIMARY KEY (team_id, club_id)
	);
INSERT INTO team_club VALUES(6,1);
INSERT INTO team_club VALUES(0,0);
INSERT INTO team_club VALUES(1,0);
INSERT INTO team_club VALUES(2,0);
INSERT INTO team_club VALUES(3,0);
INSERT INTO team_club VALUES(4,1);
INSERT INTO team_club VALUES(5,1);
COMMIT;
