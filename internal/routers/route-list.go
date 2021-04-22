package routers

import (
	"github.com/gin-gonic/gin"
	"jobor/internal/response"
)

func Query(c *gin.Context)  {
	type res struct{
		Method      string `json:"method"`
		Path        string `json:"path"`
		Handler     string `json:"handler"`
	}
	var data []res
	routes := Engine.Routes()
	for _,v := range routes{
		data = append(data, res{Method: v.Method, Path: v.Path, Handler: v.Handler})
	}
	response.Success(c, data)
}

