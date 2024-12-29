package mysqlitedb

import (
	"fmt"
	"log"
	"os"
)

func DeleteDB(dbPath string) error {
	err := os.Remove(dbPath)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return nil
	}
	log.Println("Database successfully deleted")
	return nil
}
