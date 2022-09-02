package factory

import (
	"strings"
	"testing"
)

func TestCreateCashPaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type cash must exists")
	}

	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment message wasnt correct")
	}
	t.Log("LOG:", msg)
}

func TestCreateDebitCardPaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type Debit Card is expected")
	}

	msg := payment.Pay(20.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card response message wasnt correct")
	}
	t.Log("LOG:", msg)
}

func TestNonExistingPaymentMethods(t *testing.T) {
	_, err := GetPaymentMethod(30)
	if err == nil {
		t.Error("An invalid payment id must return an error")
	}

	t.Log("LOG:", err)
}
