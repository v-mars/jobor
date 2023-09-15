package ldap

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-ldap/ldap/v3"
)

func TestLDAP_Authentication(t *testing.T) {
	// cn=ab,cn=g1,ou=coo,dc=example,dc=org
	var ldapApi = NewLDAP()
	ldapApi.Option.Addr = fmt.Sprintf("%s:%s", "127.0.0.1", "389")
	//ldapApi.Option.Addr = fmt.Sprintf("ldap://%s:%s", "127.0.0.1", "389")
	//ldapApi.Option.BaseDN = "dc=example,dc=org"
	ldapApi.Option.BaseDN = "dc=example,dc=org"
	ldapApi.Option.AuthFilter = "(cn=%s)"
	//ldapApi.Option.Attributes = []string{"uid","cn"}
	ldapApi.Option.Attributes = []string{"cn", "mail", "displayName", "dn"}
	ldapApi.Option.Username = "cn=admin,dc=example,dc=org"
	ldapApi.Option.Password = "admin"
	user, ldapAuthErr := ldapApi.Authentication("ab", "123")
	fmt.Println("ldapAuthErr:", ldapAuthErr)
	fmt.Println("user:", user)
}

func TestLDAP_Authentication1(t *testing.T) {
	aa()
}

func aa() {
	ldap.DefaultTimeout = time.Second * 3
	//conn, err := ldap.DialTLS("tcp",fmt.Sprintf("%s:%s", "127.0.0.1", "636"), &tls.Config{InsecureSkipVerify: true})
	//conn, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%s", "127.0.0.1", "389"))
	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", "389"))
	if err != nil {
		panic(err)
	}
	//err = conn.Bind("ab", "123")
	err = conn.Bind("cn=admin,dc=example,dc=org", "admin")
	if err != nil {
		fmt.Print("bind err:", err)
		return
	}

	//Filter:=fmt.Sprintf("(cn=*)")
	Filter := fmt.Sprintf("(cn=%s)", "ab")
	Attributes := []string{"cn"}
	searchRequest := ldap.NewSearchRequest(
		"ou=coo,dc=example,dc=org",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, // ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, ldap.DerefAlways
		0,                      // SizeLimit: 大小设置,一般设置为0
		0,                      // TimeLimit: 时间设置,一般设置为0
		false,                  //TypesOnly:  设置false（返回的值要多一点）
		Filter,                 // "(objectClass=*)",
		Attributes,             //Attributes 需要返回的属性值
		nil)
	cur, err := conn.Search(searchRequest)
	fmt.Println("cur.searchRequest:", searchRequest)
	if err != nil {
		fmt.Println("cur.Search err:", err)
		return
	}
	fmt.Println("entry:", cur.Entries[0].DN)
}
