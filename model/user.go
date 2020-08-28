package model

type User struct {
	Uid int `json:"uid" gorm:"column:uid;primary_key"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Status int `json:"status"`
}

type UserInfo struct {
	User
	CreateTime string `json:"createTime" gorm:"column:create_time"`
	UpdateTime string `json:"updateTime" gorm:"column:update_time"`
	Avator string `json:"avator" gorm:"column:avator"`
	Nickname string `json:"nickname" gorm:"column:nickname;default:null"`
}
