package Models

import(
	"time"
)

type Tags struct{
	ID			int
	Name		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}