package accounts

type NormalAccount struct {
	ClientName    string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *NormalAccount) Deposit(amount float64) string {
	var message string
	isPossible := amount > 0
	if isPossible {
		c.Balance += amount
		message = "The operation was sucessful"
	} else {
		message = "Something went worng"
	}

	return message

}

func (c *NormalAccount) Draft(amount float64) string {
	var message string
	isPossible := amount > 0 && amount < c.Balance

	if isPossible {
		c.Balance -= amount
		message = "The operation was sucessful"
	} else {
		message = "Something went worng"
	}

	return message
}

func (c *NormalAccount) Transfer(amount float64, destAccount *NormalAccount) string {
	var message string
	isPossible := amount > 0 && amount < c.Balance

	if isPossible {
		c.Balance -= amount
		destAccount.Balance += amount
		message = "The operation was sucessful"
	} else {
		message = "Something went worng"
	}

	return message
}
