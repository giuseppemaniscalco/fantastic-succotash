package vending_machine

import "spark_networks_assessment/pkg/repositories/products"

type Service interface {
	Charge(coins []int) bool
	Select(product *products.Product) bool
	Balance() int
}

type service struct {
	userCoins   []int
	productRepo products.Repository
}

func New(productRepo products.Repository) Service {
	return &service{
		userCoins:   make([]int, 0),
		productRepo: productRepo,
	}
}

func (s *service) Balance() int {
	var response int
	for _, coins := range s.userCoins {
		response += coins
	}

	return response
}

func (s *service) Charge(coins []int) bool {
	//TODO validate coins

	s.userCoins = append(s.userCoins, coins...)

	return true
}

func (s *service) Select(product *products.Product) bool {
	if s.productRepo.Check(product) == 0 {
		return false
	}

	//Check if we have enough coins
	if s.Balance() >= product.Value {
		return s.productRepo.Purchase(product)
	}

	return false
}
