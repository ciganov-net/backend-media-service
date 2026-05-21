package models

import "time"

type File struct {
	ID        string    `json:"id"`
	Filename  string    `json:"filename"`
	ObjectKey string    `json:"object_key"`
	Size      int64     `json:"size"`
	MimeType  string    `json:"mime_type"`
	CreatedAt time.Time `json:"created_at"`
}
