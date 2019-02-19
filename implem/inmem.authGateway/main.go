package authGateway

import (
	"github.com/err0r500/vtc-bdd/domain"
)

type gateway struct {
	loggedUser *domain.Customer
}

func (g gateway) GetCurrent() *domain.Customer {
	return g.loggedUser
}

func (g *gateway) Authenticate(customer domain.Customer) {
	g.loggedUser = &customer
}

func New() domain.AuthGateway {
	return &gateway{loggedUser: nil}
}
