package router

import (
	"admin-service/global"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Path string
	Method string
	Handler gin.HandlerFunc
}

type Group struct {
	Prefix string
	RouterChild []Router
}

var routes = []Group{
	UserRouterGroup,
	ArtileRouterGroup,
}

func bindRoute(group Group) {
	g := global.Server.Group(group.Prefix)
	for _, router := range group.RouterChild {
		g.Handle(router.Method, router.Path, router.Handler)
	}
}
func Init() {
	for _, group := range routes {
		bindRoute(group)
	}
}