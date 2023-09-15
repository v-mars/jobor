package oidc_callback

import (
	"context"
	"fmt"
	"testing"
)

func TestCallbackPasswd(t *testing.T) {
	//o := NewOIDC(context.Background(), "", "", "", "", "")
	//info, err := o.OidcCallbackWithPassword(context.Background(), "", "xxxx", "adfsasdf")
	//if err != nil {
	//	fmt.Println("err1", err)
	//	return
	//}
	//var claims json.RawMessage
	//_ = info.Claims(&claims)
	////marshal, _ := json.Marshal(info)
	//fmt.Println("userInfo:", string(claims))

	_, err := InitProvider(context.Background(), "https://sso.xxx.local")
	if err != nil {
		fmt.Println("err2", err)
		return
	}
}
