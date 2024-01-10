package address

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DAO interface {
	GetAddressById(ctx context.Context, id uint) (*Address, error)
	GetAddressesByUserId(ctx context.Context, userId uint) ([]*Address, error)
	CreateAddress(ctx context.Context, address *Address) error
	UpdateAddress(ctx context.Context, address *Address) error
	DeleteAddress(ctx context.Context, id uint) error
}

type JdbcImpl struct {
	db *gorm.DB
}

func NewJdbcImpl(db *gorm.DB) *JdbcImpl {
	return &JdbcImpl{db: db}
}

func (impl *JdbcImpl) GetAddressById(ctx context.Context, id uint) (*Address, error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	addr := &Address{}
	err := impl.db.WithContext(ctx).Take(addr, "id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return addr, nil
}

func (impl *JdbcImpl) GetAddressesByUserId(ctx context.Context, userId uint) (addrs []*Address, err error) {
	if impl.db == nil {
		return nil, fmt.Errorf("DB ERROR, err = DB is Nil")
	}
	err = impl.db.WithContext(ctx).Find(&addrs, "userId = ?", userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return addrs, nil
}

func (impl *JdbcImpl) CreateAddress(ctx context.Context, address *Address) error {
	return impl.db.WithContext(ctx).Create(address).Error
}

func (impl *JdbcImpl) UpdateAddress(ctx context.Context, address *Address) error {
	return impl.db.WithContext(ctx).Updates(address).Error
}

func (impl *JdbcImpl) DeleteAddress(ctx context.Context, id uint) error {
	return impl.db.WithContext(ctx).Delete(&Address{}, id).Error
}
