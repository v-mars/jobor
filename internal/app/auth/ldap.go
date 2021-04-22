package auth

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"golang.org/x/text/encoding/unicode"
	"jobor/internal/errors"
	"time"
)

type option struct {
	Addr       string     // addr := fmt.Sprintf("%s:%s", "10.30.108.52", "389") //addr:127.0.0.1:389
	Config     tls.Config
	BaseDN     string
	AuthFilter string
	Domain     string
	Username   string
	Password   string
}

var defaultOptions = option{
	Addr:    fmt.Sprintf("%s:%s", "10.30.108.52", "389"),
	Config:  tls.Config{InsecureSkipVerify: true},
	BaseDN:  "dc=ocean,dc=cn",
	AuthFilter:  "(&(sAMAccountName=%s))",		// (&(objectclass=group)(cn=%s*))
	Domain:  "ocean.cn",
}

func NewLDAP() *LDAP {
	l := LDAP{Option: &defaultOptions}
	return &l
}

type Result struct {
	DisplayName string `json:"displayname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	EmployeeId  string `json:"employeeId"`
}

type LDAP struct {
	Option	*option
}

func (a *LDAP) Bind(username, password string) error {
	conn, err:= a.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.Bind(username, password)
	if err != nil {
		//fmt.Printf("Ldap server disconnect. ")
		return err
	}
	return nil
}

func (a *LDAP) Conn() (*ldap.Conn, error) {
	ldap.DefaultTimeout = time.Second*3
	conn, err := ldap.DialTLS("tcp", a.Option.Addr, &a.Option.Config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (a *LDAP) User(username string, conn *ldap.Conn) (Result, error) {
	searchRequest := a.GetSearchRequest(username)
	//var cur *ldap.SearchResult
	cur, err := conn.Search(searchRequest)
	if err != nil {
		return Result{}, err
	}
	if len(cur.Entries) == 0 {
		return Result{}, errors.New("Not found user.")
	}
	var r = Result{
		DisplayName:        cur.Entries[0].GetAttributeValue("displayName"),
		Username:           cur.Entries[0].GetAttributeValue("sAMAccountName"),
		Email:              cur.Entries[0].GetAttributeValue("mail"),
		Phone:              cur.Entries[0].GetAttributeValue("mobile"),
	}
	return r,nil
}

func (a *LDAP) Group(name string, conn *ldap.Conn) (string, error) {
	Filter := fmt.Sprintf("(&(objectclass=group)(cn=%s*))", name)
	sql := ldap.NewSearchRequest(a.Option.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		Filter,
		[]string{},
		nil)
	cur, err := conn.Search(sql)
	if err != nil {
		return "", err
	}
	if len(cur.Entries) == 0 {
		return "", errors.New("Not found group.")
	}

	return "",nil
}

func (a *LDAP) GetGroupUsers(GroupName string, conn *ldap.Conn) {
	_, _ = a.Group(GroupName, conn)
}

func (a LDAP) Authentication(username, password string) (Result, error) {
	conn, err:= a.Conn()
	if err != nil {
		return Result{},err
	}
	defer conn.Close()

	err = conn.Bind(username, password)
	if err != nil {
		return Result{},err
	}
	result, err := a.User(username, conn)
	return result,err
}

func (a *LDAP) RestPassword(username, OldPassword, NewPassword string, change bool) error {
	conn, err:= a.Conn()
	if err != nil {
		//panic("Ldap server disconnect.")
		return err
	}
	defer conn.Close()

	if change{
		err = conn.Bind(username, OldPassword)
		if err != nil {
			//panic("Password error.")
			return err
		}
	}

	searchRequest := a.GetSearchRequest(username)

	var cur *ldap.SearchResult

	if cur, err = conn.Search(searchRequest); err != nil {
		//panic("Ldap server search failed.")
		return err
	}

	if len(cur.Entries) == 0 {
		//panic("Not found user.")
		return errors.New("Not found user.")
	}

	err = conn.Bind("", "")
	if err != nil {
		//panic("Ldap server unbind.")
		return err
	}

	var userDn = fmt.Sprintf("CN=%s,OU=group,%s", username, "base_dn")

	sql2 := ldap.NewModifyRequest(userDn, nil)

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, _ := utf16.NewEncoder().String("\"" + NewPassword + "\"")

	sql2.Replace("unicodePwd", []string{pwdEncoded})
	sql2.Replace("userAccountControl", []string{"512"})

	if err := conn.Modify(sql2); err != nil {
		//panic("Ldap password modify failed.")
		return err
	}
	return nil
}

func (a *LDAP) GetSearchRequest(username string) *ldap.SearchRequest {
	Filter:=fmt.Sprintf(a.Option.AuthFilter, username)
	//Filter:=fmt.Sprintf("(sAMAccountName=%s)", username)
	Attributes := []string{"sAMAccountName", "displayName", "mail", "mobile", "employeeID","givenName"}
	searchRequest := ldap.NewSearchRequest(a.Option.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		Filter,
		Attributes,
		nil)
	return searchRequest
}


func ActionLdapLogin() {
	params := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	addr := fmt.Sprintf("%s:%s", "10.30.108.52", "389") //addr:127.0.0.1:389

	conn, err := ldap.DialTLS("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		panic("Ldap server disconnect.")
		return
	}
	defer conn.Close()

	err = conn.Bind(params.Username, params.Password)
	if err != nil {
		panic("Password error.")
		return
	}

	sql := ldap.NewSearchRequest("cn=xx,cn=com",
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		fmt.Sprintf("(sAMAccountName=%s)", params.Username),
		[]string{"sAMAccountName", "displayName", "mail", "mobile", "employeeID", "givenName"},
		nil)

	var cur *ldap.SearchResult

	if cur, err = conn.Search(sql); err != nil {
		panic("Ldap server search failed.")
		return
	}

	if len(cur.Entries) == 0 {
		panic( "Not found user.")
		return
	}
	/*
	var result = struct {
		Name               string `json:"displayname"`
		Username           string `json:"username"`
		Email              string `json:"email"`
		Phone              string `json:"phone"`
		EmployeeId         string `json:"employeeId"`
	}{
		DisplayName:        cur.Entries[0].GetAttributeValue("displayName"),
		Username:           cur.Entries[0].GetAttributeValue("sAMAccountName"),
		Email:              cur.Entries[0].GetAttributeValue("mail"),
		Phone:              cur.Entries[0].GetAttributeValue("mobile"),
	}

	 */

	//return result
}

