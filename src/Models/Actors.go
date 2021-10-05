package Models

import (
	"time"

	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/DataBases/mysql"
)

type Actor struct {
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

func CreateActors(
	id int64, name string, gender bool,
	borth_date time.Time, borth_area string,
	avatar, description string,
) (string, error) {
	actor := Actor{
		ID:          id,
		Name:        name,
		Gender:      gender,
		BorthDate:   borth_date,
		BorthArea:   borth_area,
		Avatar:      avatar,
		Description: description,
	}
	res := orm.Db.Create(&actor)
	if res.Error != nil {
		return "Fail to create actor", res.Error
	} else {
		return "Success to create actor", nil
	}
}
