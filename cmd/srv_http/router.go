// Code generated by hertz generator.

package srv_http

import (
	"context"
	"embed"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/hertz-contrib/pprof"
	"html/template"
	"io/fs"
	"jobor/biz/handler"
	"jobor/biz/handler/tool"
	"jobor/biz/mw"
	"jobor/biz/pack/oidc_callback"
	"jobor/biz/response"
	"jobor/conf"
	_ "jobor/docs"
	"jobor/fsembed"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.GET("/panic", handler.Panic)
	r.GET(oidc_callback.GotoRedirect, oidc_callback.RedirectLogin)
	r.GET(oidc_callback.CallbackPath, oidc_callback.SsoCallback)
	r.GET("/routes", func(ctx context.Context, c *app.RequestContext) {
		routeInfo := r.Routes()
		hlog.Info(routeInfo)
		type res struct {
			Method  string `json:"method"`
			Path    string `json:"path"`
			Handler string `json:"handler"`
		}
		var data []res
		for _, v := range routeInfo {
			data = append(data, res{Method: v.Method, Path: v.Path, Handler: v.Handler})
		}
		response.SendDataResp(ctx, c, response.Succeed, data)
	})
	if conf.GetConf().Server.EnableSwagger {
		r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, swagger.DocExpansion("none")))
	}
	if conf.GetConf().Server.EnablePprof {
		pprof.Register(r, "api/v1/jobor/debug/pprof")
	}

	r.GET("/api/v1/login-init", mw.LoginInit)
	r.POST("/api/v1/login", mw.AuthWm.LoginHandler)
	r.GET("/api/v1/login", mw.AuthWm.LoginHandler)
	r.POST("/api/v1/logout", mw.AuthWm.LogoutHandler)
	r.GET("/api/v1/logout", mw.AuthWm.LogoutHandler)
	r.GET("/api/v1/jobor/migrate", tool.GetMigrate)
	r.GET("/api/v1/jobor/state-code", tool.GetStateCode)
	//r.POST("/api/v1/sys/gen-token", tool.GenJwtToken)
	r.POST("/api/v1/refresh_token", mw.AuthWm.RefreshHandler)
	r.GET("/api/v1/refresh_token", mw.AuthWm.RefreshHandler)
	//r.StaticFS("/", &app.FS{Root: "./fs/dist/", GenerateIndexPages: true})
	//setStaticWeb(r)
	setStaticFsWeb(r)

	auth := r.Group("/auth")
	auth.Use(mw.AuthWm.MiddlewareFunc())
	{
		auth.GET("/ping", handler.Ping)
	}

	// 服务api更新
	if conf.GetConf().Server.AutoUpdateApi {
		//api.UpdateApiByRoutes(db.DB, r)
	}
}

func setStaticWeb(e *server.Hertz) {
	e.LoadHTMLGlob("./fs/dist/*.html")
	prefix := "/"
	root := "./fs/dist/"
	fss := &app.FS{Root: root, PathRewrite: getPathRewriter(prefix), IndexNames: []string{"index.html"}}
	e.StaticFS(prefix, fss)
	//e.Static("/static/", "./fs/dist/static")
	//e.StaticFile("/favicon.ico", "./fs/dist/favicon.ico")
	e.GET("/jobor/*any", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "index.html", "")
	})
	e.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "index.html", "")
	})
}

func useEmbedWeb(h *server.Hertz) {
	//var file embed.FS
	var file = fsembed.DistFs
	h.StaticFS("/", &app.FS{
		Root:               "fs/dist",              // 根目录
		IndexNames:         []string{"index.html"}, // 索引文件
		GenerateIndexPages: true,                   // 生成索引页面
		Compress:           false,                  // 压缩
		AcceptByteRange:    false,                  // 接受字节范围
		PathRewrite:        nil,                    // 路径重写
		PathNotFound: func(ctx context.Context, c *app.RequestContext) {
			path := string(c.Path())
			// 这个函数是路径找不到绘执行这个，例如 /login /home
			// css js 文件会在 Root 文件 里面找
			// 下面匹配路径
			switch {
			case strings.HasSuffix(path, ".js"):
				return
			case strings.HasSuffix(path, ".css"):
				return
			default:
				// 必须有这一步 react vue 项目 不是 '/'路径 刷新后找不到 报 404
				// 上面匹配 js css 可以不要，样式文件 会在 Root 文件 里面找，找不到会执行上面的函数
				data, err := file.ReadFile("fs/dist/index.html") // 读取react vue 项目的 index.html
				if err != nil {
					//result.HttpError(c, xerr.ErrMsg(xerr.FIleNotExist))
					return
				}
				c.Data(200, "text/html; charset=utf-8", data)

			}
		}, // 路径未找到
		CacheDuration:        0,  // 缓存持续时间
		CompressedFileSuffix: "", // 压缩文件后缀
	})
}

func getPathRewriter(prefix string) app.PathRewriteFunc {
	// Cannot have an empty prefix
	if prefix == "" {
		prefix = "/"
	}
	// Prefix always start with a '/' or '*'
	if prefix[0] != '/' {
		prefix = "/" + prefix
	}

	// Is prefix a direct wildcard?
	isStar := prefix == "/*"
	// Is prefix a partial wildcard?
	if strings.Contains(prefix, "*") {
		isStar = true
		prefix = strings.Split(prefix, "*")[0]
		// Fix this later
	}
	prefixLen := len(prefix)
	if prefixLen > 1 && prefix[prefixLen-1:] == "/" {
		// /john/ -> /john
		prefixLen--
		prefix = prefix[:prefixLen]
	}
	return func(ctx *app.RequestContext) []byte {
		path := ctx.Path()
		if len(path) >= prefixLen {
			if isStar && string(path[0:prefixLen]) == prefix {
				path = append(path[0:0], '/')
			} else {
				path = path[prefixLen:]
				if len(path) == 0 || path[len(path)-1] != '/' {
					path = append(path, '/')
				}
			}
		}
		if len(path) > 0 && path[0] != '/' {
			path = append([]byte("/"), path...)
		}
		return path
	}
}

var (
	//prefix = fmt.Sprintf("/%s", conf.Dom)
	prefix = ""
)

func setStaticFsWeb(h *server.Hertz) {
	h.GET(prefix+"/static/*any", distFs)
	h.GET(prefix+"/assets/*any", distFs)
	h.GET(prefix+"/iconfont.js", distFs)
	h.SetHTMLTemplate(template.Must(template.New("").ParseFS(fsembed.DistFs, "dist/*.html")))
	h.GET(fmt.Sprintf("/%s/*any", conf.Dom), func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	h.GET(fmt.Sprintf("/%s/login", conf.Dom), func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	h.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

func distFs(ctx context.Context, c *app.RequestContext) {
	var static = &webUi{embed: fsembed.DistFs, path: "dist"}
	staticServer := http.FileServer(http.FS(static))
	req, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
		return
	}
	staticServer.ServeHTTP(adaptor.GetCompatResponseWriter(&c.Response), req)
}

// 嵌入普通的静态资源
type webUi struct {
	embed embed.FS // 静态资源
	path  string   // 设置embed文件到静态资源的相对路径，也就是embed注释里的路径
}

// Open 静态资源被访问的核心逻辑
func (w *webUi) Open(name string) (fs.File, error) {
	name = strings.Trim(name, prefix)
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, fmt.Errorf("http: invalid character in file path")
	}
	fullName := filepath.Join(w.path, filepath.FromSlash(path.Clean("/"+name)))
	file, err := w.embed.Open(fullName)
	return file, err
}
