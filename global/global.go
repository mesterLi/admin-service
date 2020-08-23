package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var Server *gin.Engine

var Db *gorm.DB

var Redis *redis.Client