package main

import (
	"fmt"
)

type Account struct {
	clientName    string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {

	brianAccount := Account{clientName: "Brian", agencyNumber: 123, accountNumber: 12345, balance: 155.89}

	fmt.Println(brianAccount)

}
