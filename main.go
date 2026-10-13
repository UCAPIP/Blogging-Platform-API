package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	http.HandleFunc("GET /posts", listPosts)
	http.HandleFunc("GET /posts/{id}", getPost)
	http.HandleFunc("POST /posts", createPost)
	http.HandleFunc("PUT /posts/{id}", updatePost)
	http.HandleFunc("DELETE /posts/{id}", deletePost)

	log.Println("The server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
