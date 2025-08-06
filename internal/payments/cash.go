package payments

import "fmt"

type CashPayment struct{}

func (c *CashPayment) Pay(amount float64) bool {
	fmt.Printf("Paid $%.2f via cash", amount)
	return true
}
