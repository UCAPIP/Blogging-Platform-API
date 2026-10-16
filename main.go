package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

type App struct {
	db *sql.DB
}

func main() {

	db, err := sql.Open("sqlite", "store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	app := &App{db: db}

	http.HandleFunc("GET /posts", app.listPosts)
	http.HandleFunc("GET /posts/{id}", app.getPost)
	http.HandleFunc("POST /posts", app.createPost)
	http.HandleFunc("PUT /posts/{id}", app.updatePost)
	http.HandleFunc("DELETE /posts/{id}", app.deletePost)

	log.Println("The server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
