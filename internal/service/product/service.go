package product

type Service struct {
}

func NewServise() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProduct
}
