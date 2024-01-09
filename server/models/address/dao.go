package address

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DAO interface {
	GetAddressById(id uint) (*Address, error)
	GetAddressesByUserId(userId uint) ([]*Address, error)
	CreateAddress(address *Address) error
	UpdateAddress(address *Address) error
	DeleteAddress(id uint) error
}

type JdbcImpl struct {
	db *gorm.DB
}

func (impl *JdbcImpl) GetAddressById(id uint) (*Address, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	addr := &Address{}
	err := impl.db.Take(addr, "id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return addr, nil
}

func (impl *JdbcImpl) GetAddressesByUserId(userId uint) ([]*Address, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	addrs := make([]*Address, 0, 1)
	err := impl.db.Find(&addrs, "userId = ?", userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return addrs, nil
}

func (impl *JdbcImpl) CreateAddress(address *Address) error {
	return impl.db.Create(address).Error
}

func (impl *JdbcImpl) UpdateAddress(address *Address) error {
	return impl.db.Updates(address).Error
}

func (impl *JdbcImpl) DeleteAddress(id uint) error {
	return impl.db.Delete(&Address{}, id).Error
}
