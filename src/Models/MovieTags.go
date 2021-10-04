package Models

import "time"

type MovieTag struct {
	ID         int64
	Name       string
	CreateadAt time.Time
	UpdatedAt  time.Time
}
