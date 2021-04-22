package casbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.Enforcer

type Casbin interface {
	GetContext() context.Context
	GetEnforcer() *casbin.Enforcer
}

type service struct {
	context  context.Context
	enforcer *casbin.Enforcer
	watcher  persist.Watcher
}

var getCasbin *service

func GetCasbin() *service {
	return getCasbin
}

/*func NewCasbin(db *gorm.DB, rds redis.RedisInterface) (Casbin, error) {
	m := casbin.NewModel()
	// [request_definition]
	m.AddDef("r", "r", "sub, obj, act")

	//[policy_definition]
	m.AddDef("p", "p", "sub, obj, act")

	//[role_definition]
	m.AddDef("g", "g", "_, _")

	// [policy_effect]
	m.AddDef("e", "e", "some(where (p.eft == allow))")

	// [matchers]
	m.AddDef("m", "m", "r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch3(r.obj, p.obj)) && keyMatch2(r.act, p.act)")

	watcher, err := wrds.NewWatcher(rds)
	if err != nil {
		return nil, err
	}

	adapter := gormAdapter.NewAdapterByDB(db)
	e, err := casbin.NewEnforcerSafe(m, adapter)
	if err != nil {
		return nil, err
	}

	e.AddFunction("keyMatch3", util.KeyMatch3Func)

	e.SetWatcher(watcher)
	_ = watcher.SetUpdateCallback(func(string) {
		if err = e.LoadPolicy(); err != nil {
			_ = fmt.Errorf("err: %v", err)
		}
	})

	//ctx := context.WithValue(context.Background(), kitcasbin.CasbinModelContextKey, m)
	//ctx = context.WithValue(ctx, kitcasbin.CasbinPolicyContextKey, adapter)
	//ctx = context.WithValue(ctx, kitcasbin.CasbinEnforcerContextKey, e)
	ctx := context.WithValue(context.Background(), kitcasbin.CasbinEnforcerContextKey, e)

	//e.EnableLog(cf.GetBool("server", "debug"))
	e.EnableLog(true)

	getCasbin = &service{context: ctx, enforcer: e}
	return getCasbin, nil

}*/

func (c *service) GetContext() context.Context {
	return c.context
}

func (c *service) GetEnforcer() *casbin.Enforcer {
	return c.enforcer
}

func (c *service) Close() {
	c.watcher.Close()
}

////a URL path or a * pattern like /alice_data/*
//func KeyMatch(key1 string, key2 string) bool {
//	i := strings.Index(key2, "*")
//	if i == -1 {
//		return key1 == key2
//	}
//
//	if len(key1) > i {
//		return key1[:i] == key2[:i]
//	}
//	return key1 == key2[:i]
//}
//
//// a URL path or a : pattern like /alice_data/:resource
//func KeyMatch2(key1 string, key2 string) bool {
//	re := regexp.MustCompile(`(.*):[^/](.*)`)
//	for {
//		if !strings.Contains(key2, "/:") {
//			break
//		}
//
//		key2 = re.ReplaceAllString(key2, "$1[^/]+$2")
//	}
//
//	return RegexMatch(key1, key2)
//}
//
//
//// a URL path or a {} pattern like /alice_data/{resource}
//func KeyMatch3(key1 string, key2 string) bool {
//	re := regexp.MustCompile(`(.*)\{[^/]+\}(.*)`)
//	for {
//		if !strings.Contains(key2, "/{") {
//			break
//		}
//
//		key2 = re.ReplaceAllString(key2, "$1[^/]+$2")
//	}
//
//	return RegexMatch(key1, key2)
//}
//
//func RegexMatch(key1 string, key2 string) bool {
//	res, err := regexp.MatchString(key2, key1)
//	if err != nil {
//		panic(err)
//	}
//	return res
//}

func NewCasbin(dbSourceName string) (*casbin.Enforcer,error) {
	//dbSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "ocean", "123456","127.0.0.1","3306", "ops")
	adapter, err := gormAdapter.NewAdapter("mysql", dbSourceName, true) // Your driver and data source.
	if err!=nil {return nil, err}

	m := model.NewModel()
	// [request_definition]
	m.AddDef("r", "r", "sub, dom, obj, act")			// "sub, dom, obj, act" user/role 域/租户 url_path	 Method

	//[policy_definition]
	m.AddDef("p", "p", "sub, dom, obj, act")			// "sub, dom, obj, act" user/role 域/租户  url_path	 Method

	//[role_definition]
	m.AddDef("g", "g", "_, _, _")					// "_, _, _"  [user/role  user/role  域/租户]
	//第三个 _ 表示域/租户的名称, 此部分不应更改。 然后，政策可以是：
	//
	//p, admin, tenant1, data1, read
	//p, admin, tenant2, data2, read
	//
	//g, alice, admin, tenant1
	//g, alice, user, tenant2

	// [policy_effect]
	m.AddDef("e", "e", "some(where (p.eft == allow))")

	// [matchers]
	//m.AddDef("m", "m", "r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch3(r.obj, p.obj)) && keyMatch2(r.act, p.act)")
	//var matcherValueV1 = "g(r.sub, p.sub) && (keyMatch2(r.obj, p.obj) || keyMatch3(r.obj, p.obj)) && " +
	//	"keyMatch2(r.act, p.act)"
	var matcherValueV1 = "g(r.sub, p.sub, r.dom) && r.dom == p.dom && (keyMatch2(r.obj, p.obj) || keyMatch3(r.obj, p.obj)) && " +
		"r.act == p.act"
	//m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && keyMatch2(r.dom, p.dom) && r.obj == p.obj && r.act == p.act
	//var matcherValueV1 = "g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.obj"
	m.AddDef("m", "m", matcherValueV1)

	e, err := casbin.NewEnforcer(m, adapter)
	if err!=nil {return nil, err}

	//e.EnableLog(true)


	// Load the policy from DB.
	err = e.LoadPolicy()
	if err!=nil {return nil, err}

	// Load the Model from DB.
	//err = e.LoadModel()
	//if err!=nil {return nil, err}

	// Check the permission.
	//var dom = "sys"
	//if len(dom)==0{dom="sys"}
	//ok1, _ := e.Enforce("admin", dom,"/api/v1/sys/user", "GET")
	//ok2, _ := e.Enforce("root1", dom,"/api/v1/sys/user", "GET")
	//fmt.Println("ok1:", ok1)
	//fmt.Println("ok2:", ok2)
	// Modify the policy.
	//_, _ = e.AddNamedGroupingPolicy("r", "sub, obj, act")
	//_, _ = e.AddPolicy("admin", dom,"/api/v1/sys/user", "GET")
	//_, _ = e.AddPolicy("admin", dom,"/api/v1/sys/user", "POST")
	////_, _ = e.Add("admin", "/api/v1/sys/user", "POST")
	//_, _ = e.AddGroupingPolicy("admin", "role1", dom)
	//_, _ = e.AddGroupingPolicy("root", "admin", dom)

	// e.RemovePolicy(...)
	//roles, err := e.GetImplicitRolesForUser("root", dom)
	//permissions, _ := e.GetImplicitPermissionsForUser("admin", dom)
	//fmt.Println("roles:", roles, err)
	//fmt.Println("permissions:", permissions)

	// Save the policy back to DB.
	//err = e.SavePolicy()
	//if err!=nil {return nil, err}
	Enforcer =e
	return e, nil
}