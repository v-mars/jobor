package ldap

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/unicode"
)

const (
	defaultFilter  = "(&(sAMAccountName=%s))"
	DefaultTimeout = time.Second * 3
)

var (
	AdAttributes       = []string{"sAMAccountName", "displayName", "mail", "mobile", "employeeID", "givenName"}
	OpenldapAttributes = []string{"cn", "mail", "displayName", "uid"}
)

type option struct {
	Enabled    bool
	Addr       string // addr := fmt.Sprintf("%s:%s", "10.30.108.52", "389") //addr:127.0.0.1:389
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
	Addr:       fmt.Sprintf("%s:%s", "127.0.0.1", "389"),
	Config:     tls.Config{InsecureSkipVerify: true},
	BaseDN:     "dc=example,dc=org",
	AuthFilter: defaultFilter, // (&(sAMAccountName=%s)) (&(objectclass=group)(cn=%s*))  (cn=%s)
	Attributes: AdAttributes,  // []string{"cn", "mail","displayName"}
	Domain:     "example.org",
}

func NewLDAP() *LDAP {
	l := LDAP{Option: &defaultOptions, Logger: logrus.New()}
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
	Option *option
	Logger *logrus.Logger
}

func (a *LDAP) Bind(username, password string) error {
	conn, err := a.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.Bind(username, password)
	if err != nil {
		return err
	}
	hlog.Debugf("ldap user[%s] bind success", username)
	return nil
}

func (a *LDAP) Connect() (*ldap.Conn, error) {
	ldap.DefaultTimeout = DefaultTimeout
	conn, err := ldap.Dial("tcp", a.Option.Addr)
	if a.Option.Tls {
		conn, err = ldap.DialTLS("tcp", a.Option.Addr, &a.Option.Config)
	}
	if err != nil {
		return nil, err
	}
	hlog.Debugf("ldap addr[%s] conn success", a.Option.Addr)
	return conn, nil
}

func (a *LDAP) Authentication(username, password string) (Result, error) {
	var r = Result{}
	conn, err := a.Connect()
	if err != nil {
		return r, err
	}
	defer conn.Close()

	err = conn.Bind(a.Option.Username, a.Option.Password)
	if err != nil {
		hlog.Errorf("ldap admin bind err: %s", err)
		return r, err
	}
	hlog.Debugf("ldap admin bind success")
	entry, err := a.User(username, conn)
	if entry == nil {
		hlog.Errorf("Not found ldap user info.")
		return r, fmt.Errorf("Not found ldap user info.")
	}
	hlog.Debugf("ldap search user[%s] success", username)
	err = conn.Bind(entry.DN, password)
	if err != nil {
		hlog.Errorf("ldap user authentication err: %s", err)
		return Result{}, err
	}
	hlog.Debugf("ldap user[%s] authentication success", username)
	if a.Option.AuthFilter == defaultFilter {
		r = Result{
			DisplayName: entry.GetAttributeValue("displayName"),
			Username:    entry.GetAttributeValue("sAMAccountName"),
			Email:       entry.GetAttributeValue("mail"),
			Phone:       entry.GetAttributeValue("mobile"),
		}
	} else {
		r = Result{
			DisplayName: entry.GetAttributeValue("displayName"),
			Username:    entry.GetAttributeValue("cn"),
			Email:       entry.GetAttributeValue("mail"),
		}
	}
	return r, err
}

func (a *LDAP) GetUser(username string) (Result, error) {
	var r = Result{}
	conn, err := a.Connect()
	if err != nil {
		return r, err
	}
	defer conn.Close()

	err = conn.Bind(a.Option.Username, a.Option.Password)
	if err != nil {
		hlog.Errorf("ldap admin bind err: %s", err)
		return r, err
	}
	hlog.Debugf("ldap admin bind success")
	entry, err := a.User(username, conn)
	if entry == nil {
		hlog.Errorf("Not found ldap user info.")
		return r, fmt.Errorf("not found ldap user info")
	}
	hlog.Debugf("ldap search user[%s] success", username)
	if a.Option.AuthFilter == defaultFilter {
		r = Result{
			DisplayName: entry.GetAttributeValue("displayName"),
			Username:    entry.GetAttributeValue("sAMAccountName"),
			Email:       entry.GetAttributeValue("mail"),
			Phone:       entry.GetAttributeValue("mobile"),
		}
	} else {
		r = Result{
			DisplayName: entry.GetAttributeValue("displayName"),
			Username:    entry.GetAttributeValue("cn"),
			Email:       entry.GetAttributeValue("mail"),
		}
	}
	return r, nil
}

func (a *LDAP) User(username string, conn *ldap.Conn) (*ldap.Entry, error) {
	searchRequest := a.GetSearchRequest(username)
	cur, err := conn.Search(searchRequest)
	if err != nil {
		hlog.Errorf("ldap search err: %s", err)
		return nil, err
	}
	if len(cur.Entries) == 0 {
		hlog.Errorf("Not found ldap user.")
		return nil, fmt.Errorf("Not found ldap user.")
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
		return "", fmt.Errorf("Not found group.")
	}

	return "", nil
}

func (a *LDAP) GetGroupUsers(GroupName string, conn *ldap.Conn) {
	_, _ = a.Group(GroupName, conn)
}

func (a *LDAP) GetSearchRequest(username string) *ldap.SearchRequest {
	//Filter:=fmt.Sprintf("(sAMAccountName=%s)", username)
	Filter := fmt.Sprintf(a.Option.AuthFilter, username)
	Attributes := a.Option.Attributes
	if a.Option.AuthFilter != defaultFilter {
		Attributes = OpenldapAttributes
	}
	searchRequest := ldap.NewSearchRequest(a.Option.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, // ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, ldap.DerefAlways
		0,                      // SizeLimit: 大小设置,一般设置为0
		0,                      // TimeLimit: 时间设置,一般设置为0
		false,                  //TypesOnly:  设置false（返回的值要多一点）
		Filter,                 // "(objectClass=*)",
		Attributes,             //Attributes 需要返回的属性值
		nil)                    //Controls:  控制
	return searchRequest
}

func (a *LDAP) RestPassword(username, OldPassword, NewPassword string, change bool) error {
	conn, err := a.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	if change {
		err = conn.Bind(a.Option.Username, a.Option.Password)
		if err != nil {
			//panic("Password error.")
			return err
		}
	}

	entry, err := a.User(username, conn)
	if err != nil {
		return err
	}

	//var userDn = fmt.Sprintf("CN=%s,OU=group,%s", username, "base_dn")

	sql2 := ldap.NewModifyRequest(entry.DN, nil)

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16.NewEncoder().String("\"" + NewPassword + "\"")
	if err != nil {
		hlog.Errorf("ldap user[%s] new password encoder failed.", username)
		return err
	}
	sql2.Replace("unicodePwd", []string{pwdEncoded})
	//sql2.Replace("userAccountControl", []string{"512"})

	if err = conn.Modify(sql2); err != nil {
		hlog.Errorf("ldap user[%s] password modify failed.", username)
		return err
	}
	return nil
}

func GetLdapApi(Enable, Tls bool, Addr, BaseDN, ManagerDN, BindPass, AuthFilter string, logger *logrus.Logger) (*LDAP, error) {
	if !Enable || len(Addr) == 0 || len(ManagerDN) == 0 {
		return nil, fmt.Errorf("ldap data is none")
	}
	var ldapApi = NewLDAP()
	ldapApi.Logger = logger
	ldapApi.Option.Addr = Addr
	ldapApi.Option.AuthFilter = AuthFilter
	ldapApi.Option.BaseDN = BaseDN
	ldapApi.Option.Username = ManagerDN
	ldapApi.Option.Password = BindPass
	ldapApi.Option.Tls = Tls
	return ldapApi, nil
}
