package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func redisConnect() (rdb *redis.Client) {
	//rdb = Client
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return
}

func subMessage(channel string) {
	rdb := redisConnect()
	pubsub := rdb.Subscribe(context.Background(), channel)
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		panic(any(err))
	}

	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func Pub() {
	var ctx = context.Background()
	rdp := redisConnect()
	n, err := rdp.Publish(ctx, "reload_config", "hello ocean ").Result()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("%d clients received the message\n", n)
}

func Sub() {
	var ctx = context.Background()
	rdp := redisConnect()
	pubsub := rdp.Subscribe(ctx, "reload_config")
	defer pubsub.Close()
	var c = time.After(123 * time.Second)
	for {
		select {
		case <-c:
			fmt.Println("已经超时")
			return
		case <-ctx.Done():
			fmt.Println("已经结束")
			return
		case msg := <-pubsub.Channel():
			fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
			//for msg := range pubsub.Channel() {
			//	fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
			//}
		}
	}
}

func redisDemo() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "xxxx:6379",
		Password: "xxxx", // no password set
		DB:       0,      // use default DB
	})
	/*_ = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        []string{"localhost:6379", "127.0.0.1:6379"},
		Password:     "",
		MaxRedirects: 8, // 当遇到网络错误或者MOVED/ASK重定向命令时，最多重试几次，默认8

		//只含读操作的命令的"节点选择策略"。默认都是false，即只能在主节点上执行。
		ReadOnly: false, // 置为true则允许在从节点上执行只含读操作的命令
		// 默认false。 置为true则ReadOnly自动置为true,表示在处理只读命令时，可以在一个slot对应的主节点和所有从节点中选取Ping()的响应时长最短的一个节点来读数据
		RouteByLatency: false,
		// 默认false。置为true则ReadOnly自动置为true,表示在处理只读命令时，可以在一个slot对应的主节点和所有从节点中随机挑选一个节点来读数据
		RouteRandomly: false,

		//每一个redis.Client的连接池容量及闲置连接数量，而不是cluterClient总体的连接池大小。实际上没有总的连接池
		//而是由各个redis.Client自行去实现和维护各自的连接池。
		PoolSize:     15, // 连接池最大socket连接数，默认为5倍CPU数， 5 * runtime.NumCPU
		MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		//命令执行失败时的重试策略
		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时，-1表示取消读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		//闲置连接检查包括IdleTimeout，MaxConnAge
		IdleCheckFrequency: 60 * time.Second, //闲置连接检查的周期，无默认值，由ClusterClient统一对所管理的redis.Client进行闲置连接检查。初始化时传递-1给redis.Client表示redis.Client自己不用做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		IdleTimeout:        5 * time.Minute,  //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:         0 * time.Second,  //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
	})*/

	var ctx = context.Background()
	setResult, err := rdb.SetNX(ctx, "deployment_dev_k8s_godemo", "deploying", 20*time.Second).Result()
	if err != nil {
		fmt.Println("setnx failed:", err)
		return
	}
	fmt.Println("setResult:", setResult)

	time.Sleep(1 * time.Second)

	_, _ = rdb.Set(ctx, "aaaa", []byte("dasjflkdajsflkjlk dakfjlk daf"), 20*time.Second).Result()
	r1, _ := rdb.Get(ctx, "aaaa").Result()
	r2, _ := rdb.Get(ctx, "aaaa2").Result()
	fmt.Println("r1:", r1, len(r1))
	fmt.Println("r2:", r2, len(r2))

	setResult, err = rdb.SetNX(ctx, "key1", "value", 20*time.Second).Result()
	if err != nil {
		fmt.Println("setnx failed:", err)
		return
	}
	fmt.Println("setResult:", setResult)
}
