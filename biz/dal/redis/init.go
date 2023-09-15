package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/conf"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var ResultNil = redis.Nil

const (
	FailLoginPrefix   = "fail_login_"
	FailCaptchaPrefix = "captcha_"
)

func Init(c *conf.Config) {
	Rdb = InitClient(c)
}

func Publish(ctx context.Context, channel string, message interface{}) error {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("channel [%s] publish the message Marshal err: %s", channel, err)
	}
	n, err := Rdb.Publish(ctx, channel, string(msgBytes)).Result()
	if err != nil {
		return fmt.Errorf("channel [%s] publish the message err: %s\n", channel, err)
	}
	logger.Infof("channel [%s] message is publish success, the %d clients received the message\n", channel, n)
	return nil
}

func Subscribe(ctx context.Context, callFunc func(channel, payload string) error, channels ...string) error {
	pubSub := Rdb.Subscribe(ctx, channels...)
	defer pubSub.Close()
	log.Printf("channel=%s subscribe start\n", channels)
	for {
		select {
		//case <-time.After(123*time.Second):
		//	log.Println("已经超时")
		//	return
		case <-ctx.Done():
			log.Printf("channels=%s消息订阅已经结束.\n", channels)
			return nil
		case msg := <-pubSub.Channel():
			log.Printf("channel=%s recevie message=%s\n", msg.Channel, msg.Payload)
			err := callFunc(msg.Channel, msg.Payload)
			if err != nil {
				log.Printf("channel=%s callFunc err: %s", msg.Channel, err)
				return fmt.Errorf("channel=%s exec callFunc err: %s", msg.Channel, err)
			}
		}
	}
}

// InitClient
/*连接Redis*/
func InitClient(c *conf.Config) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		//Addr:     "0.0.0.0",
		Addr:     c.Redis.Address,
		Password: c.Redis.Password, // no password set
		DB:       c.Redis.Db,       // use default DB
		//MaxRetries: 8,
		//MaxRedirects: 8, // 当遇到网络错误或者MOVED/ASK重定向命令时，最多重试几次，默认8
		DialTimeout: 5 * time.Second,
		ReadTimeout: 5 * time.Second,
	})

	pong, err := client.Ping(context.TODO()).Result()

	if err != nil {
		panic("redis连接失败，请检查配置.\n 错误：" + err.Error())
	}
	if pong != "" {
		hlog.Infof("redis ping is %s, connect success!", pong)
	}
	return client
}

/*连接Redis集群*/
func initClusterClient() (err error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping(context.TODO()).Result()
	if err != nil {
		return err
	}
	return nil
}

/* 连接Redis哨兵模式 */
func initSentinelClient() (err error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping(context.TODO()).Result()
	if err != nil {
		return err
	}
	return nil
}

func NonnullErr(err error) bool {
	return errors.Is(err, ResultNil)
}

// 监视watch_count的值，并在值不变的前提下将其值+1
/*key := "watch_count"
err = client.Watch(func(tx *redis.Tx) error {
	n, err := tx.Get(key).Int()
	if err != nil && err != redis.Nil {
		return err
	}
	_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Set(key, n+1, 0)
		return nil
	})
	return err
}, key)*/
