package main

import (
	"database/sql"
	"encoding/json"

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

func loadTasks(db *sql.DB) (*[]Post, error) {
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

func findTask(db *sql.DB, id int) (*Post, error) {

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
