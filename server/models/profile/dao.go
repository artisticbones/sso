package profile

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DAO interface {
	GetProfileByUserId(userId uint) (*Profile, error)
	GetAllProfiles() ([]*Profile, error)
	CreateProfile(user *Profile) *gorm.DB
	UpdateProfile(user *Profile) *gorm.DB
	DeleteProfile(id uint) *gorm.DB
}

type JdbcImpl struct {
	db *gorm.DB
}

func (impl *JdbcImpl) GetProfileByUserId(userId uint) (*Profile, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	user := &Profile{}
	err := impl.db.Take(user, "userId = ?", userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, nil
}

func (impl *JdbcImpl) GetAllProfiles() ([]*Profile, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	users := make([]*Profile, 0, 1)
	err := impl.db.Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return users, nil
}

func (impl *JdbcImpl) CreateProfile(user *Profile) *gorm.DB {
	return impl.db.Create(user)
}

func (impl *JdbcImpl) UpdateProfile(user *Profile) *gorm.DB {
	return impl.db.Updates(user)
}

func (impl *JdbcImpl) DeleteProfile(userId uint) *gorm.DB {
	return impl.db.Delete(&Profile{}, "userId = ?", userId)
}
