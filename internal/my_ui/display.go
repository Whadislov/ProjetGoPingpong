package myapp

import (
	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// Display launches the UI
func Display(sqlDB *msql.Database, golangDB *mt.Database) {
	StarterPage(sqlDB, golangDB)
}
