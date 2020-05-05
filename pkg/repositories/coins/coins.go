package coins

import "errors"

const (
	invalidAmountError  = "invalid amount"
	negativeAmountError = "negative amount"
	zeroAmountError     = "zero amount"
)

var (
	validCoinAmount = map[int]bool{
		1:  true,
		5:  true,
		10: true,
		25: true,
	}
)

type Repository interface {
	List() []int
	Add(amount int) error
}

type repository struct {
	amount []int
}

func New() Repository {
	return &repository{
		amount: make([]int, 0),
	}
}

func (r *repository) List() []int {
	return r.amount
}

func (r *repository) Add(amount int) error {
	if amount < 0 {
		return errors.New(negativeAmountError)
	}
	if amount == 0 {
		return errors.New(zeroAmountError)
	}
	if _, ok := validCoinAmount[amount]; !ok {
		return errors.New(invalidAmountError)
	}

	r.amount = append(r.amount, amount)

	return nil
}
