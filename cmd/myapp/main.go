package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	partner := internal.NewPartner("Dmitry", 23, 10000, 1000)

	startTransaction(partner)

	fmt.Printf("%+v\n", partner)
}
func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
