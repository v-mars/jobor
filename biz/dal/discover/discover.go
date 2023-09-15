package discover

import (
	"fmt"
	"jobor/conf"
	utils2 "jobor/pkg/utils"
	"log"
	"strings"

	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	hc "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/utils"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"github.com/hertz-contrib/registry/redis"
)

func RegistryWeb(options *[]hc.Option) {
	tags := map[string]string{"app": strings.ToLower(conf.GetConf().Server.Service),
		"group": conf.GetEnv(), "dyeing": conf.GetDyeing()}
	if conf.GetConf().Server.EnableRegistry {
		localIP := utils2.GetOutBoundIP("tcp", conf.GetConf().Redis.Address)
		var r registry.Registry
		//var err error
		switch conf.GetConf().Server.RegistryCenter {
		case "consul":
			localIP = utils2.GetOutBoundIP("tcp", conf.GetConf().Consul.Address)
			cc := consulapi.DefaultConfig()
			cc.Address = conf.GetConf().Consul.Address
			cc.Token = conf.GetConf().Consul.Token
			consulClient, err := consulapi.NewClient(cc)
			if err != nil {
				log.Fatal(err)
				return
			}
			// build a consul register with the consul client
			r = consul.NewConsulRegister(consulClient)
		case "redis":
			localIP = utils2.GetOutBoundIP("tcp", conf.GetConf().Redis.Address)
			r = redis.NewRedisRegistry(conf.GetConf().Redis.Address)
		//case "etcd":
		//	localIP = utils2.GetOutBoundIP("tcp", conf.GetConf().Etcd.HttpAddress)
		//	r, err = etcd.NewEtcdRegistry([]string{conf.GetConf().Etcd.HttpAddress})
		//	if err != nil {
		//		panic(err)
		//	}
		//case "zookeeper", "zk":
		//	localIP = utils2.GetOutBoundIP("tcp", conf.GetConf().Zk.HttpAddress)
		//	r, err = zookeeper.NewZookeeperRegistry([]string{conf.GetConf().Zk.HttpAddress}, 40*time.Second)
		//	if err != nil {
		//		panic(err)
		//	}
		case "default":
			log.Fatalf("尚未支持该注册中心[%s]类型\n", conf.GetConf().Server.RegistryCenter)
		}
		addr := fmt.Sprintf("%s%s", localIP, conf.GetConf().Server.HttpAddress)
		*options = append(*options, server.WithRegistry(r, &registry.Info{
			ServiceName: conf.GetConf().Server.Service,
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        tags,
		}))
	}
}

func RegistryRPC(options *[]hc.Option) {

}
