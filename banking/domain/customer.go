package domain

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// 2nd port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
