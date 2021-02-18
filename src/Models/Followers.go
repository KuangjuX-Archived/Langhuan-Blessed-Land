package Models

import(
	"time"
	"errors"

	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/DataBases/mysql"
)

type Follower struct{
	ID			int
	UserID		int
	FollowerID	int
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

func (user *User)IsExistFollow(follower_id int)(bool, error){
	res := orm.Db.Where("user_id = ? and follower_id = ?", user.ID, follower_id).First(&Follower{})
	err := res.Error
	if err != nil && err == orm.ErrorRecordNotFound {
		return false, nil
	}else if err == nil {
		return true, nil
	}

	return false, err
}

func (user *User)Follow(follower_id int)(error){
	follower := &Follower{
		UserID: int(user.ID),
		FollowerID: follower_id,
	}
	if res := orm.Db.Create(follower); res.Error != nil{
		return res.Error
	}

	return nil
}

func (user *User)UnFollow(follower_id int)(error){
	res := orm.Db.Where("user_id = ? and follower_id = ?", user.ID, follower_id).Delete(Follower{})
	if res.Error != nil{
		return res.Error
	}
	return nil
}


func GetFollowersByPage(page, size, user_id int)([]Follower, error){
	var followers []Follower
	followers = make([]Follower, 0)
	DB := orm.Db
	if page >= 0 && size >= 0{
		res := DB.Where("user_id = ?", user_id).Limit(size).Offset((page - 1)*size).Find(&followers)
		if res.Error != nil{
			return []Follower{}, res.Error
		}

		return followers, nil
	}

	return []Follower{}, errors.New("Invalid Page or Size.")
}