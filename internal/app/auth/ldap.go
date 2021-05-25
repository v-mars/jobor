package auth

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"golang.org/x/text/encoding/unicode"
	"jobor/internal/errors"
	"jobor/pkg/logger"
	"time"
)

const (
	defaultFilter = "(&(sAMAccountName=%s))"
	DefaultTimeout = time.Second*3
)

var (
	AdAttributes = []string{"sAMAccountName", "displayName", "mail", "mobile", "employeeID","givenName"}
	OpenldapAttributes = []string{"cn", "mail","displayName", "uid"}
)

type option struct {
	Addr       string     // addr := fmt.Sprintf("%s:%s", "10.30.108.52", "389") //addr:127.0.0.1:389
	Config     tls.Config
	BaseDN     string
	AuthFilter string
	Attributes []string
	Domain     string
	Username   string
	Password   string
	Tls        bool
}

var defaultOptions = option{
	Addr:    fmt.Sprintf("%s:%s", "127.0.0.1", "389"),
	Config:  tls.Config{InsecureSkipVerify: true},
	BaseDN:  "dc=example,dc=org",
	AuthFilter:  defaultFilter,		// (&(sAMAccountName=%s)) (&(objectclass=group)(cn=%s*))  (cn=%s)
	Attributes: AdAttributes,  // []string{"cn", "mail","displayName"}
	Domain:  "example.org",
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
	conn, err:= a.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.Bind(username, password)
	if err != nil {
		return err
	}
	logger.Debugf("ldap user[%s] bind success", username)
	return nil
}

func (a *LDAP) Connect() (*ldap.Conn, error) {
	ldap.DefaultTimeout = DefaultTimeout
	conn, err := ldap.Dial("tcp", a.Option.Addr)
	if a.Option.Tls{
		conn, err = ldap.DialTLS("tcp", a.Option.Addr, &a.Option.Config)
	}
	if err != nil {
		return nil, err
	}
	logger.Debugf("ldap addr[%s] conn success", a.Option.Addr)
	return conn, nil
}

func (a *LDAP) User(username string, conn *ldap.Conn) (*ldap.Entry, error) {
	searchRequest := a.GetSearchRequest(username)
	cur, err := conn.Search(searchRequest)
	if err != nil {
		logger.Errorf("ldap search err: %s", err)
		return nil, err
	}
	if len(cur.Entries) == 0 {
		logger.Errorf("Not found ldap user.")
		return nil, errors.New("Not found ldap user.")
	}
	entry := cur.Entries[0]
	return entry, err
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
	var r = Result{}
	conn, err:= a.Connect()
	if err != nil {
		return r,err
	}
	defer conn.Close()

	err = conn.Bind(a.Option.Username, a.Option.Password)
	if err != nil {
		logger.Errorf("ldap admin bind err: %s", err)
		return r,err
	}
	entry, err := a.User(username, conn)
	if entry==nil {
		logger.Errorf("Not found ldap user info.")
		return r, errors.New("Not found ldap user info.")
	}
	err = conn.Bind(entry.DN, password)
	if err != nil {
		logger.Errorf("ldap user authentication err: %s", err)
		return Result{},err
	}
	if a.Option.AuthFilter == defaultFilter {
		r = Result{
			DisplayName:        entry.GetAttributeValue("displayName"),
			Username:           entry.GetAttributeValue("sAMAccountName"),
			Email:              entry.GetAttributeValue("mail"),
			Phone:              entry.GetAttributeValue("mobile"),
		}
	} else {
		r = Result{
			DisplayName:        entry.GetAttributeValue("displayName"),
			Username:           entry.GetAttributeValue("cn"),
			Email:              entry.GetAttributeValue("mail"),
		}
	}
	return r,err
}

func (a *LDAP) RestPassword(username, OldPassword, NewPassword string, change bool) error {
	conn, err:= a.Connect()
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
	//Filter:=fmt.Sprintf("(sAMAccountName=%s)", username)
	Filter:=fmt.Sprintf(a.Option.AuthFilter, username)
	Attributes := a.Option.Attributes
	if a.Option.AuthFilter!=defaultFilter{Attributes=OpenldapAttributes}
	searchRequest := ldap.NewSearchRequest(a.Option.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,		// ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, ldap.DerefAlways
		0,		// SizeLimit: 大小设置,一般设置为0
		0, 	// TimeLimit: 时间设置,一般设置为0
		false,	//TypesOnly:  设置false（返回的值要多一点）
		Filter,					// "(objectClass=*)",
		Attributes,				//Attributes 需要返回的属性值
		nil)	//Controls:  控制
	return searchRequest
}

