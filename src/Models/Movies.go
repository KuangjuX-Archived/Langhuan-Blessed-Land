package Models

import (
	"time"
)

// 电影元信息
type Movies struct {
	ID           int64
	Name         string
	Directors    string
	Actors       string
	ScreenWriter string
	Language     string
	Length       int16
	Tags         string
	Description  string
	Score        int32
	ReleaseData  time.Time
	CreateadAt   time.Time
	UpdatedAt    time.Time
}
