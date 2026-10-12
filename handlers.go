package main

import (
	"encoding/json"
	"net/http"
)

//func handlerPosts(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/posts" {
//		http.NotFound(w, r)
//		return
//	}
//
//	formattedJson, err := loadTasks()
//	if err != nil {
//		log.Printf("ERROR: Unable to load tasks from DB: %v", err)
//		http.Error(w, "Internal server error: Unable to load data", http.StatusInternalServerError)
//		return
//	}
//
//	jsonData, _ := json.Marshal(formattedJson)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(jsonData)
//}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	targetTask := findTask(id)

	jsonData, _ := json.Marshal(targetTask)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
