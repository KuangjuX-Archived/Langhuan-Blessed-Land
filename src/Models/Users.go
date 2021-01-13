package Models

import(
    orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/Databases"
)

type User struct{
	ID       int64  `json:"id"`       
    Username string `json:"username"` 
    Password string `json:"password"` 
    Email    string `json:"email"`
}

func (user *User) Users() (users []User, err error) {
    if err = orm.Db.Find(&users).Error; err != nil {
        return
    }
    return
}

func CreatUser(username, password, email string) int {
    var user = User{
        Username: username,
        Password: password,
        Email: email,
    }
    
    result := orm.Db.Create(&user)

    if result.Error != nil {
       return 1
    }else
    {
        return 0
    }
}