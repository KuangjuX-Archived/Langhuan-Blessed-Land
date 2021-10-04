package Models

import "time"

type TagMovieMap struct {
	ID        int64
	TagID     int64
	MovieID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
