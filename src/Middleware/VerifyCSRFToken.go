package Middleware

import(
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var padding = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
	'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
	'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9',
}


var (
	// TokenLength is the length of csrf token
	TokenLength = 16

	// TokenKey is the key of csrf token
	// it could be in get-query, post-form or header
	TokenKey = "X-Csrf-Token"

	// TokenCookie is the name of token cookie
	TokenCookie = "X-Csrf-Token"

	// DefaultExpire is the default expire time of cookie
	DefaultExpire = 3600 * 6

	// RandomSec is the flag which represents the random-source
	// will be changed after each period of time
	RandomSec = false

	// randSource will be changed every DefaultExpire time
	randSource = rand.New(rand.NewSource(time.Now().UnixNano()))

	// GenerateToken returns random CSRF token
	GenerateToken = func() string {
		result := make([]byte, TokenLength)
		for i := 0; i < TokenLength; i++ {
			result[i] = padding[randSource.Intn(62)]
		}
		return string(result)
	}
)

func init() {
	if RandomSec {
		go func() {
			for {
				time.Sleep(time.Duration(DefaultExpire) * time.Second)
				randSource = rand.New(rand.NewSource(time.Now().UnixNano()))
			}
		}()
	}
}

// SetCSRFToken set CSRF token in cookie while no token in cookie now
func SetCSRFToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie(TokenCookie)
		if err != nil {
			c.SetCookie(TokenCookie, GenerateToken(), DefaultExpire, "/", "", false, false)
		}
		c.Next()
	}
}

// XCSRF verify the token
// if not match, returns 403
func XCSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		token, ok := c.GetQuery(TokenKey)
		if !ok {
			token = c.GetHeader(TokenKey)
		}

		cookie, err := c.Cookie(TokenCookie)
		if token == "" || err != nil || cookie != token {
			c.AbortWithStatus(403)
		}
		c.Next()
	}
}