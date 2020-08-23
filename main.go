package main

import (
	"admin-service/global"
	"admin-service/middleware"
	"admin-service/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"strings"
)

const (
	userName = "root"
	password = "123456-d"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "thankslyh"
)

func linkSql() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err := gorm.Open("mysql", path)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(10)
	db.DB().SetMaxOpenConns(10)
	global.Db = db
	//defer db.Close()
}

func linkRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "123456",
		DB: 9,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	global.Redis = client
}
func main() {
	r := gin.Default()
	go linkSql()
	//go linkRedis()
	r.Use(middleware.JWTAuth())
	global.Server = r
	router.Init()
	if err := r.Run(":8989"); err == nil {
		fmt.Println("server running at 8989")
	}
}