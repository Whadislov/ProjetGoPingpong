package mysqlitedb

import (
	"sync"
)

var initOnce sync.Once
var DbPath string = "database.db"
var sqlDB *Database
