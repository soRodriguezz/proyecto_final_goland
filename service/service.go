package service

type Service struct {
	Name  string
	Price float64
}

func NewService(name string, price float64) *Service {
	return &Service{
		Name:  name,
		Price: price,
	}
}
