package Models

import(
	"time"

	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/Databases"
)

type Article struct{
	ID 		int64		`json:"id" gorm:"primaryKey"`
	UserID  int64		`json:"user_id"`
	TagID	int64		`json:"tag_id"`
	Title	string		`json:"title"`
	Content	string		`json:"content"`
	Created time.Time	`json:"created_at"`
	Update	time.Time	`json:"update_at"`
}

func GetAllArticles() ([]Article, error){
	var articles []Article
	result := orm.Db.Find(&articles)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func GetArticleByTag(tag_id string) (Article, error){
	var article Article
	result := orm.Db.Where("tag_id = ?", tag_id).First(&article)
	err := result.Error

	if err != nil{
		return article, err
	}

	return article, err
}