package domain

import (
	"strings"
)

type Account struct {
	ID    string
	Solde int
	Avoir int
}

func (account Account) Charge(origin, destination string) Account {
	if !strings.Contains(origin, "Paris") {
		return account
	}

	fromSolde := 40
	fromAvoir := 10
	if strings.Contains(destination, "Paris") {
		fromSolde = 30
		fromAvoir = 0
	}

	if account.Avoir > fromSolde+fromAvoir {
		fromAvoir = fromSolde + fromAvoir
		fromSolde = 0
	}

	return Account{
		ID:    account.ID,
		Solde: account.Solde - fromSolde,
		Avoir: account.Avoir - fromAvoir}
}
