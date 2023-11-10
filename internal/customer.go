package internal

import "errors"

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}

func CalcPrice(c Customer, price int) (int, error) {
	if c.Debt < price {
		return 0, errors.New("debt is less than price")
	}
	sum := c.Debt - price
	return sum, nil
}
