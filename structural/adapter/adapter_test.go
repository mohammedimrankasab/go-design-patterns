package adapter

import (
	"testing"
)

func TestPaymentAdapter(t *testing.T) {

	legacyGateway := &LegacyPaymentGateway{
		AccountID: "legacy-account",
	}

	paymentProcessor :=
		NewPaymentAdapter(
			legacyGateway,
		)

	err := paymentProcessor.Charge(
		100.50,
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}
}

func TestPaymentAdapterImplementsInterface(
	t *testing.T,
) {

	var processor PaymentProcessor = NewPaymentAdapter(
		&LegacyPaymentGateway{},
	)

	if _, ok := processor.(*PaymentAdapter); !ok {
		t.Fatal(
			"expected payment processor to implement PaymentProcessor",
		)
	}
}
