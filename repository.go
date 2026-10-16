package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"time"
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

var ErrNotFound = errors.New("post not found")

func loadPosts(db *sql.DB) ([]Post, error) {

	posts := []Post{}

	rows, err := db.Query("SELECT id, title, content, category, tags, created_at, updated_at FROM posts")
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Post
		var tagsRaw string
		if err := rows.Scan(&p.Id, &p.Title, &p.Content, &p.Category, &tagsRaw, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return posts, err
		}
		if err := json.Unmarshal([]byte(tagsRaw), &p.Tags); err != nil {
			return posts, err
		}
		if p.Tags == nil {
			p.Tags = []string{}
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func findPost(db *sql.DB, id int) (*Post, error) {
	var p Post
	var tagsRaw string

	err := db.QueryRow(
		"SELECT id, title, content, category, tags, created_at, updated_at FROM posts WHERE id = ?",
		id,
	).Scan(&p.Id, &p.Title, &p.Content, &p.Category, &tagsRaw, &p.CreatedAt, &p.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
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

func addPost(db *sql.DB, input *PostInput) (int64, error) {

	now := time.Now().Format(time.RFC3339)

	if input.Tags == nil {
		input.Tags = []string{}
	}
	tagsJSON, err := json.Marshal(input.Tags)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	res, err := db.Exec("INSERT INTO posts (title, content, category, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		input.Title, input.Content, input.Category, string(tagsJSON), now, now,
	)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	log.Println(id, "ID Post was created succesfuly")
	return id, nil

}

func editPost(db *sql.DB, input *PostInput, id int) error {
	now := time.Now().Format(time.RFC3339)

	if input.Tags == nil {
		input.Tags = []string{}
	}
	tagsJSON, err := json.Marshal(input.Tags)
	if err != nil {
		return err
	}

	res, err := db.Exec("UPDATE posts SET title = ?, content = ?, category = ?, tags = ?, updated_at = ? WHERE id = ?",
		input.Title, input.Content, input.Category, string(tagsJSON), now, id,
	)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

func deleting(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}
