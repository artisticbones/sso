package address

import "github.com/artisticbones/sso/server/models/global"

type Address struct {
	global.Model
	UserId   uint   `json:"userId" gorm:"column:userId;type:uint;size:32;not null"`
	Country  string `json:"country" gorm:"column:country;type:varchar(255);not null;index:address"`
	Province string `json:"province" gorm:"column:province;type:varchar(255);not null;index:address"`
	City     string `json:"city" gorm:"column:city;type:varchar(255);not null;index:address"`
	Distinct string `json:"distinct" gorm:"column:distinct;type:varchar(255);not null;index:address"`
	Detail   string `json:"detail" gorm:"column:detail;type:varchar(255);not null"`
	Tag      string `json:"tag" gorm:"column:tag;type:varchar(255);default:'';index:tag"`
}

func (Address) Table() string {
	return "sso_addresses"
}
