package accountRW

import (
	"github.com/err0r500/vtc-bdd/domain"
)

type rw struct {
	accountRepo []domain.Account
}

func (r *rw) Set(inAccount domain.Account) {
	for k, account := range r.accountRepo {
		if account.ID == inAccount.ID {
			r.accountRepo[k] = inAccount
		}
	}
}

func (r rw) Get(id string) *domain.Account {
	for _, account := range r.accountRepo {
		if account.ID == id {
			return &account
		}
	}

	return nil
}

func (r *rw) Create(account domain.Account) {
	r.accountRepo = append(r.accountRepo, account)
}

func New() domain.AccountRW {
	return &rw{
		accountRepo: nil,
	}
}
