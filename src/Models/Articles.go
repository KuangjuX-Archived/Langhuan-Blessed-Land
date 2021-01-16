package Models

import(
	"time"

	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/Databases"
)

type Article struct{
	ID 		int64		`json:"id" gorm:"primaryKey" form:"id"`
	UserID  int64		`json:"user_id" form:"user_id"`
	TagID	int64		`json:"tag_id" form:"tag_id"`
	Title	string		`json:"title"`
	Content	string		`json:"content"`
	Created time.Time	`json:"created_at"`
	Update	time.Time	`json:"update_at"`
}

type Pagination struct{
	Ok    bool        `json:"ok"`
	Size  uint        `form:"size" json:"size"`
	Page  uint        `form:"page" json:"page"`
	Data  interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total uint        `json:"total"`
}

type ArticlePages struct{
	Article
	Pagination
	FromTime string `form:"from_time"` 
	ToTime   string `form:"to_time"`  
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

func GetArticlesByTag(tag_id string) ([]Article, error){
	var articles []Article
	result := orm.Db.Where("tag_id = ?", tag_id).Find(&articles)
	err := result.Error

	if err != nil{
		return articles, err
	}

	return articles, err
}

func (o ArticlePages) Search() (list *[]Article, total uint, err error){
	list = &[]ArticlePages{}

	tx := orm.Db(o.Article)


}