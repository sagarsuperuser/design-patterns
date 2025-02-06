package main

import "fmt"

type PaymentStratedy interface {
	Pay(amount float64)
}

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) {
	fmt.Println("Paid with Credit Card: ", amount)
}

type PayPal struct{}

func (p PayPal) Pay(amount float64) {
	fmt.Println("Paid with Pay Pal: ", amount)

}

type PaymentProcesssor struct {
	stratedy PaymentStratedy
}

func (p *PaymentProcesssor) SetStatedy(st PaymentStratedy) {
	p.stratedy = st
}

func (p *PaymentProcesssor) Pay(amount float64) {
	p.stratedy.Pay(amount)
}

func main() {
	processor := &PaymentProcesssor{stratedy: CreditCard{}}
	processor.Pay(200.00)
	processor.SetStatedy(PayPal{})
	processor.Pay(300.00)

}
