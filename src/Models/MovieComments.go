package Models

import "time"

type MovieComments struct {
	ID          int64
	UserID      int64
	MovieID     int64
	ToCommentID int64
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
