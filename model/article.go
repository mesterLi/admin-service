package model

type Tag struct {
	Id string
	Name string
}

type ArticleInfo struct {
	Title string `json:"title"`
	SubTitle string `json:"subTitle"`
	Id int `json:"id" gorm:"column:article_id"`
	Content string `json:"content"`
	Tags string `json:"tags"`
	ReadNum int `json:"readNum"`
	CommentNum int `json:"commentNum"`
	StarNum int `json:"starNum"`
	CreateTime string `json:"createTime"`
	UpdateTime string`json:"updateTime"`
	Status int `json:"status"`
}
