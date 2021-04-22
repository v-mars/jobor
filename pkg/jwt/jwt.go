package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const defaultKey = "jS2SnJdxxmTKRNQYh"

// tokenInfo 令牌信息
type tokenInfo struct {
	AccessToken string `json:"access_token"` // 访问令牌
	TokenType   string `json:"token_type"`   // 令牌类型
	ExpiresAt   int64  `json:"expires_at"`   // 令牌到期时间
}

func (t tokenInfo) GetAccessToken() string {
	return t.AccessToken
}

func (t tokenInfo) GetTokenType() string {
	return t.TokenType
}

func (t tokenInfo) GetExpiresAt() int64 {
	return t.ExpiresAt
}

func (t tokenInfo) EncodeToJSON() ([]byte, error) {
	return t.EncodeToJSON()
}

type options struct {
	TokenType     string
	Expired       int
	SigningMethod jwt.SigningMethod
	SigningKey    interface{}
	KeyFunc       jwt.Keyfunc
	Claims        jwt.MapClaims
}

// Option 定义参数项
type Option func(*options)

var defaultOptions = options{
	TokenType:     "Bearer",
	Expired:       3600,
	SigningMethod: jwt.SigningMethodHS256, 	//signingMethod: jwt.SigningMethodHS512,
	SigningKey:    []byte(defaultKey),
	KeyFunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(defaultKey), nil
	},
}

// SetSigningMethod 设定签名方式
func (a *JWTAuth)SetSigningMethod(method jwt.SigningMethod) {
	a.Options.SigningMethod = method
}

// SetSigningKey 设定签名key
func (a *JWTAuth)SetSigningKey(key []byte) {
	a.Options.SigningKey = key
}

// SetExpired 设定令牌过期时长(unit, default3600)
func (a *JWTAuth) SetExpired(expired int) {
	a.Options.Expired = expired
}

// SetKeyfunc 设定验证key的回调函数
func (a *JWTAuth)SetKeyfunc(keyFunc jwt.Keyfunc){
	a.Options.KeyFunc = keyFunc
}

// SetClaims 设定SetClaims
func (a *JWTAuth) SetClaims(claims jwt.MapClaims) {
	a.Options.Claims = claims
}

// New 创建认证实例
func New(opts ...Option) JWTAuth {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	return JWTAuth{
		Options:  &o,
	}
}

type NewFunc func(*JWTAuth)

// JWTAuth jwt认证
type JWTAuth struct {
	Options  *options
}

// GenerateToken 生成令牌
func (a *JWTAuth) GenerateToken() (TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.Options.Expired) * time.Second).Unix()
	token := jwt.New(a.Options.SigningMethod)
	//claims := make(jwt.MapClaims)
	//claims["name"] = "xxx"
	a.Options.Claims["exp"] = expiresAt
	a.Options.Claims["iat"] = now.Unix()  // 签发人
	a.Options.Claims["iss"] = "ops tools go jwt."
	//claims["jti"] = "ops tools go jwt."  // id
	//fmt.Println("a.Options.claims:", a.Options.Claims)
	token.Claims = a.Options.Claims

	tokenString, err := token.SignedString(a.Options.SigningKey)
	if err != nil {
		return nil, err
	}

	info := tokenInfo{
		ExpiresAt:   expiresAt,
		TokenType:   a.Options.TokenType,
		AccessToken: tokenString,
	}
	return info, nil
}

// 解析令牌
func (a *JWTAuth) ParseToken(tokenString string, keys string) (map[string]interface{}, error) {
	key := defaultKey
	if len(keys) > 0 {
		key = keys
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(key), nil
	})
	//fmt.Println("token:", token, len(keys), keys)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, err
	} else if !token.Valid {
		return nil, ErrInvalidToken
	}
	return token.Claims.(jwt.MapClaims), nil
}


// 定义错误
var (
	TokenExpired     error  = errors.New("token is expired")
	TokenNotValidYet error  = errors.New("token not active yet")
	TokenMalformed   error  = errors.New("that's not even a token")
	TokenInvalid     error  = errors.New("invalid token")
	ErrInvalidToken  error  = errors.New("invalid token")
)

// TokenInfo 令牌信息
type TokenInfo interface {
	// 获取访问令牌
	GetAccessToken() string
	// 获取令牌类型
	GetTokenType() string
	// 获取令牌到期时间戳
	GetExpiresAt() int64
	// JSON编码
	EncodeToJSON() ([]byte, error)
}

// Auther 认证接口
type Auther interface {
	// 生成令牌
	GenerateToken() (TokenInfo, error)

	// 销毁令牌
	//DestroyToken(ctx context.Context, accessToken string) error

	// 解析用户ID
	//ParseUserID(ctx context.Context, accessToken string) (string, error)

	// 释放资源
	//Release() error
}