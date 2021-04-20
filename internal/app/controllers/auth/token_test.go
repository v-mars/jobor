package auth

import (
	"fmt"
	sys "ops-go/internal/app/controllers/sys/user"
	"ops-go/internal/config"
	"ops-go/pkg/convert"
	goJWT "ops-go/pkg/jwt"
	"testing"
)

func TestGerToken(t *testing.T) {
	var cla map[string]interface{}
	var u = sys.InfoUser{ID:1,Name:"ocean.zhang",Nickname:"ocean.zhang",Username:"ocean.zhang",Roles:[]string{"admin"}}
	if err := convert.StructToMapOut(u, &cla); err!=nil{
		return
	}
	tokenApi := goJWT.New()
	tokenApi.SetClaims(cla)
	tokenApi.Options.TokenType = config.Configs.JWT.TokenType //"Bearer"
	tokenApi.Options.SigningKey = []byte(config.Configs.JWT.TokenKey)
	tokenApi.Options.Expired = 36000
	tokenApi.Options.Claims = cla
	tokenString, _ := tokenApi.GenerateToken()
	fmt.Println("data token:", tokenString.GetAccessToken())


}