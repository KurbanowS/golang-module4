package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	c := 1000
	fmt.Println(internal.CalcPrice(*cust, c))
	cust.CalcDiscount()
	fmt.Printf("%+v\n", cust)
	cust.WrOffDebt()
}
