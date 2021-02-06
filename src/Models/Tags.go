package Models

import(
	"time"

	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/DataBases/mysql"
)

type Tag struct{
	ID			int
	Name		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

func GetTagByID(id int)(interface{}, error){
	var tag Tag
	if err := orm.Db.Where("id = ?", id).First(&tag).Error; err != nil{
		return nil, err
	}
	return tag, nil
}