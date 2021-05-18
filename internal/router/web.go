package router

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	Fs "jobor/fs"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

func setEmbedWeb(e *gin.Engine)  {
	//tmpl := template.Must(template.New("").ParseFS(Fs.IndexFs, "dist/index.html"))
	//Engine.SetHTMLTemplate(tmpl)
	// Vue静态资源
	webuiAntObj := &webUi{
		embed: Fs.DistFs,
		path:  "dist",
	}
	// 设置Ant Design前端访问，try file到index.html
	webuiAntIndexObj := &webUiIndex{
		webUi: webuiAntObj,
	}
	var static=&webUi{embed: Fs.DistFs,path: "dist/static"}
	var favicon=&webUi{embed: Fs.DistFs,path: "dist/favicon.ico"}
	e.StaticFS("/static", http.FS(static))
	e.StaticFS("/favicon.ico/", http.FS(favicon))
	e.GET("/", func(c *gin.Context) {
		c.FileFromFS("index", http.FS(webuiAntIndexObj))
	})
}

func setStaticWeb(e *gin.Engine)  {
	e.LoadHTMLGlob("./fs/dist/*.html")
	e.Static("/static/", "./fs/dist/static")
	e.StaticFile("/favicon.ico", "./fs/dist/favicon.ico")
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}


// 嵌入普通的静态资源
type webUi struct {
	embed embed.FS // 静态资源
	path  string   // 设置embed文件到静态资源的相对路径，也就是embed注释里的路径
}

// Open 静态资源被访问的核心逻辑
func (w *webUi) Open(name string) (fs.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, fmt.Errorf("http: invalid character in file path")
	}
	fullName := filepath.Join(w.path, filepath.FromSlash(path.Clean("/"+name)))
	file, err := w.embed.Open(fullName)
	return file, err
}

// Ant Design前端页面，需要该方式，实现刷新，访问到前端index.html
type webUiIndex struct {
	webUi *webUi
}

func (w *webUiIndex) Open(name string) (fs.File, error) {
	return w.webUi.Open("index.html")
}


