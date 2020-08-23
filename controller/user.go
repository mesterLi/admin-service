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
	"time"
)

func Login(c *gin.Context) {
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
	global.Db.Where("username = ?", rBody.Username).First(&curUser)
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

func Info(c *gin.Context) {
	var info model.UserInfo
	uid, exit := c.Get("uid")
	if exit == false {
		c.JSON(http.StatusOK, gin.H{
			"status": 40009,
			"data": "没有该用户信息！",
			"message": "没有该用户信息！",
		})
		return
	}
	global.Db.Where("uid = ?", uid).First(&info)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": info,
	})
}

func Create(c *gin.Context) {
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
	body.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	err := global.Db.Where("username = ?", body.Username).Find(&AdminUser{})
	if err.Error == nil {
		fmt.Println("该用户已存在")
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"message": "该用户已存在",
		})
		return
	}
	global.Db.Create(body)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "创建用户成功",
	})
}