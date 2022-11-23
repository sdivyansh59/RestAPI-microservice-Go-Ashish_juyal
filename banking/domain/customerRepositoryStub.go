package domain

// here wew will introduce adapter for 2nd port


type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "New Delhi", "110011", "2000-02-10", "1"},
		{"1002", "Prashant", "Bihar", "110022", "2008-08-18", "2"},
		{"1005", "Divyansh", "Uttarkhand", "250011", "1998-11-04", "1"},
	}
	return CustomerRepositoryStub{customers}
}