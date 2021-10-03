package Models

import "time"

type Actors struct {
	ID          int64
	Name        string
	Gender      bool
	BorthDate   time.Time
	BorthArea   string
	Avatar      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
