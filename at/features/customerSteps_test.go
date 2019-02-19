package feature_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	accountRW "github.com/err0r500/vtc-bdd/implem/inmem.accountRW"
	authGateway "github.com/err0r500/vtc-bdd/implem/inmem.authGateway"
	userRW "github.com/err0r500/vtc-bdd/implem/inmem.userRW"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/err0r500/vtc-bdd/domain"
	vtcRW "github.com/err0r500/vtc-bdd/implem/inmem.vtcRW"
	"github.com/err0r500/vtc-bdd/uc"
	"github.com/pkg/errors"
)

type runner struct {
	i           uc.Interactor
	currBooking *domain.Booking
}

func (r *runner) reset() {
	r.i = uc.Interactor{
		VtcRepo:      vtcRW.New(),
		CustomerRepo: userRW.New(),
		AuthGateway:  authGateway.New(),
		AccountRepo:  accountRW.New(),
	}
	r.currBooking = nil
}

func (r *runner) desClientsExistent(dataTable *gherkin.DataTable) error {
	for key, row := range dataTable.Rows {
		if key == 0 {
			continue
		}

		customer := domain.NewCustomer(row.Cells[0].Value, row.Cells[1].Value, row.Cells[2].Value)
		r.i.CustomerRepo.Create(customer)

		if customer != r.i.CustomerRepo.All()[key-1] {
			return errors.New("wooops")
		}
	}
	return nil
}

func (r *runner) desVTCExistent(dataTable *gherkin.DataTable) error {
	for key, value := range dataTable.Rows {
		if key == 0 {
			continue
		}

		vtc := domain.NewVtc(value.Cells[0].Value, value.Cells[1].Value, value.Cells[2].Value)
		r.i.VtcRepo.Create(vtc)
		if vtc != r.i.VtcRepo.All()[key-1] {
			return errors.New("wooops")
		}
	}
	return nil
}

func (r *runner) jeSuisAuthentifiEnTantQue(firstname string) error {
	currentUser := r.i.AuthGateway.GetCurrent()
	if currentUser == nil {
		user := r.i.CustomerRepo.GetByFirstname(firstname)
		if user == nil {
			return errors.New("user not found")
		}

		r.i.AuthGateway.Authenticate(*user)
		if *user != *r.i.AuthGateway.GetCurrent() {
			return errors.New("failed to insert user")
		}
		return nil
	}

	if currentUser.FirstName != firstname {
		return errors.New("other user currently logged in")
	}

	return nil
}

func (r *runner) leSoldeDeMonCompteEstDeEurosTTCAvecEurosTTCDavoir(rawSolde, rawAvoir string) error {
	solde, err := strconv.Atoi(rawSolde)
	if err != nil {
		return err
	}
	avoir, err := strconv.Atoi(rawAvoir)
	if err != nil {
		return err
	}

	currUser := r.i.AuthGateway.GetCurrent()
	if currUser == nil {
		return errors.New("user not logged in")
	}

	currAccount := r.i.AccountRepo.Get(currUser.ID)
	if currAccount == nil {
		r.i.AccountRepo.Create(domain.Account{ID: currUser.ID, Solde: solde, Avoir: avoir})
		return nil
	}
	if currAccount.Solde != solde || currAccount.Avoir != avoir {
		return errors.New("wrong amount")
	}

	return nil
}

func (r *runner) jeTenteDeRserverLeVTCDe(vtcFirstname, origin, destination string) error {
	r.currBooking = r.i.Bookvtc(vtcFirstname, origin, destination)
	return nil
}

func (r *runner) laRservationEstEffective() error {
	if r.currBooking == nil {
		return errors.New("reservation non effective")
	}
	return nil
}

func (r runner) laRservationNestPasEffective() error {
	if r.currBooking != nil {
		return errors.New("reservation effective")
	}
	return nil
}

func (r runner) etUneAlertePourInsuffisanceDeSoldeSeLve() error {
	return godog.ErrPending
}

func (r runner) jeNeSuisPasAuthentifi() error {
	return godog.ErrPending
}

func (r runner) etUneAlertePourIdentificationDuClientImpossibleSeLve() error {
	return godog.ErrPending
}

func SearchSteps(s *godog.Suite) {
	r := &runner{}

	s.BeforeScenario(func(interface{}) {
		r.reset()
	})

	s.Step(`^des clients existent:$`, r.desClientsExistent)
	s.Step(`^des VTC existent:$`, r.desVTCExistent)
	s.Step(`^je suis authentifié en tant que "([^"]*)"$`, r.jeSuisAuthentifiEnTantQue)
	s.Step(`^le solde de mon compte est de "([^"]*)" euros TTC avec "([^"]*)" euros TTC d\'avoir$`, r.leSoldeDeMonCompteEstDeEurosTTCAvecEurosTTCDavoir)
	s.Step(`^je tente de réserver le VTC "([^"]*)" de "([^"]*)" à "([^"]*)"$`, r.jeTenteDeRserverLeVTCDe)
	s.Step(`^la réservation est effective$`, r.laRservationEstEffective)
	s.Step(`^la réservation n\'est pas effective$`, r.laRservationNestPasEffective)
	s.Step(`^et une alerte pour insuffisance de solde se lève$`, r.etUneAlertePourInsuffisanceDeSoldeSeLve)
	s.Step(`^je ne suis pas authentifié$`, r.jeNeSuisPasAuthentifi)
	s.Step(`^et une alerte pour identification du client impossible se lève$`, r.etUneAlertePourIdentificationDuClientImpossibleSeLve)
}

func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		SearchSteps(s)
	}, godog.Options{
		Format:    "progress",
		Paths:     []string{"."},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
