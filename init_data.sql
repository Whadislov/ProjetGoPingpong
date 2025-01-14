BEGIN;

DROP TABLE IF EXISTS player_team CASCADE;
DROP TABLE IF EXISTS player_club CASCADE;
DROP TABLE IF EXISTS team_club CASCADE;
DROP TABLE IF EXISTS players CASCADE;
DROP TABLE IF EXISTS teams CASCADE;
DROP TABLE IF EXISTS clubs CASCADE;

CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    age INTEGER,
    ranking INTEGER,
    forehand VARCHAR,
    backhand VARCHAR,
    blade VARCHAR
);

INSERT INTO players (id, name, age, ranking, forehand, backhand, blade) VALUES
(0, 'Julien', 27, 1632, 'Victas V22 Double Extra max', 'Victas V20 Double Extra 2mm', 'Koki Niwa Wood'),
(1, 'Robin', 21, 1755, 'Unknown', 'Unknown', 'Unknown'),
(2, 'Martin', 40, 1553, 'Tibhar Hybrid K3 2mm', 'Tibhar Hybrid K3 2mm', 'Tibhar Félix Lebrun Hyper Carbon inner'),
(3, 'Lasse', 20, 1189, 'Victas V22 Double Extra max', 'Victas V20 Double Extra max', 'Koki Niwa Wood'),
(4, 'Benjamin', 40, 983, 'Unknown', 'Unknown', 'Unknown'),
(5, 'Nicolas', 30, 747, 'Unknown', 'Unknown', 'Unknown'),
(6, 'Arnaud', 45, 626, 'Unknown', 'Unknown', 'Unknown'),
(7, 'Niklas', 20, 1048, 'Tibhar Hybrid MK', 'Unknown', 'Unknown');

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

INSERT INTO teams (id, name) VALUES
(0, 'Mannschaft 1'),
(1, 'Mannschaft 2'),
(2, 'Mannschaft 3'),
(3, 'Mannschaft 5'),
(4, 'D2'),
(5, 'D3'),
(6, 'D4');

CREATE TABLE clubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

INSERT INTO clubs (id, name) VALUES
(0, 'TSG Heilbronn'),
(1, 'SC Turretot');

CREATE TABLE player_club (
    player_id INTEGER,
    club_id INTEGER,
    PRIMARY KEY (player_id, club_id),
    CONSTRAINT fk_player FOREIGN KEY (player_id) REFERENCES players (id) ON DELETE CASCADE,
    CONSTRAINT fk_club FOREIGN KEY (club_id) REFERENCES clubs (id) ON DELETE CASCADE
);

INSERT INTO player_club (player_id, club_id) VALUES
(6, 1),
(7, 0),
(0, 1),
(0, 0),
(1, 0),
(2, 0),
(3, 0),
(4, 1),
(5, 1);

CREATE TABLE player_team (
    player_id INTEGER,
    team_id INTEGER,
    PRIMARY KEY (player_id, team_id),
    CONSTRAINT fk_player_team_player FOREIGN KEY (player_id) REFERENCES players (id) ON DELETE CASCADE,
    CONSTRAINT fk_player_team_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE
);

INSERT INTO player_team (player_id, team_id) VALUES
(6, 6),
(0, 4),
(0, 1),
(0, 2),
(1, 0),
(1, 1),
(2, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(5, 4);

CREATE TABLE team_club (
    team_id INTEGER,
    club_id INTEGER,
    PRIMARY KEY (team_id, club_id),
    CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE,
    CONSTRAINT fk_team_club FOREIGN KEY (club_id) REFERENCES clubs (id) ON DELETE CASCADE
);

INSERT INTO team_club (team_id, club_id) VALUES
(6, 1),
(0, 0),
(1, 0),
(2, 0),
(3, 0),
(4, 1),
(5, 1);

COMMIT;