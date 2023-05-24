package customer

type Customer struct {
	Name  string
	Phone string
}

func NewCustomer(name string, phone string) *Customer {
	return &Customer{
		Name:  name,
		Phone: phone,
	}
}
