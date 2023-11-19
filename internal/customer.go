package internal

import (
	"errors"
)

const DEFAULT_DISCOUNT = 500

type Customer struct {
	Name     string
	Age      int
	balance  int
	debt     int
	discount bool
}

type Overdue struct {
	Customer
}

type Debtor interface {
	GetBalance() int
	GetDebt() int
	WrOffDebt() error
}

func (o *Overdue) WrOffDebt() error {
	if o.debt >= o.balance {
		return errors.New("Not possible write off")
	}

	o.balance -= o.debt
	o.debt = 0

	return nil
}

func (c *Customer) CalcDiscount() (int, error) {
	if !c.discount {
		return 0, errors.New("Discount not avaible")
	}
	result := DEFAULT_DISCOUNT - c.debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func (o *Overdue) GetBalance() int {
	return o.balance
}

func (o *Overdue) GetDebt() int {
	return o.debt
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Overdue {
	c := &Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		debt:     debt,
		discount: discount,
	}

	cc := &Overdue{
		*c,
	}
	return cc
}

type Discounter interface {
	CalcDiscount() (int, error)
}

func CalcPrice(c Discounter) (int, error) {
	debt, err := c.CalcDiscount()
	if err != nil {
		return 0, err
	}
	return debt, err
}
