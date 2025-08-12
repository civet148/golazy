package gen

import (
	_ "embed"
	"github.com/civet148/golazy/parser"
	"github.com/civet148/golazy/utils"
	"strings"
)

const (
	defaultMiddlewareCors = "Cors"    //default cors domain middleware
	defaultMiddlewareJwt  = "JwtAuth" //default jwt middleware
)

const (
	corsMiddlewareImpl = `method := c.Request.Method

		//set header for cross-domain
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Auth-Token, *")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, content-Disposition")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Set("content-type", "application/json, text/plain, multipart/form-data, */*")

		// abort with options method (code=204)
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}`

	jwtAuthMiddlewareImports = `"encoding/json"
"time"
"github.com/civet148/log"
"github.com/dgrijalva/jwt-go"
`
	jwtAuthMiddlewareConstants = `const (
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
`

	jwtAuthMiddlewareImpl = `// White list check for request path
		if ok := m.WhiteList[c.Request.RequestURI]; ok {
			c.Next()
			return
		}`
	jwtAuthMiddlewareExtraFunctions = `
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
`
)

//go:embed tpls/middleware.tpl
var middlewareImplementCode string

func genMiddleware(cfg *Config, api *parser.ApiService) error {

	middlewares := utils.GetMiddleware(api.Server)
	for _, item := range middlewares {
		middlewareFilename := strings.TrimSuffix(strings.ToLower(item), "middleware") + "_middleware"
		filename, err := utils.FileNamingFormat(cfg.Style, middlewareFilename)
		if err != nil {
			return err
		}

		impl := getDefaultMiddlewareImpl(item)

		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		err = genFile(fileGenConfig{
			dir:             cfg.OutDir,
			subdir:          middlewareDir,
			filename:        filename + ".go",
			templateName:    "contextTemplate",
			category:        category,
			builtinTemplate: middlewareImplementCode,
			data: map[string]any{
				"name":                  strings.Title(name),
				"hasNetHttp":            impl.HasNetHttp,
				"hasWhiteList":          impl.HasWhiteList,
				"constants":             impl.Constants,
				"defaultMiddlewareImpl": impl.Implement,
				"importPkg":             impl.ImportPkg,
				"extraFunctions":        impl.ExtraFunctions,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

type middlewareImpl struct {
	HasNetHttp     bool
	HasWhiteList   bool
	ImportPkg      string
	Constants      string
	Implement      string
	ExtraFunctions string
}

func getDefaultMiddlewareImpl(name string) *middlewareImpl {
	switch name {
	case defaultMiddlewareCors:
		return &middlewareImpl{
			HasNetHttp:     true,
			HasWhiteList:   false,
			ImportPkg:      "",
			Implement:      corsMiddlewareImpl,
			ExtraFunctions: "",
		}
	case defaultMiddlewareJwt:
		return &middlewareImpl{
			HasNetHttp:     false,
			HasWhiteList:   true,
			ImportPkg:      jwtAuthMiddlewareImports,
			Implement:      jwtAuthMiddlewareImpl,
			Constants:      jwtAuthMiddlewareConstants,
			ExtraFunctions: jwtAuthMiddlewareExtraFunctions,
		}
	}
	return &middlewareImpl{
		HasNetHttp:     false,
		HasWhiteList:   false,
		ImportPkg:      "",
		Constants:      "",
		Implement:      "",
		ExtraFunctions: "",
	}
}
