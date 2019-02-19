package userRW

import (
	"github.com/err0r500/vtc-bdd/domain"
)

type rw struct {
	repo []domain.Customer
}

func (rw rw) GetByFirstname(firstname string) *domain.Customer {
	for _, customer := range rw.repo {
		if customer.FirstName == firstname {
			return &customer
		}
	}
	return nil
}

func (rw *rw) Create(customer domain.Customer) {
	rw.repo = append(rw.repo, customer)
}

func (rw rw) All() []domain.Customer {
	return rw.repo
}

func New() domain.CustomerRW {
	return &rw{
		repo: []domain.Customer{},
	}
}
