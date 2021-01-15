package Models

import(
	"time"

	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/Databases"
)

type Articles struct{
	ID 		int64		`json:"id"`
	UserID  int64		`json:"user_id"`
	TagID	int64		`json:"tag_id"`
	Title	string		`json:"title"`
	Content	string		`json:"content"`
	Created time.Time	`json:"created_at"`
	Update	time.Time	`json:"update_at"`
}

func GetAllArticles() ([]Articles, error){
	var articles []Articles
	result := orm.Db.Find(&articles)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}