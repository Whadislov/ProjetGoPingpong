package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./wasm")))
	log.Printf("App started on localhost:7001")
	log.Fatal("Server error: ", http.ListenAndServe(":7001", nil))

}
