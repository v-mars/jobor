# Jobor分布式定时任务
## ✨ 功能特性
- 通过raft一致性算法，实现多server/controller/master的高可用，不同于传统的分布式只实现worker端的高可用，调度端只能是单点来避免同一任务同一时间被重复调度执行，从而达到了整个服务（Server,Worker）的高可用，保证了系统的健壮稳定性。
- worker高可用，并且通过路由标识，worker可以部署在不同环境，实现不同环境worker的权限控制，worker的系统环境依赖（Python,Golang,执行依赖的文件）。
- 调度server与worker通过grpc通信。
- 支持LDAP（openldap,AD）协议用户认证。
- 支持多种任务脚本 [ api/restful请求, shell, python3 ] e.g
- 基于Casbin实现的权限认证

## 架构图
![avatar](./img/struct.png)

## 构建
```
go build -o ./app ./main.go
go build -o ./app ./cmd/worker/main.go
```

## 启动Server
```
go build -o ./app ./main.go

./app -c ./conf/config.yaml
```


## 启动Worker
```
go build -o ./app ./cmd/worker/main.go
./app -c ./conf/worker.yaml
```

## 默认
username: admin
password: admin

## DB
```
字符集：utf8mb4
在配置文件最后一行粘贴以下语句
/etc/mysql/conf.d/mysql.cnf

[mysqld]
sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION

```


## Jobor预览
![avatar](./img/jobor-dash.jpeg)

## TODO 
### task
- [x] 支持server/controller/master(通过raft一致性算法)的高可用，一个Raft集群通常包含2*N+1个服务器，允许系统有N个故障服务器。
- [x] ldap(支持openldap,AD 认证)
- [x] server <-- gRPC --> worker
- [x] task abort
- [x] task timeout
- [x] api/restful [GET, POST, PUT, DELETE] task
- [x] shell task
- [x] python3 task
- [x] golang task
- [x] server task
- [x] father task
- [x] children task
- [x] worker 预执行（如：执行python 前先执行 pip install xx）
- [x] worker 节点支持：agent和ssh两种模式
- [x] 路由标识多选
- [ ] 任务缓存执行

## 🤝 特别感谢
- golang 1.20
- hertz
- hertz-swagger
- kitex
- gorm
- casbin
- mysql 8.0
- redis 5
- 等

## 支持
1、希望大家多多支持，给项目一个star

2、该项目花费了作者大量时间，如果你觉的该项目对你有用，希望可以友情赞助一下

<img src="./img/wechat.jpeg" width=200 height=200>


## 交流/商务联系
```
如果您只是使用本项目的话，您可以在提出您使用中需要改进的地方，我会尽快修改。
如果您是想基于此项目二次开发的话，您可以提出您在开发过程中的任何疑问，我会尽快答复并讲解。
```
<img src="./img/Wechatid.jpeg" width=200 height=200>


