package Models

import(
	"time"
)

type LikesMap struct{
	ID			int64			`json:"id"`
	UserID		int64			`json:"user_id"`
	ArticleID	int64			`json:"article_id"`
	CreatedAt 	time.Time		`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt 	time.Time		`json:"updated_at" gorm:"autoUpdateTime"`
}