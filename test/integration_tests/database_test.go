package integrationtests

import (
	"database/sql"
	"log"
	"testing"

	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	_ "github.com/mattn/go-sqlite3"
)

// TestDatabaseConnection tests the database connection and table creation.
func TestDatabaseConnection(t *testing.T) {
	// Connect to the database
	db, err := msql.ConnectToDB("mockDB.db")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Conn.Close()

	// Verify the connection
	err = db.Conn.Ping()
	if err != nil {
		t.Fatalf("Failed to verify connection to database: %v", err)
	}

	// Verify that the tables are created
	tables := []string{"players", "teams", "clubs", "player_club"}
	for _, table := range tables {
		if !tableExists(db.Conn, table) {
			t.Fatalf("Table %s does not exist", table)
		}
	}
}

// tableExists checks if a table exists in the database.
func tableExists(db *sql.DB, tableName string) bool {
	query := `SELECT name FROM sqlite_master WHERE type='table' AND name=?;`
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatalf("Failed to check if table %s exists: %v", tableName, err)
	}
	return name == tableName
}
