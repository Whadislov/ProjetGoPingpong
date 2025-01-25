package my_types

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    string `json:"created_at"`
}

type Club struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"players"`
	TeamIDs   map[int]string `json:"teams"`
}

type Player struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Age      int            `json:"age"`
	Ranking  int            `json:"ranking"`
	Material []string       `json:"material"`
	TeamIDs  map[int]string `json:"teams"`
	ClubIDs  map[int]string `json:"clubs"`
}

type Team struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"players"`
	ClubID    map[int]string `json:"clubs"`
}

type Database struct {
	Users   map[int]*User   `json:"users"`
	Clubs   map[int]*Club   `json:"clubs"`
	Teams   map[int]*Team   `json:"teams"`
	Players map[int]*Player `json:"players"`
}
