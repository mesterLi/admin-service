package controller

import (
	"admin-service/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type idParam struct {
	Id string `uri:"id" binding:"required,uuid"`
}
const (
	ARTICLE_STATUS_ALL = iota
	ARTICLE_STATUS_ON
	ARTICLE_STATUS_OFF
	ARTICLE_STATUS_REMOVE
)
type articleControl struct {
	List, GetInfo gin.HandlerFunc
}

var Article = articleControl{
	List: list,
	GetInfo: getInfo,
}
func list(c *gin.Context) {
	fmt.Println(c.FullPath())
	page := c.Param("page")
	limit := c.Param("limit")
	tag := c.Param("tag")
	status := c.Param("status")
	fmt.Println(page,limit,tag,status)
	fmt.Println(ARTICLE_STATUS_ALL, ARTICLE_STATUS_OFF, ARTICLE_STATUS_ON, ARTICLE_STATUS_REMOVE)
	list := []model.ArticleInfo{}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": map[string] interface{}{
			"list": list,
			"total": 0,
			"limit": 10,
		},
		"message": "获取列表成功",
	})
}

func getInfo(c *gin.Context) {
	var param idParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": model.ArticleInfo{},
		"id": c.Param("id"),
	})
}
