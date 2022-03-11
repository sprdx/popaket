package users

import (
	"context"
	"popaket/businesses/users"
	"popaket/helpers"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.DomainRepository {
	return &UserRepository{Conn: conn}
}

func (u *UserRepository) Create(ctx context.Context, user users.Domain) (users.Domain, error) {

	password, err := helpers.HashPassword(user.Password)
	if err != nil {
		return users.Domain{}, err
	}

	createdUser := Users{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Password: password,
		MSISDN:   user.MSISDN,
	}

	insertErr := u.Conn.Create(&createdUser).Error
	if insertErr != nil {
		return users.Domain{}, insertErr
	}

	return createdUser.ToDomain(), nil
}

func (u *UserRepository) Login(ctx context.Context, msisdn string, password string) (users.Domain, error) {
	var user Users
	err := u.Conn.Where("msisdn = ?", msisdn).First(&user).Error
	if err != nil {
		return users.Domain{}, err
	}

	if !helpers.IsMatched(user.Password, password) {
		return users.Domain{}, err
	}

	user.Token, _ = helpers.CreateToken(user.Id)

	tx := u.Conn.Save(&user)
	if tx.Error != nil {
		return users.Domain{}, tx.Error
	}

	return user.ToDomain(), nil
}
