package adapter

// PaymentProcessor defines the interface
// expected by the application.
type PaymentProcessor interface {
	Charge(amount float64) error
}

// LegacyPaymentGateway represents
// an existing third-party system.
type LegacyPaymentGateway struct {
	AccountID string
}

// MakePayment is the old API.
// It accepts amount in cents.
func (l *LegacyPaymentGateway) MakePayment(
	amountInCents int,
) error {

	// simulate legacy payment processing

	return nil
}

// PaymentAdapter converts the legacy API
// into the application's expected interface.
type PaymentAdapter struct {
	legacyGateway *LegacyPaymentGateway
}

// NewPaymentAdapter creates an adapter.
func NewPaymentAdapter(
	gateway *LegacyPaymentGateway,
) *PaymentAdapter {

	return &PaymentAdapter{
		legacyGateway: gateway,
	}
}

// Charge implements PaymentProcessor.
func (p *PaymentAdapter) Charge(
	amount float64,
) error {

	amountInCents := int(amount * 100)

	return p.legacyGateway.MakePayment(
		amountInCents,
	)
}
