package model

import (
	"time"
)

type Tag struct {
	Id string
	Name string
}

type ArticleInfo struct {
	Title string `json:"title"`
	SubTitle string `json:"subTitle"`
	Id string `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Content string `json:"content"`
	Tags []Tag `json:"tags"`
	ReadNum int `json:"readNum"`
	CommentNum int `json:"commentNum"`
	StarNum int `json:"starNum"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time`json:"updateTime"`
}

func (ArticleInfo) TableName() string {
	return "articles"
}
