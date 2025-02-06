package main

import "fmt"

// Interface
type PaymentMethod interface {
	Pay(amount float64)
}

// Implementations
type CreditCard struct{}

func (c CreditCard) Pay(amount float64) { fmt.Println("Paid with Credit Card:", amount) }

type PayPal struct{}

func (p PayPal) Pay(amount float64) { fmt.Println("Paid with PayPal:", amount) }

// Factory function
func GetPaymentMethod(method string) PaymentMethod {
	switch method {
	case "credit":
		return CreditCard{}
	case "paypal":
		return PayPal{}
	default:
		return nil
	}
}

func main() {
	payment := GetPaymentMethod("credit")
	payment.Pay(100) // Output: Paid with Credit Card: 100
}
