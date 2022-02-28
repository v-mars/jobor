package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sync"
)

var Configs = Config{
	Server: Server{
		Name:     name,
		Mode:     "DEBUG",
		IP:       "0.0.0.0",
		BindHost: "0.0.0.0",
		Port:     "5000",
		RootPath: rootPath,
		LogPath:  "./logs",
		LogLevel: "DEBUG",
		GRPCPort: 50052,
		CronType: "s",
	},
	JWT: JWT{
		TokenType: "Bearer",
		TokenKey: "",
		RefreshKey: "",
		Age: 3600,
	},
	Ldap: Ldap{},
	Raft: Raft{
		Bootstrap: true,
		HttpAddress: "127.0.0.1:2869",
		TcpAddress: "127.0.0.1:2889",
		DataDir: "./raft_data",
	},
}

func getDefaultName() string {
	hostname, _ := os.Hostname()
	hostRune := []rune(hostname)
	if len(hostRune) > 32 {
		hostRune = hostRune[:32]
	}
	return string(hostRune)
}

func SetConf(conf Config) {
	lock.Lock()
	defer lock.Unlock()
	Configs = conf
}

func GetConf() Config {
	lock.RLock()
	defer lock.RUnlock()
	return Configs
}

var lock = new(sync.RWMutex)
var name = getDefaultName()
var rootPath, _ = os.Getwd()


// Config 配置参数
type Config struct {
	Web         Web
	Server      Server
	Gorm        Gorm
	MySQL       MySQL
	Sqlite3     Sqlite3
	Redis       Redis
	Session     Session
	JWT 		JWT
	Email		Email
	Ldap		Ldap
	Raft		Raft
}

// Web 站点配置参数
type Web struct {
	Domain       string
	StaticPath   string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type Raft struct {
	Bootstrap   bool   `json:"bootstrap"`
	HttpAddress string `json:"httpAddress"`
	TcpAddress  string `json:"tcpAddress"`
	DataDir     string `json:"dataDir"`
	JoinAddress string `json:"joinAddress"`
}


type Server struct {
	Name         string  `json:"Name"`
	ConfigFile   string  `json:"configFile"`
	RootPath     string
	Mode         string
	IP           string
	BindHost     string
	Port         string
	GRPCPort     int

	CronType     string
	LogPath      string
	LogLevel     string

	GitlabHost   string
	GitlabToken  string

	JenkinsUser  string
	JenkinsToken string
	JenkinsHost  string
	JenkinsJob   string
}

type Ldap struct {
	Addr       string   `json:"addr"`
	BaseDn     string   `json:"baseDn"`
	BindDn     string   `json:"bindDn"`
	BindPass   string   `json:"bindPass"`
	AuthFilter string   `json:"authFilter"`
	Attributes []string `json:"attributes"`
	Domain     string   `json:"domain"`
	Tls        bool     `json:"tls"`
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool
	DBType       string
	MaxLifetime  int
	MaxOpenConn  int
	MaxIdleConn  int
	TablePrefix  string
	DSN          string
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
	Charset    string   `default:"utf8mb4"`
}

// DSN MySQL 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

func (a MySQL) CasbinDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", a.User, a.Password, a.Host, a.Port, a.DBName)
}

// Sqlite3 配置参数
type Sqlite3 struct {
	Path string
}

// DSN Sqlite3 数据库连接串
func (a Sqlite3) DSN() string {
	return a.Path
}

type DataBases struct {
	Engine     string            `json:"engine" default:"mysql"`
	Name       string            `json:"name" default:"ops"`
	User       string            `json:"user" default:"root"`
	Password   string            `json:"password" default:"123456"`
	Host       string            `json:"host" default:"127.0.0.1"`
	Port       string            `json:"port" default:"3306"`
	Charset    string            `json:"charset" default:"utf8mb4"`
	Parameters string            `json:"parameters"`
	Options    map[string]string `json:"options"`
}

func (a DataBases) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.Name, a.Parameters)
}

type Redis struct {
	Host        string
	Port        int
	Password    string
	DB          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type Session struct {
	SessionCookieAge            int      // 设置session失效时间,单位为秒
	SessionSaveEveryRequest     bool     // 是否每次请求都保存Session，默认修改之后才保存（默认）
	SessionExpireAtBrowserClose bool     // 是否关闭浏览器使得Session过期（默认）
	SessionCookieName           string
	SessionPrefix               string
	SessionEngine               string
}

type JWT struct {
	TokenType    string		// AUTHORIZATION  Bearer Authorization
	TokenKey     string
	RefreshKey   string
	Age          int
}

type Email struct {
	SMTPHost   string `json:"smtpHost"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	From       string `json:"from"`
	Tls        bool   `json:"tls"`
	Anonymous  bool   `json:"anonymous"`
	SkipVerify bool   `json:"skipVerify"`
}

// LoadConfig 加载配置
func LoadConfig(configPath string) (err error) {
	var c = &Configs
	//var server = c.Server

	if _, err := toml.DecodeFile(configPath, &c); err != nil {
		log.Fatal("加载配置失败:", err)
	}
	//if server.Mode != "" {
	//	c.Server.Mode = server.Mode
	//} else if len(c.Server.Mode) == 0 {
	//	c.Server.Mode = gin.ReleaseMode
	//}

	//if server.IP != "" {
	//	c.Server.IP = server.IP
	//} else if len(c.Server.IP) == 0 {
	//	c.Server.IP = "0.0.0.0"
	//}

	//if server.Port != "" {
	//	c.Server.Port = server.Port
	//} else if len(c.Server.Port) == 0 {
	//	c.Server.Port = "8000"
	//}
	//fmt.Println("ac:", *ac)
	//jsonBytes, _ := json.Marshal(*ac)
	//fmt.Println("ac:", string(jsonBytes))
	if len(c.MySQL.Parameters) == 0 {
		c.MySQL.Parameters = "charset=utf8&parseTime=True&loc=Asia%2FShanghai"
	}


	//c.Sqlite3.Path = ""

	c.Session.SessionCookieAge = 60 * 60
	c.Session.SessionExpireAtBrowserClose = true
	c.Session.SessionSaveEveryRequest = true
	c.Session.SessionPrefix = ""
	c.Session.SessionCookieName = "sessionid"
	c.Session.SessionEngine = "redis"
	if c.JWT.TokenType == "" {
		c.JWT.TokenType = "Bearer"
	}
	if c.JWT.TokenKey == "" {
		c.JWT.TokenKey = "jS2SnJdxxmTKRNQYh"
	}
	if c.JWT.RefreshKey == "" {
		c.JWT.RefreshKey = "jreJdXxT0kenYh"
	}
	if c.JWT.Age == 0 {
		c.JWT.Age = 60 * 60
	}
	return nil
}

const Channel = "jobor.config"

type ReloadConfig struct {
	Op string
	Target string
	Data interface{}
}




