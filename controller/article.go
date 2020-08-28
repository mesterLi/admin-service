package controller

import (
	"admin-service/global"
	"admin-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	List: article_list,
	GetInfo: getInfo,
}
func article_list(c *gin.Context) {
	var total int
	list := []model.ArticleInfo{}
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	title := c.DefaultQuery("title", "")
	subTitle := c.DefaultQuery("subTitle", "")
	status := c.DefaultQuery("status", "0")
	statusInt, _ := strconv.Atoi(status)
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	var Db = global.Db.Table("articles")
	if statusInt != ARTICLE_STATUS_ALL {
		Db = Db.Where("status = ?", statusInt)
	}
	if title != "" {
		Db = Db.Where("title = ?", title)
	}
	if subTitle != "" {
		Db = Db.Where("subTitle = ?", title)
	}
	Db.Limit(limitInt).Offset((pageInt - 1) * limitInt).Find(&list).Count(&total)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": map[string] interface{}{
			"list": list,
			"total": total,
		},
		"message": "获取列表成功",
	})
}

func getInfo(c *gin.Context) {
	var info model.ArticleInfo
	id := c.Param("id")
	if id == "list" {
		article_list(c)
		return
	}
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "没有该文章",
			"data": model.ArticleInfo{},
		})
		return
	}
	Db := global.Db.Table("articles").Where("id = ?", id).First(&info)
	if info.Id == 0 || Db.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"message": "没有该文章",
			"data": model.ArticleInfo{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": info,
		"message": "success",
	})
}
