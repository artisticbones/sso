package address

import "gorm.io/gorm"

type DAO interface {
	GetAddressById(id uint) (*Address, error)
	GetAddressesByUserId(userId uint) ([]*Address, error)
	CreateAddress(address *Address) *gorm.DB
	UpdateAddress(address *Address) *gorm.DB
	DeleteAddress(address *Address) *gorm.DB
}
