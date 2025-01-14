package mysqlitedb

import (
	"fmt"
	"sync"
)

var initOnce sync.Once
var sqliteDbPath string = "database.db"
var sqlDB *Database

// Const for PostgreSQL
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "wn7-00407"
	dbname   = "ttapp_database"
)

// PostgreSQL info
var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
