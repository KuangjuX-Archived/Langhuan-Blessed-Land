package Models

import(
	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/Databases"
)

type User struct{
	id       int64  `json:"id"`       
    username string `json:"username"` 
    password string `json:"password"` 
}

func (user *User) Users() (users []User, err error) {
    if err = orm.Db.Find(&users).Error; err != nil {
        return
    }
    return
}