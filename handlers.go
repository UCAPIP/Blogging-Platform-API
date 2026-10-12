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

	targetTask, err := findTask(database, id)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, _ := json.Marshal(targetTask)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
