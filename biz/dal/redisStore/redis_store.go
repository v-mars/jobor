package redisStore

import (
	"fmt"
	"jobor/conf"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
)

var Store redis.Store

func Init() {
	Store = Connect()
}

func Connect() redis.Store {
	store, err := redis.NewStoreWithDB(10, "tcp",
		conf.GetConf().Redis.Address,
		conf.GetConf().Redis.Password,
		fmt.Sprintf("%d", conf.GetConf().Redis.Db),
		[]byte(conf.GetConf().Authentication.AuthSecret))
	if err != nil {
		hlog.Fatalf("session init redis err, %s", err)
	}
	store.Options(GetSessionOption())
	return store
}

func GetSessionOption() sessions.Options {
	return sessions.Options{
		Path:   "/",
		MaxAge: conf.GetConf().Authentication.MaxAge,
	}
}

func GetRedisStore() (rediStore *redis.RediStore, err error) {
	rediStore, err = redis.GetRedisStore(Store)
	if err != nil {
		return nil, err
	}
	return rediStore, err
}
