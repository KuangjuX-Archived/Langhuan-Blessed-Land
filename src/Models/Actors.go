package Models

import "time"

type Actors struct {
	ID          int64
	Name        string
	Gender      bool
	BorthData   time.Time
	BorthArea   string
	Avatar      string
	Description string
	CreateadAt  time.Time
	UpdatedAt   time.Time
}
