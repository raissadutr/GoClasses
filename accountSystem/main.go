package main

import (
	c "accountSystem/accounts"
	"fmt"
)

func main() {

	brianAccount := c.NormalAccount{ClientName: "Brian", AgencyNumber: 123, AccountNumber: 12345, Balance: 155.89}
	hannahAccount := c.NormalAccount{ClientName: "Hannah", Balance: 1000}

	fmt.Println(brianAccount)
	fmt.Println(brianAccount.Deposit(200), brianAccount.Balance)
	fmt.Println(brianAccount.Draft(100), brianAccount.Balance)
	fmt.Println(hannahAccount.Transfer(500, &brianAccount), hannahAccount.Balance)

}
