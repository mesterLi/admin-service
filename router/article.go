package router

import (
	"admin-service/controller"
	"net/http"
)

var ArtileRouterGroup = Group{
	Prefix: "/admin/article",
	RouterChild: []Router{
		{
			Path: "/list",
			Method: http.MethodGet,
			Handler: controller.Article.List,
		},
		{
			Path: "/list/:id",
			Method: http.MethodGet,
			Handler: controller.Article.GetInfo,
		},
	},
}
