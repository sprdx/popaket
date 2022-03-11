package responses

import (
	"popaket/businesses/users"
)

type UserResponse struct {
	ID       string
	Name     string
	Username string
	MSISDN   string
}

type LoginResponse struct {
	Token string
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:       domain.Id,
		Name:     domain.Name,
		Username: domain.Username,
		MSISDN:   domain.MSISDN,
	}
}
