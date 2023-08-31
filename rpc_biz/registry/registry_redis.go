package registry

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// SUBSCRIPTIONS_

type Redis struct {
	c *redis.Client
	Subscriptions
}

func (m *Redis) Registry(port int32, ttl int64) error {
	var d = time.Second * time.Duration(ttl)
	ticker := time.NewTicker(d)
	go func() {
		for {
			err := m.Keepalive(ttl)
			if err != nil {
				fmt.Printf("保持连接失败：%s", err)
			}
			//fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
			<-ticker.C
		}
	}()

	return nil
}

// 保持服务器与redis的同步更新
func (m *Redis) Keepalive(interval int64) error {
	var key = fmt.Sprintf("routing_key.%s__%s", m.RoutingKey, m.Ip)
	_ = m.c.Set(context.TODO(), key, "aa", time.Second*time.Duration(interval+1))
	return nil
}

func (m *Redis) UnRegistry() error {
	var key = fmt.Sprintf("routing_key.%s__%s", m.RoutingKey, m.Ip)
	_ = m.c.Del(context.TODO(), key)
	return nil
}
