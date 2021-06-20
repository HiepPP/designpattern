package factory

import (
	"errors"
	"fmt"
)

// Delegating the creation of new instance of structures to a different part of the program
// Working at the interface level instead of with concrete implementations
// Grouping families of objects to obtain a family object creator


const (
	Cash = 1
	Visa = 2
)

type PaymentMethod interface {
	Pay(amount float32) (bool, string)
}

func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case Cash:
		return new(CashPM), nil
	case Visa:
		return new(VisaPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", method))
	}
}

type CashPM struct {
}

func (c *CashPM) Pay(amount float32) (bool, string) {
	return true, fmt.Sprintf("%0.2f paid using cash\n", amount)
}

type VisaPM struct {
}

func (v *VisaPM) Pay(amount float32) (bool, string) {
	return true, fmt.Sprintf("%0.2f paid using visa\n", amount)
}
