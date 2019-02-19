package vtcRW

import (
	"github.com/err0r500/vtc-bdd/domain"
)

type rw struct {
	repo []domain.Vtc
}

func (rw rw) GetByFirstname(firstname string) *domain.Vtc {
	for _, vtc := range rw.repo {
		if vtc.Firstname == firstname {
			return &vtc
		}
	}
	return nil
}

func (rw rw) All() []domain.Vtc {
	orig := rw.repo
	rw.repo = []domain.Vtc{}
	return orig
}

func (rw *rw) Create(vtc domain.Vtc) {
	rw.repo = append(rw.repo, vtc)
}

func New() domain.VtcRW {
	return &rw{
		repo: []domain.Vtc{},
	}
}
