package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	http.HandleFunc("/posts", handlerPosts)
	http.HandleFunc("/posts/{id}", handlerPost)

	log.Println("The server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
