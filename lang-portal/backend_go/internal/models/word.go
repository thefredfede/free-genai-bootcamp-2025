package models

type Word struct {
	ID       int      `json:"id"`
	Japanese string   `json:"japanese"`
	Romaji   string   `json:"romaji"`
	English  string   `json:"english"`
	Parts    []string `json:"parts"`
}
