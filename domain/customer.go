package domain

type Customer struct {
	ID        string
	FirstName string
	lastName  string
}

func NewCustomer(id, firstName, lastName string) Customer {
	return Customer{id, firstName, lastName}
}
