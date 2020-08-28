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
			Handler: controller.User.Login,
			IsAuth: false,
		},
		{
			Path: "/create",
			Method: http.MethodPost,
			Handler: controller.User.Create,
			IsAuth: true,
		},
		{
			Path: "/:uid",
			Method: http.MethodPut,
			Handler: controller.User.Update,
		},
		{
			Path: "/:uid",
			Method: http.MethodGet,
			Handler: controller.User.Info,
			IsAuth: true,
		},
		{
			Path: "/:uid",
			Method: http.MethodDelete,
			Handler: controller.User.Delete,
		},
	},
}