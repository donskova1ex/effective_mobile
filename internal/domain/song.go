package domain

import "time"

type Song struct {
	ID          int32     `json:"id" db:"id"`
	GroupName   string    `json:"group_name" db:"group_name"`
	SongName    string    `json:"song_name" db:"song_name"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Text        string    `json:"text" db:"text"`
	Link        string    `json:"link" db:"link"`
}
