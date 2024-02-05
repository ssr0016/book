package model

import "fmt"

var (
	ErrIDNotFound = fmt.Errorf("id not found")
)

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}
