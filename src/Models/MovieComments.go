package Models

import "time"

type MovieComments struct {
	ID           int64
	UserID       int64
	MoviesID     int64
	ToCommentsID int64
	Content      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
