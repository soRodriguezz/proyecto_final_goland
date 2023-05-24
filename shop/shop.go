package shop

type Shop struct {
	Name     string
	Location string
}

func NewShop(name string, location string) *Shop {
	return &Shop{
		Name:     name,
		Location: location,
	}
}
