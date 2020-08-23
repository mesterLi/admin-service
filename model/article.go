package model

import "time"

type Tag struct {
	Id string
	Name string
}

type ArticleInfo struct {
	Title string `json:"title" db:"title"`
	SubTitle string `json:"subTitle" db:"subTitle"`
	Id string `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
	Tags []Tag `json:"tags" db:"tags"`
	ReadNum int `json:"readNum" db:"readNum"`
	CommentNum int `json:"commentNum" db:"commentNum"`
	StarNum int `json:"starNum" db:"starNum"`
	CreateTime time.Time `json:"createTime" db:"createTime"`
	UpdateTime time.Time`json:"updateTime" db:"updateTime"`
}
