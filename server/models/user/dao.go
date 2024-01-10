package user

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DAO interface {
	GetUserById(ctx context.Context, id uint) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
	CreateUser(ctx context.Context, user *User) *gorm.DB
	UpdateUser(ctx context.Context, user *User) *gorm.DB
	DeleteUser(ctx context.Context, id uint) *gorm.DB
}

type JdbcImpl struct {
	db *gorm.DB
}

func NewJdbcImpl(db *gorm.DB) *JdbcImpl {
	return &JdbcImpl{db: db}
}

func (impl *JdbcImpl) GetUserById(ctx context.Context, id uint) (*User, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	user := &User{}
	err := impl.db.WithContext(ctx).Take(user, "id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, nil
}

// GetAllUsers get all users from database.
// When the amount of data is large, this function may cause library dragging.
func (impl *JdbcImpl) GetAllUsers(ctx context.Context) (users []*User, err error) {
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

func (impl *JdbcImpl) CreateUser(ctx context.Context, user *User) *gorm.DB {
	return impl.db.WithContext(ctx).Create(user)
}

func (impl *JdbcImpl) UpdateUser(ctx context.Context, user *User) *gorm.DB {
	return impl.db.WithContext(ctx).Updates(user)
}

func (impl *JdbcImpl) DeleteUser(ctx context.Context, id uint) *gorm.DB {
	return impl.db.WithContext(ctx).Delete(&User{}, id)
}
