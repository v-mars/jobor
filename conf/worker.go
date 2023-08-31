package conf

import (
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
	WorkerName    string   `yaml:"worker_name"`
	HttpAddress   string   `yaml:"http_address"`
	GRpcAddr      string   `yaml:"grpc_addr"`
	RoutingKey    string   `yaml:"routing_key"`
	Weight        int      `yaml:"weight"`
	Servers       []string `yaml:"servers"`
	LogLevel      string   `yaml:"log_level"`
	LogFileName   string   `yaml:"log_file_name"`
	LogMaxSize    int      `yaml:"log_max_size"`
	LogMaxBackups int      `yaml:"log_max_backups"`
	LogMaxAge     int      `yaml:"log_max_age"`
	Concurrency   int      `yaml:"concurrency"`
	Redis         Redis    `yaml:"redis"`
}

func initWorkerConf() {
	confFileRelPath := FlagWorker
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	wconf = new(WorkerConfig)
	//wconf.HttpAddress = ":2002"
	wconf.GRpcAddr = ":20021"
	wconf.LogLevel = "./log/std.log"
	wconf.LogLevel = "debug"
	wconf.Concurrency = 20
	wconf.Weight = 100
	err = yaml.Unmarshal(content, wconf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err = validator.Validate(wconf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}
	if wconf.WorkerName == "" {
		wconf.WorkerName, _ = os.Hostname()
	}

	//pretty.Printf("%+v\n", conf)
}

var FlagWorker = "conf/worker.yaml"
