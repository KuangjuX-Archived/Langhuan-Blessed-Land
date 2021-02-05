package Models

import(
    "errors"
    "golang.org/x/crypto/bcrypt"
    "time"
	"fmt"
	"reflect"

    "github.com/dgrijalva/jwt-go"
	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/DataBases/mysql"
)

type User struct{
	ID       	int64  		`json:"id"`       
	Username 	string 		`json:"username"` 
	Nickname 	string		`json:"nickname"`
    Password 	string 		`json:"password"` 
	Email    	string 		`json:"email"`
	Avatar	 	string 		`json:"avatar"`
	CreatedAt 	time.Time	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"autoUpdateTime"`
}

var AppSecret = ""
var AppIss = "github.com/KuangjuX/Lang-Huan-Blessed-Land"
var ExpireTime = time.Hour * 24

type userStdClaims struct {
	jwt.StandardClaims
	*User
}



func hash(password string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Hash(password string) ([]byte, error) {
	return hash(password)
}


/*jwt*/
func jwtGenerateToken(m *User, d time.Duration) (string, error) {
	m.Password = ""
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", m.ID),
		Issuer:    AppIss,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		User:           m,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	
	return tokenString, err
}

// parser token to get user message
func JwtParserUser(tokenString string) (*User, error){
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}

	claims := userStdClaims{}
	_ , err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return claims.User, nil
}


//User is exist?
func IsExistUser(identify string)(bool, error) {
	res := orm.Db.Where("username = ? or email = ?", identify, identify).First(&User{})
	err := res.Error
	if err != nil && err != orm.ErrorRecordNotFound{
		return false, err
	}
	if err != orm.ErrorRecordNotFound{
		// fmt.Printf("User has exist.\n")
		return true, nil
	}
	return false, nil
}




func CreatUser(username, password, email string) (string, error) {
    hash_password, err := hash(password)
    if err != nil{
        return "Fail to create", err
    }
    var user = User{
        Username: username,
        Password: string(hash_password),
        Email: email,
    }
    
    res := orm.Db.Create(&user)
    if res.Error == nil {
        return "Success to create", nil
    }else {
        return "Fail to create", res.Error
    }

}

func (user *User) CreateUser() (error){
	hash_password, err := hash(user.Password)
	if err != nil{
		return err
	}

	user.Password = string(hash_password)

	res := orm.Db.Create(user)

	if res.Error == nil {
        return nil
    }else {
        return res.Error
    }
}

func (user *User) Login() (string, error) {
    user.ID = 0
    if user.Password == "" {
        return "", errors.New("password is required")
    }
    input_password := user.Password

    err := orm.Db.Where("username = ? or email = ?",user.Username, user.Email).First(user).Error
    if err != nil{
        return "", err
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input_password)); err != nil {
		return "", err
	}
	user.Password = ""
	token, err := jwtGenerateToken(user, ExpireTime)
	return token, err

}

func (user *User)OAuthLogin() (string, error) {
	if err := orm.Db.Where("username = ? or email = ?", user.Username, user.Email).First(user).Error; err != nil {
		return "", err
	} 

	user.Password = ""
	token, err := jwtGenerateToken(user, ExpireTime)
	return token, err
}


func ModifyNickname(user_id int64, new_nickname string) (error){
	DB := orm.Db
	result := DB.Model(User{}).Where("id = ?", user_id).Updates(User{Nickname : new_nickname})

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func ModifyEmail(user_id int64, new_email string) (error){
	DB := orm.Db
	result := DB.Model(User{}).Where("id = ?", user_id).Updates(User{Email: new_email})
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func ModifyAvatar(user_id int64, new_avatar string) (error){
	DB := orm.Db
	result := DB.Model(User{}).Where("id = ?", user_id).Updates(User{Avatar: new_avatar})
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// Get Last User by key
func GetLastUserInfoByKey(key string)(string, error){
	// get the last user in database
	var user User
	if res := orm.Db.Last(&user); res.Error != nil{
		return "", res.Error
	}

	tmp := reflect.ValueOf(user)

	ret := tmp.FieldByName(key).String()
	return ret, nil
}