package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

func main() {
	customer := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	fmt.Printf("%+v\n", customer)
}
func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
func startTransactionDynamic(debtor interface{}) error {
	answer, ok := debtor.(internal.Debtor)
	if !ok {
		return errors.New("Incorrect type")
	}
	return answer.WrOffDebt()
}
