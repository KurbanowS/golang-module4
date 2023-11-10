package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	c := 1000
	fmt.Println(internal.CalcPrice(*cust, c))

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not avaible")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	fmt.Printf("%+v\n", cust)

}
