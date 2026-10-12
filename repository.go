package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Tags      string `json:"tags"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

//func loadTasks() ([]Task, error) {
//	database, err := sql.Open("sqlite3", "./store.db")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer database.Close()
//
//	if err = database.Ping(); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("DB is ready!")
//}

func findTask(idQ string) string {

	database, err := sql.Open("sqlite", "store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is ready!")
	var article Task
	stringQuery := "SELECT id, title, content, category, tags, created_at, updated_at FROM posts WHERE id = " + idQ
	err = database.QueryRow(stringQuery).Scan(&article.Id, &article.Title, &article.Content, &article.Category, &article.Tags, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Задача не найдена")
		} else {
			log.Fatal(err)
		}
	}

	jsonData, err := json.Marshal(article)
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}
