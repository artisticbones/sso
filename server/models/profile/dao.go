package profile

import "gorm.io/gorm"

type DAO interface {
	GetProfileByUserId(userId uint) (*Profile, error)
	GetAllProfiles() ([]*Profile, error)
	CreateProfile(user *Profile) *gorm.DB
	UpdateProfile(user *Profile) *gorm.DB
	DeleteProfile(id uint) *gorm.DB
}
