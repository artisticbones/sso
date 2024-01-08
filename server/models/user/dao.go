package user

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DAO interface {
	GetUserById(id uint) (*User, error)
	GetAllUsers() ([]*User, error)
	CreateUser(user *User) *gorm.DB
	UpdateUser(user *User) *gorm.DB
	DeleteUser(id uint) *gorm.DB
}

type JdbcImpl struct {
	db *gorm.DB
}

func NewJdbcImpl(db *gorm.DB) *JdbcImpl {
	return &JdbcImpl{db: db}
}

func (impl *JdbcImpl) GetUserById(id uint) (*User, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	user := &User{}
	err := impl.db.Take(user, "id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, nil
}

func (impl *JdbcImpl) GetAllUsers() ([]*User, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	users := make([]*User, 0, 1)
	err := impl.db.Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return users, nil
}

func (impl *JdbcImpl) CreateUser(user *User) *gorm.DB {
	return impl.db.Create(user)
}

func (impl *JdbcImpl) UpdateUser(user *User) *gorm.DB {
	return impl.db.Updates(user)
}

func (impl *JdbcImpl) DeleteUser(id uint) *gorm.DB {
	return impl.db.Delete(&User{}, id)
}
