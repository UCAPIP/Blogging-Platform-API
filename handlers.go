package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

func (a *App) listPosts(w http.ResponseWriter, r *http.Request) {

	formattedJson, err := loadPosts(a.db)
	if err != nil {
		log.Printf("ERROR: Unable to load tasks from DB: %v", err)
		writeError(w, http.StatusInternalServerError, "Internal server error", "Unable to load data.")
		return
	}

	jsonData, err := json.Marshal(formattedJson)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "Internal server error", "")
		return
	}
	writeJSON(w, http.StatusOK, jsonData)
}

func (a *App) getPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Bad Request", "The 'id' parameter must be a valid integer number.")
		return
	}

	targetTask, err := findPost(a.db, id)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "Not Found", "The requested post could not be found.")
		return
	}
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "Internal Server Error", "An unexpected error occurred on the server. Please try again later.")
		return
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

func validate(input *PostInput) []string {
	var problems []string
	if strings.TrimSpace(input.Title) == "" {
		problems = append(problems, "title is required")
	}
	if strings.TrimSpace(input.Content) == "" {
		problems = append(problems, "content is required")
	}
	if strings.TrimSpace(input.Category) == "" {
		problems = append(problems, "category is required")
	}
	return problems
}

func (a *App) createPost(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var input PostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON body", "")
		return
	}

	if problems := validate(&input); len(problems) > 0 {
		writeError(w, http.StatusBadRequest, "validation failed", strings.Join(problems, "; "))
		return
	}

	id, err := addPost(a.db, &input)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "Internal Server Error", "An unexpected error occurred on the server. Please try again later.")
		return
	}

	post, err := findPost(a.db, int(id))
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "internal error", "")
		return
	}

	writeJSON(w, http.StatusCreated, post)
}

func (a *App) updatePost(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Bad Request", "The 'id' parameter must be a valid integer number.")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var input PostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body", "")
		return
	}

	if problems := validate(&input); len(problems) > 0 {
		writeError(w, http.StatusBadRequest, "validation failed", strings.Join(problems, "; "))
		return
	}

	err = editPost(a.db, &input, id)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "internal error", "")
		return
	}

	post, err := findPost(a.db, id)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "internal error", "")
		return
	}

	writeJSON(w, http.StatusOK, post)
}

func (a *App) deletePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Bad Request", "The 'id' parameter must be a valid integer number.")
		return
	}

	err = deleting(a.db, id)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "internal error", "")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
