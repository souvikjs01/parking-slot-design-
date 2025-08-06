package payments

import "fmt"

type CardPayment struct {
	CardNumber string
}

func (c *CardPayment) Pay(amount float64) bool {
	fmt.Printf("Paid $%.2f via card ending in %s", amount, c.CardNumber[len(c.CardNumber)-4:])
	return true
}
