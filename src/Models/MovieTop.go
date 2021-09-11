package Models

import "time"

type MovieTop struct {
	ID        int64
	MovieID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
