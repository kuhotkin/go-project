package models

type project struct {
	ID    string `json:"id" db:"id"`
	Title string `json:"name" db:"title"`
}
