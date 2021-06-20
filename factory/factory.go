package factory

import (
	"errors"
)

const (
	Cash = 1
	Visa = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error){
	return nil, errors.New("Not implement yet")
}

type CashPM struct {

}

func (c *CashPM) Pay(amount float32) string{
	return "cash"
}

type VisaPM struct {

}

func (v *VisaPM) Pay(amount float32) string{
	return "visa"
}