package api

import (
	"log"
	"net/http"
)

func RunApi() {
	http.HandleFunc("/users", GetUsers)
	http.HandleFunc("/users/create", CreateUser)
	http.HandleFunc("/users/get", GetUser)
	http.HandleFunc("/users/update", UpdateUser)
	http.HandleFunc("/users/delete", DeleteUser)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
