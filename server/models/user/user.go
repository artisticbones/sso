package user

import "github.com/artisticbones/sso/server/models/global"

type User struct {
	global.Model
	Email    string `json:"email" gorm:"column:email;type:varchar(255);unique;not null"`
	Username string `json:"username" gorm:"column:username;type:varchar(255);not null"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(255);default:''"`
}

func (User) Table() string {
	return "sso_users"
}
