package users

import (
	"context"
)

type Domain struct {
	Id       string
	Name     string
	Username string
	Password string
	MSISDN   string
	Token    string
}

// Data Access Layer ke Domain
type DomainRepository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, msisdn string, password string) (Domain, error)
	GetById(ctx context.Context, id string) (Domain, error)
}

type DomainService interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, msisdn string, password string) (Domain, error)
	GetById(ctx context.Context, token string) (Domain, error)
}
