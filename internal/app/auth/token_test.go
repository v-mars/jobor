package auth

import (
	"fmt"
	sys "jobor/internal/app/sys/user"
	"jobor/internal/config"
	"jobor/pkg/convert"
	goJWT "jobor/pkg/jwt"
	"testing"
)

func TestGerToken(t *testing.T) {
	var cla map[string]interface{}
	var u = sys.InfoUser{ID: 1,Name:"admin",Nickname:"admin",Username:"admin",Roles:[]string{"admin"}}
	if err := convert.StructToMapOut(u, &cla); err!=nil{
		return
	}
	tokenApi := goJWT.New()
	tokenApi.SetClaims(cla)
	tokenApi.Options.TokenType = config.Configs.JWT.TokenType //"Bearer"
	tokenApi.Options.SigningKey = []byte("jS2SnJdxxmTKRNQYh")
	tokenApi.Options.Expired = 3600*24*365*10
	tokenApi.Options.Claims = cla
	tokenString, _ := tokenApi.GenerateToken()
	fmt.Println("data token:", tokenString.GetAccessToken())
	parseToken, err := tokenApi.ParseToken(tokenString.GetAccessToken(), "jS2SnJdxxmTKRNQYh")
	if err != nil {
		return
	}
	fmt.Println("parseToken:", parseToken)


}