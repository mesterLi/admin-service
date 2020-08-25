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
			IsAuth: false,
		},
		{
			Path: "/info",
			Method: http.MethodGet,
			Handler: controller.Info,
			IsAuth: true,
		},
		{
			Path: "/create",
			Method: http.MethodPost,
			Handler: controller.Create,
			IsAuth: true,
		},
	},
}