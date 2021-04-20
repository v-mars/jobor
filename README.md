# jobor

### 火星计划之定时任务

# 命令
```
./app -h
Usage:
   [command]

Available Commands:
  client      Start Run jobor client
  help        Help about any command
  server      Start Run Jobor Server
  version     Print the version of Jobor

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.

```

# 启动Server
```
 
./app server -h
Welcome User Jobor Server

Usage:
   server [flags]

Examples:
## 启动命令 ./app server -p 5000 -c ./configs/config.toml -f ./logs

Flags:
  -c, --config string   config file, example: ./configs/config.toml
  -h, --help            help for server
  -i, --ip string       服务IP (default "0.0.0.0")
  -l, --level string    日志级别(DEBUG, INFO, WARNING e.g)
  -f, --log string      日志目录(/data/logs e.g) (default "./logs")
  -m, --mode string     启动模式(release, debug, test e.g)
  -p, --port string     服务启动的端口: 5000 (default "5000")

./app server -p 5000 -c ./configs/config.toml -f ./logs
```

# 启动Worker
```
./app worker
```