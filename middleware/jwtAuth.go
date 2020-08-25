package middleware

import (
	"admin-service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var white_list = [] string {
	"/admin/user/login",
	"/admin/user/logout",
}
func isLogin(url string) bool {
	var isNeedLogin = true
	for _, u := range white_list {
		if url == u {
			isNeedLogin = false
			break
		}
	}
	return isNeedLogin
}
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAuth, isExist := c.Get("isAuth")
		fmt.Println("isAuth", isExist, isAuth)
		if isLogin(c.Request.RequestURI) == false {
			c.Next()
			return
		}
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"message": "没有的登录信息",
			})
			c.Abort()
			return
		}
		status, err := utils.ValidadteToken(token)
		if err != nil || !status {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"message": "token失效",
			})
			c.Abort()
			return
		}
		var uid int
		utils.GetUidFromToken(token, &uid)
		c.Set("uid", uid)
		c.Next()
	}
}
