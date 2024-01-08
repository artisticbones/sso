package profile

import "github.com/artisticbones/sso/server/models/global"

type Profile struct {
	global.Model
	UserId uint   `json:"userId" gorm:"column:userId;type:uint;size:32;not null"`
	Name   string `json:"name" gorm:"column:name;type:varchar(255)"`
	IDCard string `json:"IDCard" gorm:"column:IDCard;type:varchar(255)"`
	Avatar string `json:"avatar" gorm:"column:avatar;type:varchar(255)"`
	Gender string `json:"gender" gorm:"column:gender;type:enum('男','女','保密');not null;default:'保密'"`
	Intro  string `json:"intro" gorm:"column:intro;type:varchar(512)"`
}

func (Profile) Table() string {
	return "sso_profiles"
}
