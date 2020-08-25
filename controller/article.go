package controller

import (
	"admin-service/global"
	"admin-service/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type idParam struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type listParams struct {
	Page, Limit, Title, SubTitle string
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
	var params listParams
	params.Page = c.DefaultQuery("page", "1")
	params.Limit = c.DefaultQuery("limit", "10")
	params.Title = c.DefaultQuery("title", "")
	params.SubTitle = c.DefaultQuery("subTitle", "")
	fmt.Println(params)
	list := []model.ArticleInfo{}
	global.Db.Where(&model.ArticleInfo{Title: params.Title, SubTitle: params.SubTitle}).Limit(params.Limit).Find(&list)
	fmt.Println(list)
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
