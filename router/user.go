package router

import (
	"admin-service/controller"
	"net/http"
)

var UserRouterGroup = Group{
	Prefix: "/admin/user",
	RouterChild: []Router{
		{
			Path: "/login",
			Method: http.MethodPost,
			Handler: controller.Login,
		},
		{
			Path: "/info",
			Method: http.MethodGet,
			Handler: controller.Info,
		},
		{
			Path: "/create",
			Method: http.MethodPost,
			Handler: controller.Create,
		},
	},
}