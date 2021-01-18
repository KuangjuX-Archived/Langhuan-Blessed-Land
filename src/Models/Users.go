package Models

import(
    "errors"
    "golang.org/x/crypto/bcrypt"
    "time"
    "fmt"

    "github.com/dgrijalva/jwt-go"
    orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/Databases"
)

type User struct{
	ID       int64  `json:"id"`       
	Username string `json:"username"` 
	Nickname string	`json:"nickname"`
    Password string `json:"password"` 
	Email    string `json:"email"`
	Avatar	 string `json:"avatar"`
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


func CreatUser(username, password, email string) (string, error) {
    hash_password, err := hash(password)
    if err != nil{
        return "创建失败", err
    }
    var user = User{
        Username: username,
        Password: string(hash_password),
        Email: email,
    }
    
    result := orm.Db.Create(&user)
    if result.Error == nil {
        return "创建成功", nil
    }else {
        return "创建失败", result.Error
    }

}

func (user *User) Login() (string, error) {
    user.ID = 0
    if user.Password == "" {
        return "", errors.New("password is required")
    }
    input_password := user.Password

    err := orm.Db.Where("username = ? or email = ?",user.Username, user.Email).First(&user).Error
    if err != nil{
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input_password)); err != nil {
		return "", err
	}
	user.Password = ""
	data, err := jwtGenerateToken(user, time.Hour*24*365)
	return data, err

}

func ModifyNickname(user_id int64, new_nickname string) (error){
	DB := orm.Db
	DB = DB.Model(User{}).Where("user_id = ?", user_id).Updates(User{Nickname : new_nickname})
	if DB.Error != nil {
		return DB.Error
	}
	return nil
}

func ModifyEmail(user_id int64, new_email string) (error){
	DB := orm.Db
	DB = DB.Model(User{}).Where("user_id = ?", user_id).Updates(User{Email: new_email})
	if DB.Error != nil {
		return DB.Error
	}
	return nil
}

func ModifyAvatar(user_id int64, new_avatar string) (error){
	DB := orm.Db
	DB = DB.Model(User{}).Where("user_id = ?", user_id).Updates(User{Avatar: new_avatar})
	if DB.Error != nil {
		return DB.Error
	}
	return nil
}