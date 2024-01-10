package profile

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DAO interface {
	GetProfileByUserId(ctx context.Context, userId uint) (*Profile, error)
	GetAllProfiles(ctx context.Context) ([]*Profile, error)
	CreateProfile(ctx context.Context, user *Profile) *gorm.DB
	UpdateProfile(ctx context.Context, user *Profile) *gorm.DB
	DeleteProfile(ctx context.Context, id uint) *gorm.DB
}

type JdbcImpl struct {
	db *gorm.DB
}

func (impl *JdbcImpl) GetProfileByUserId(ctx context.Context, userId uint) (*Profile, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	user := &Profile{}
	err := impl.db.WithContext(ctx).Take(user, "userId = ?", userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, nil
}

// GetAllProfiles get all users' profile from database.
// When the amount of data is large, this function may cause library dragging.
func (impl *JdbcImpl) GetAllProfiles(ctx context.Context) (users []*Profile, err error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	err = impl.db.WithContext(ctx).Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return users, nil
}

func (impl *JdbcImpl) CreateProfile(ctx context.Context, user *Profile) *gorm.DB {
	return impl.db.WithContext(ctx).Create(user)
}

func (impl *JdbcImpl) UpdateProfile(ctx context.Context, user *Profile) *gorm.DB {
	return impl.db.WithContext(ctx).Updates(user)
}

func (impl *JdbcImpl) DeleteProfile(ctx context.Context, userId uint) *gorm.DB {
	return impl.db.WithContext(ctx).Delete(&Profile{}, "userId = ?", userId)
}
