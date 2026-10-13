package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

type Post struct {
	Id        int      `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

func loadPosts(db *sql.DB) (*[]Post, error) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	var posts []Post
	var tagsRaw string

	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.Id, &p.Title, &p.Content, &p.Category, &tagsRaw, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(tagsRaw), &p.Tags); err != nil {
			return nil, err
		}
		if p.Tags == nil {
			p.Tags = []string{}
		}
		posts = append(posts, p)
	}
	rows.Err()
	return &posts, nil
}

func findPost(db *sql.DB, id int) (*Post, error) {

	var p Post
	var tagsRaw string

	err := db.QueryRow(
		"SELECT id, title, content, category, tags, created_at, updated_at FROM posts WHERE id = ?",
		id,
	).Scan(&p.Id, &p.Title, &p.Content, &p.Category, &tagsRaw, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(tagsRaw), &p.Tags); err != nil {
		return nil, err
	}
	if p.Tags == nil {
		p.Tags = []string{}
	}

	return &p, nil
}

func addPost(db *sql.DB, pi *PostInput) {

	now := time.Now().Format(time.RFC3339)

	if pi.Tags == nil {
		pi.Tags = []string{}
	}
	tagsJSON, err := json.Marshal(pi.Tags)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := db.Exec("INSERT INTO posts (title, content, category, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		pi.Title, pi.Content, pi.Category, string(tagsJSON), now, now,
	)

	if err != nil {
		log.Println(err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(id, "ID Post was created succesfuly")

}
