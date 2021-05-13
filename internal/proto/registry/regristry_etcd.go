package registry

//import (
//	"context"
//	"flag"
//	"fmt"
//	"go.etcd.io/etcd/clientv3"
//	"google.golang.org/service"
//	"net"
//	"os"
//	"os/signal"
//	"strings"
//	"syscall"
//	"time"
//)
//
type Etcd struct {
	Subscriptions
}

func (e Etcd) Registry(port int32, ttl int64) error {
	panic("implement me")
}

func (e Etcd) Keepalive(ttl int64) error {
	panic("implement me")
}

func (e Etcd) UnRegistry() error {
	panic("implement me")
}

//
//const schema = "ns"
//
//var host = "127.0.0.1" //服务器主机
//
//var (
//	Port        = flag.Int("Port", 3000, "listening port")                           //服务器监听端口
//	ServiceName = flag.String("ServiceName", "greet_service", "service name")        //服务名称
//	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address") //etcd的地址
//)
//
//var cli *clientv3.Client
//
////将服务地址注册到etcd中
//func register(etcdAddr, serviceName, serverAddr string, ttl int64) error {
//	var err error
//
//	if cli == nil {
//		//构建etcd client
//		cli, err = clientv3.New(clientv3.Config{
//			Endpoints:   strings.Split(etcdAddr, ";"),
//			DialTimeout: 15 * time.Second,
//		})
//		if err != nil {
//			fmt.Printf("连接etcd失败：%s\n", err)
//			return err
//		}
//	}
//
//	//与etcd建立长连接，并保证连接不断(心跳检测)
//	ticker := time.NewTicker(time.Second * time.Duration(ttl))
//	go func() {
//		key := "/" + schema + "/" + serviceName + "/" + serverAddr
//		for {
//			resp, err := cli.Get(context.Background(), key)
//			//fmt.Printf("resp:%+v\n", resp)
//			if err != nil {
//				fmt.Printf("获取服务地址失败：%s", err)
//			} else if resp.Count == 0 { //尚未注册
//				err = keepAlive(serviceName, serverAddr, ttl)
//				if err != nil {
//					fmt.Printf("保持连接失败：%s", err)
//				}
//			}
//			<-ticker.C
//		}
//	}()
//
//	return nil
//}
//
////保持服务器与etcd的长连接
//func keepAlive(serviceName, serverAddr string, ttl int64) error {
//	//创建租约
//	leaseResp, err := cli.Grant(context.Background(), ttl)
//	if err != nil {
//		fmt.Printf("创建租期失败：%s\n", err)
//		return err
//	}
//
//	//将服务地址注册到etcd中
//	key := "/" + schema + "/" + serviceName + "/" + serverAddr
//	_, err = cli.Put(context.Background(), key, serverAddr, clientv3.WithLease(leaseResp.ID))
//	if err != nil {
//		fmt.Printf("注册服务失败：%s", err)
//		return err
//	}
//
//	//建立长连接
//	ch, err := cli.KeepAlive(context.Background(), leaseResp.ID)
//	if err != nil {
//		fmt.Printf("建立长连接失败：%s\n", err)
//		return err
//	}
//
//	//清空keepAlive返回的channel
//	go func() {
//		for {
//			<-ch
//		}
//	}()
//	return nil
//}
//
////取消注册
//func unRegister(serviceName, serverAddr string) {
//	if cli != nil {
//		key := "/" + schema + "/" + serviceName + "/" + serverAddr
//		cli.Delete(context.Background(), key)
//	}
//}
//
//
//func mainF() {
//	flag.Parse()
//
//	//监听网络
//	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *Port))
//	if err != nil {
//		fmt.Println("监听网络失败：", err)
//		return
//	}
//	defer listener.Close()
//
//	//创建grpc句柄
//	srv := service.NewServer()
//	defer srv.GracefulStop()
//
//	//将greetServer结构体注册到grpc服务中
//	//proto.RegisterGreetServer(srv, &greetServer{})
//
//	//将服务地址注册到etcd中
//	serverAddr := fmt.Sprintf("%s:%d", host, *Port)
//	fmt.Printf("greeting server address: %s\n", serverAddr)
//	register(*EtcdAddr, *ServiceName, serverAddr, 5)
//
//	//关闭信号处理
//	ch := make(chan os.Signal, 1)
//	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
//	go func() {
//		s := <-ch
//		unRegister(*ServiceName, serverAddr)
//		if i, ok := s.(syscall.Signal); ok {
//			os.Exit(int(i))
//		} else {
//			os.Exit(0)
//		}
//	}()
//
//	//监听服务
//	err = srv.Serve(listener)
//	if err != nil {
//		fmt.Println("监听异常：", err)
//		return
//	}
//}