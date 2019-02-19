package domain

type Vtc struct {
	ID        string
	Firstname string
	Lastname  string
}

func NewVtc(id, firstName, lastName string) Vtc {
	return Vtc{id, firstName, lastName}
}
