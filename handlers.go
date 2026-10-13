package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"
)

func (a *App) listPosts(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/posts" {
		http.NotFound(w, r)
		return
	}

	formattedJson, err := loadPosts(a.db)
	if err != nil {
		log.Printf("ERROR: Unable to load tasks from DB: %v", err)
		http.Error(w, "Internal server error: Unable to load data", http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(formattedJson)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (a *App) getPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	targetTask, err := findPost(a.db, id)
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

func (a *App) createPost(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var input PostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	addPost(a.db, &input)

	log.Printf("%+v", input)
}

func (a *App) updatePost(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}
	var input PostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	editPost(a.db, &input, id)
}

func (a *App) deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("is ready!")
}
