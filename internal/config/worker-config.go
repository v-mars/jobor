package config

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"log"
)

var WorkerConfig = Worker{
	Name: getDefaultName(),
	Mode:     "DEBUG",
	IP:       "0.0.0.0",
	Port:     20052,
	Weight:   100,
	LogPath:  "./logs",
	LogLevel: "DEBUG",
	RoutingKey: "default",
	Servers:  make([]string,0),
	Version: "worker version 1.0.1",
}


type Worker struct {
	Name         string `json:"name"`
	Mode         string `json:"mode"`
	Addr         string `json:"addr"`
	IP           string `json:"ip"`
	Port         int32 `json:"port"`
	Weight       int32 `json:"weight"`
	RoutingKey   string `json:"routing_key" `
	LogPath      string `json:"log_path"`
	LogLevel     string `json:"log_level"`
	Servers      []string `json:"servers"`
	Version      string `json:"version"`
}

// LoadWorkerConfig 加载配置
func LoadWorkerConfig(configPath string) (err error) {
	var config struct{
		Worker *Worker
	}
	var c = config
	c.Worker = &WorkerConfig
	//var c = &WorkerConfig

	if _, err = toml.DecodeFile(configPath, &c); err != nil {
		log.Fatal("Worker加载配置失败:", err)
	}

	if len(c.Worker.Mode) == 0 {
		c.Worker.Mode = gin.ReleaseMode
	}

	if len(c.Worker.IP) == 0 {
		c.Worker.IP = "0.0.0.0"
	}

	if c.Worker.Port == 0 {
		c.Worker.Port = 20052
	}
	return nil
}