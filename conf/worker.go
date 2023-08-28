package conf

import (
	"flag"
	"github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	wconf *WorkerConfig
)

// GetWorkerConf gets configuration instance
func GetWorkerConf() *WorkerConfig {
	once.Do(initWorkerConf)
	return wconf
}

type WorkerConfig struct {
	HttpAddress   string   `yaml:"http_address"`
	RoutingKey    string   `yaml:"routing_key"`
	Weight        int      `yaml:"weight"`
	Servers       []string `yaml:"servers"`
	LogLevel      string   `yaml:"log_level"`
	LogFileName   string   `yaml:"log_file_name"`
	LogMaxSize    int      `yaml:"log_max_size"`
	LogMaxBackups int      `yaml:"log_max_backups"`
	LogMaxAge     int      `yaml:"log_max_age"`
	Concurrency   int      `yaml:"concurrency"`
}

func initWorkerConf() {
	confFileRelPath := FlagWorker
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	wconf = new(WorkerConfig)
	wconf.HttpAddress = ":2002"
	wconf.LogLevel = "./log/std.log"
	wconf.LogLevel = "debug"
	wconf.Concurrency = 20
	wconf.Weight = 100
	wconf.RoutingKey = "default"
	err = yaml.Unmarshal(content, wconf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err = validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}

	//pretty.Printf("%+v\n", conf)
}

var FlagWorker string

func init() {
	flag.StringVar(&FlagWorker, "wc", "conf/worker.yaml", "config path, eg: -conf conf/worker.yaml")
	flag.StringVar(&FlagWorker, "wconf", "conf/worker.yaml", "config path, eg: -conf conf/worker.yaml")
	flag.Parse()
}
