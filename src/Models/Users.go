package Models

import(
    orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/Databases"
    "golang.org/x/crypto/bcrypt"
    "strings"
)

type User struct{
	ID       int64  `json:"id"`       
    Username string `json:"username"` 
    Password string `json:"password"` 
    Email    string `json:"email"`
}

// TrimUsername removes whitespaces in the username
func TrimUsername(username *string) {
	trimmed := strings.Trim(*username, " ")
	*username = trimmed
}

func hash(password string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// EncryptPassword encrypts the password
func EncryptPassword(pass *string) error {
	hash, err := hash(*pass)
	if err != nil {
		return err
	}
	*pass = string(hash)
	return nil
}

// Validate checks if the user has input the required fields
func Validate(username, password string) error {
	TrimUsername(&username)
	if username == "" {
		return ErrUsernameRequired
	}
	if password == "" {
		return ErrPasswordRequired
	}
	return nil
}

// VerifyPassword verifies that the password matches the hash
func VerifyPassword(hash, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
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