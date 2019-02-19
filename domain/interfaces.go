package domain

type CustomerRW interface {
	Create(customer Customer)
	All() []Customer
	GetByFirstname(string) *Customer
}

type VtcRW interface {
	GetByFirstname(string) *Vtc
	Create(vtc Vtc)
	All() []Vtc
}

type AuthGateway interface {
	GetCurrent() *Customer
	Authenticate(Customer)
}

type AccountRW interface {
	Create(account Account)
	Get(string) *Account
	Set(Account)
}
