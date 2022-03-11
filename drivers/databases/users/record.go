package users

import (
	"popaket/businesses/users"
)

type Users struct {
	Id       string `gorm:"primaryKey"`
	Name     string
	Username string `gorm:"unique"`
	Password string
	MSISDN   string `gorm:"unique"`
	Token    string
}

func (Users) TableName() string {
	return "users"
}

func (u *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		MSISDN:   u.MSISDN,
		Token:    u.Token,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:       domain.Id,
		Name:     domain.Name,
		Username: domain.Username,
		MSISDN:   domain.MSISDN,
	}
}
