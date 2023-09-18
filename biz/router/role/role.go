// Code generated by hertz generator. DO NOT EDIT.

package role

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	role "jobor/biz/handler/role"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			{
				_sys := _v1.Group("/sys", _sysMw()...)
				_sys.GET("/role", append(_getroleMw(), role.GetRole)...)
				_role := _sys.Group("/role", _roleMw()...)
				_role.PUT("/:id", append(_putroleMw(), role.PutRole)...)
				_role.DELETE("/:id", append(_deleteroleMw(), role.DeleteRole)...)
				_sys.POST("/role", append(_postroleMw(), role.PostRole)...)
				_sys.GET("/role-tree", append(_getroletreeMw(), role.GetRoleTree)...)
				_sys.GET("/roles", append(_getroleallMw(), role.GetRoleAll)...)
			}
		}
	}
}
