package controller

import (
	"admin-service/global"
	"admin-service/model"
	"admin-service/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)
const FORMATE_TEM = "2006-01-02 15:04:05"
const (
	USER_STATUS_ON = iota + 1
	USER_STATUS_OFF
	USER_STATUS_DEL
)
type UserType struct {
	Login, Create, Info, List, Update, Delete gin.HandlerFunc
}

var User = UserType{
	Login: login,
	Create: create,
	Info: info,
	List: user_list,
	Update: user_update,
	Delete: user_delete,
}

func login(c *gin.Context) {
	var rBody *model.User
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("login err:{}", err)
		return
	}
	err = json.Unmarshal(body, &rBody)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	var curUser model.User
	//global.Db.First(&curUser)
	global.Db.Table("admin_users").Where("username = ?", rBody.Username).First(&curUser)
	fmt.Println("curUser", curUser)
	if curUser.Username == "" {
		fmt.Println("没有该账号～")
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"message": "没有该账号～",
		})
		return
	}
	if curUser.Password != rBody.Password {
		fmt.Println("密码错误～")
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"message": "密码错误～",
		})
		return
	}
	token, err := utils.GetToken(curUser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": token,
	})
	fmt.Println(curUser.Username)
	//err = global.Redis.Set(curUser.Username, token, time.Hour * 2).Err()
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func info(c *gin.Context) {
	var info model.UserInfo
	uid := c.Param("uid")
	fmt.Println("uid", uid == "list")
	if uid == "list" {
		user_list(c)
		return
	}
	if uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 40009,
			"data": nil,
			"message": "没有该用户信息！",
		})
		return
	}
	global.Db.Table("admin_users").Where("uid = ?", uid).First(&info)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": info,
	})
}

func create(c *gin.Context) {
	type AdminUser struct {
		Username string `json:"username" form:"username" binding:"required" gorm:"primary_key"`
		Password string `json:"password" form:"password" binding:"required"`
		CreateTime string `json:"create_time"`
	}
	var body AdminUser
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	body.CreateTime = time.Now().Format(FORMATE_TEM)
	err := global.Db.Table("admin_users").Where("username = ?", body.Username).Find(&AdminUser{})
	if err.Error == nil {
		fmt.Println("该用户已存在")
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"message": "该用户已存在",
		})
		return
	}
	global.Db.Table("admin_users").Create(body)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "创建用户成功",
	})
}

func user_list(c *gin.Context) {
	var users []model.UserInfo
	var total int
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	username := c.DefaultQuery("username", "")
	status := c.DefaultQuery("status", "")
	statusInt, _ := strconv.Atoi(status)
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	Db := global.Db.Table("admin_users")
	if username != "" {
		Db = Db.Where("username = ?", username)
	}
	fmt.Println("status=", status)
	if status != "" {
		Db = Db.Where("status = ?", statusInt)
	}
	//fmt.Println(pageInt)
	Db = Db.Limit(limit).Offset((pageInt - 1) * limitInt)
	Db.Find(&users).Count(&total)
	fmt.Println(users)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": map[string]interface{}{
			"list": users,
			"total": total,
		},
	})
}

func user_update(c *gin.Context) {
	type updateJson struct {
		Nickname, Avatar, CreateTime string
	}
	var putJson updateJson
	var updateInfo model.UserInfo
	c.ShouldBindJSON(&putJson)
	uid := c.Param("uid")
	if uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"data": nil,
			"message": "请传入要更新的用户ID",
		})
		return
	}
	Db := global.Db.Table("admin_users").Where("uid = ?", uid).First(&updateInfo)
	if updateInfo.Uid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"data": nil,
			"message": "没有该用户",
		})
		return
	}
	if putJson.Nickname != "" {
		updateInfo.Nickname = putJson.Nickname
	}
	if putJson.Avatar != "" {
		updateInfo.Avator = putJson.Avatar
	}
	updateInfo.UpdateTime = time.Now().Format(FORMATE_TEM)
	if err := Db.Update(&updateInfo); err.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"data": nil,
			"message": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": updateInfo,
		"message": "更新成功",
	})
}

func user_delete(c *gin.Context) {
	var info model.UserInfo
	uid := c.Param("uid")
	if uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"message": "请传入要删除的用户id",
		})
		return
	}
	Db := global.Db.Table("admin_users").Where("uid = ?", uid).First(&info)
	if info.Uid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"message": "没有该用户",
		})
		return
	}
	info.Status = USER_STATUS_DEL
	if err := Db.Update(&info); err.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "删除成功",
	})
}