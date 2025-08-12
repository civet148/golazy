package middleware

import (
	"github.com/gin-gonic/gin"

	"encoding/json"
	"github.com/civet148/log"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	CLAIM_EXPIRE       = "claim_expire"
	CLAIM_ISSUE_AT     = "claim_iat"
	CLAIM_USER_SESSION = "user_session"
)

const (
	HEADER_AUTH_TOKEN      = "Authorization"
	DEFAULT_TOKEN_DURATION = 24 * time.Hour // default one year  = 8760 hour
)

const (
	jwtTokenSecret = "7bdf27cffd5fd105af4efb20b1090bbe"
)

type JwtCode int

const (
	JWT_CODE_SUCCESS             JwtCode = 0
	JWT_CODE_ERROR_CHECK_TOKEN   JwtCode = -1
	JWT_CODE_ERROR_PARSE_TOKEN   JwtCode = -2
	JWT_CODE_ERROR_INVALID_TOKEN JwtCode = -3
	JWT_CODE_ERROR_TOKEN_EXPIRED JwtCode = -4
)

var codeMessages = map[JwtCode]string{
	JWT_CODE_SUCCESS:             "JWT_CODE_SUCCESS",
	JWT_CODE_ERROR_CHECK_TOKEN:   "JWT_CODE_ERROR_CHECK_TOKEN",
	JWT_CODE_ERROR_PARSE_TOKEN:   "JWT_CODE_ERROR_PARSE_TOKEN",
	JWT_CODE_ERROR_INVALID_TOKEN: "JWT_CODE_ERROR_INVALID_TOKEN",
	JWT_CODE_ERROR_TOKEN_EXPIRED: "JWT_CODE_ERROR_TOKEN_EXPIRED",
}

type JwtAuthMiddleware struct {
	WhiteList map[string]bool
}

func NewJwtAuthMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		WhiteList: map[string]bool{}, //添加路由白名单
	}
}

func (m *JwtAuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// White list check for request path
		if ok := m.WhiteList[c.Request.RequestURI]; ok {
			c.Next()
			return
		}

		//TODO: add your middleware logic here

		// Pass through to next handler
		c.Next()
	}
}

// generate JWT token
func GenerateToken(session interface{}, duration ...interface{}) (token string, err error) {

	var d time.Duration
	var claims = make(jwt.MapClaims)

	if len(duration) == 0 {
		d = DEFAULT_TOKEN_DURATION
	} else {
		var ok bool
		if d, ok = duration[0].(time.Duration); !ok {
			d = DEFAULT_TOKEN_DURATION
		}
	}
	var data []byte
	data, err = json.Marshal(session)
	if err != nil {
		return token, log.Errorf(err.Error())
	}
	sign := jwt.New(jwt.SigningMethodHS256)
	claims[CLAIM_EXPIRE] = time.Now().Add(d).Unix()
	claims[CLAIM_ISSUE_AT] = time.Now().Unix()
	claims[CLAIM_USER_SESSION] = string(data)
	sign.Claims = claims

	token, err = sign.SignedString([]byte(jwtTokenSecret))
	return token, err
}

// parse JWT token claims
func ParseToken(c *gin.Context) error {
	strAuthToken := GetAuthToken(c)
	if strAuthToken == "" {
		return log.Errorf("[JWT] request header have no any key '%s'", HEADER_AUTH_TOKEN)
	}
	claims, err := ParseTokenClaims(strAuthToken)
	if err != nil {
		return log.Errorf(err.Error())
	}
	c.Keys = make(map[string]interface{})
	c.Keys[CLAIM_EXPIRE] = int64(claims[CLAIM_EXPIRE].(float64))
	if c.Keys[CLAIM_EXPIRE].(int64) < time.Now().Unix() {
		return log.Errorf("[JWT] token [%s] expired at %v\n", strAuthToken, c.Keys[CLAIM_EXPIRE])
	}

	c.Keys[CLAIM_EXPIRE] = int64(claims[CLAIM_EXPIRE].(float64))
	c.Keys[CLAIM_ISSUE_AT] = int64(claims[CLAIM_ISSUE_AT].(float64))
	c.Keys[CLAIM_USER_SESSION] = claims[CLAIM_USER_SESSION].(string)
	return nil
}

func ParseTokenClaims(strAuthToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(strAuthToken, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtTokenSecret), nil
	})
	if err != nil {
		return jwt.MapClaims{}, log.Errorf("[JWT] parse token error [%s]", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, log.Errorf("[JWT] parse token error: no claims found")
	}
	return claims, nil
}

func GetAuthToken(c *gin.Context) string {
	strToken := c.Request.Header.Get(HEADER_AUTH_TOKEN)
	log.Debugf("AuthToken [%s]", strToken)
	return strToken
}

func GetAuthSessionFromToken(strAuthToken string, session interface{}) error {
	claims, err := ParseTokenClaims(strAuthToken)
	if err != nil {
		return log.Errorf(err.Error())
	}
	strSessionJson := claims[CLAIM_USER_SESSION].(string)
	err = json.Unmarshal([]byte(strSessionJson), session)
	if err != nil {
		return log.Errorf(err.Error())
	}
	return nil
}

func GetAuthSessionFromContext(c *gin.Context, session interface{}) error {
	strAuthToken := GetAuthToken(c)
	return GetAuthSessionFromToken(strAuthToken, session)
}
