package users

import (
	context "context"
	"errors"
	"time"
)

type Usecase struct {
	Repo           DomainRepository
	contextTimeout time.Duration
}

func NewUsecase(repo DomainRepository, timeout time.Duration) *Usecase {
	return &Usecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (u *Usecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if string(domain.MSISDN[0:2]) != "62" {
		err := errors.New("error")
		return Domain{}, err
	}

	return u.Repo.Create(ctx, domain)
}

func (u *Usecase) Login(ctx context.Context, msisdn string, password string) (Domain, error) {
	if msisdn == "" || password == "" {
		err := errors.New("error")
		return Domain{}, err
	}

	return u.Repo.Login(ctx, msisdn, password)
}

func (u *Usecase) GetById(ctx context.Context, id string) (Domain, error) {

	return u.Repo.GetById(ctx, id)
}
