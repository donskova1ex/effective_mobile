package domain

import "time"

type Song struct {
	id           int32     `json:"id" db:"id"`
	group_name   string    `json:"group_name" db:"group_name"`
	song_name    string    `json:"song_name" db:"song_name"`
	release_date time.Time `json:"release_date" db:"release_date"`
	text         string    `json:"text" db:"text"`
	link         string    `json:"link" db:"link"`
}
