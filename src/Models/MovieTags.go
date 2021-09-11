package Models

import "time"

type MovieTags struct {
	ID         int64
	Name       string
	CreateadAt time.Time
	UpdatedAt  time.Time
}
