package router

import (
	"admin-service/controller"
	"net/http"
)

var ArtileRouterGroup = Group{
	Prefix: "/admin/article",
	RouterChild: []Router{
		{
			Path: "/:id",
			Method: http.MethodGet,
			Handler: controller.Article.GetInfo,
			IsAuth: false,
		},
	},
}
