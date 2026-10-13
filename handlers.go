package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"
)

func listPosts(w http.ResponseWriter, r *http.Request) {

	database, err := sql.Open("sqlite", "store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is ready!")

	if r.URL.Path != "/posts" {
		http.NotFound(w, r)
		return
	}

	formattedJson, err := loadPosts(database)
	if err != nil {
		log.Printf("ERROR: Unable to load tasks from DB: %v", err)
		http.Error(w, "Internal server error: Unable to load data", http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(formattedJson)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	database, err := sql.Open("sqlite", "store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is ready!")

	targetTask, err := findPost(database, id)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, _ := json.Marshal(targetTask)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

type PostInput struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

func createPost(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite", "store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is ready!")

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var input PostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	addPost(database, &input)

	log.Printf("%+v", input)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite", "store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is ready!")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	var input PostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	editPost(database, &input, id)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite", "store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is ready!")
}
