package factory

import (
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type cash must exist")
	}

	success, msg := payment.Pay(10.30)
	if !success {
		t.Error("Fail payment")
	}
	t.Log("LOG:", msg)
}
