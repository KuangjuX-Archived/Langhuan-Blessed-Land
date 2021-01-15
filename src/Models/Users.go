package Models

import(
    "strings"
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
    Password string `json:"password"` 
    Email    string `json:"email"`
}

var AppSecret = ""
var AppIss = "github.com/KuangjuX/Lang-Huan-Blessed-Land"
var ExpireTime = time.Hour * 24

type userStdClaims struct {
	jwt.StandardClaims
	*User
}


// ErrUsernameRequired occurs when the username field is left blank
var ErrUsernameRequired = errors.New("validate: username required")

// ErrPasswordRequired occurs when the password field is left blank
var ErrPasswordRequired = errors.New("validate: password required")

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

// 0 username
// 1 email
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