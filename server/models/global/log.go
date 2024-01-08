package global

type OptLog struct {
	Model
	UserId   uint   `gorm:"column:userId;type:uint;size:32" json:"userId"`
	Url      string `gorm:"column:url;type:varchar(255);not null;default:''" json:"url"`
	Method   string `gorm:"column:method;type:varchar(32);not null;default:'';index:method" json:"method"`
	Body     string `gorm:"column:body;type:varchar(255)" json:"body"`
	ClientIP string `gorm:"column:clientIP;type:varchar(255);not null;default:'';index:clientIP"`
}

func (OptLog) Table() string {
	return "sso_opt_log"
}

type AuthLog struct {
	Model
	Email    string `gorm:"column:email;type:varchar(255);index:email" json:"email"`
	ClientIP string `json:"clientIP" gorm:"column:clientIP;type:varchar(255);not null;default:''"`
	Device   string `json:"device" gorm:"column:device;type:varchar(255);not null;default:''"`
	Location string `json:"location" gorm:"column:location;type:varchar(255);not null;default:''"`
}

func (AuthLog) Table() string {
	return "sso_auth_log"
}
