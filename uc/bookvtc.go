package uc

import (
	"github.com/err0r500/vtc-bdd/domain"
)

type Interactor struct {
	AccountRepo  domain.AccountRW
	VtcRepo      domain.VtcRW
	CustomerRepo domain.CustomerRW
	AuthGateway  domain.AuthGateway
}

func (i Interactor) Bookvtc(vtcDriver, origin, destination string) *domain.Booking {

	i.AccountRepo.Set(
		i.AccountRepo.Get(i.AuthGateway.GetCurrent().ID).Charge(origin, destination),
	)

	return &domain.Booking{
		Customer: *i.AuthGateway.GetCurrent(),
		Vtc:      *i.VtcRepo.GetByFirstname(vtcDriver),
	}
}
