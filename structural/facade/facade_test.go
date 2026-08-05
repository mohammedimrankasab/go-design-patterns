package facade

import (
	"errors"
	"testing"
)

type mockInventory struct {
	called bool
	err    error
}

func (m *mockInventory) Reserve(order Order) error {
	m.called = true
	return m.err
}
func (m *mockInventory) Release(order Order) error {
	m.called = true
	return m.err
}

type mockPayment struct {
	called bool
	err    error
}

func (m *mockPayment) Process(order Order) error {
	m.called = true
	return m.err
}
func (m *mockPayment) Refund(order Order) error {
	m.called = true
	return m.err
}

type mockNotification struct {
	called bool
	err    error
}

func (m *mockNotification) Send(order Order) error {
	m.called = true
	return m.err
}

func TestPlaceOrderSuccess(t *testing.T) {

	inventory := &mockInventory{}
	payment := &mockPayment{}
	notification := &mockNotification{}

	facade := NewOrderFacade(
		inventory,
		payment,
		notification,
	)

	err := facade.PlaceOrder(Order{
		ID:     "order-1",
		Amount: 199.99,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !inventory.called {
		t.Fatal("inventory should be called")
	}

	if !payment.called {
		t.Fatal("payment should be called")
	}

	if !notification.called {
		t.Fatal("notification should be called")
	}
}

func TestInventoryFailure(t *testing.T) {

	inventory := &mockInventory{
		err: errors.New("inventory unavailable"),
	}

	payment := &mockPayment{}
	notification := &mockNotification{}

	facade := NewOrderFacade(
		inventory,
		payment,
		notification,
	)

	err := facade.PlaceOrder(Order{})

	if err == nil {
		t.Fatal("expected error")
	}

	if payment.called {
		t.Fatal("payment should not execute")
	}

	if notification.called {
		t.Fatal("notification should not execute")
	}
}

func TestPaymentFailure(t *testing.T) {

	inventory := &mockInventory{}

	payment := &mockPayment{
		err: errors.New("payment failed"),
	}

	notification := &mockNotification{}

	facade := NewOrderFacade(
		inventory,
		payment,
		notification,
	)

	err := facade.PlaceOrder(Order{})

	if err == nil {
		t.Fatal("expected error")
	}

	if !inventory.called {
		t.Fatal("inventory should execute")
	}

	if notification.called {
		t.Fatal("notification should not execute")
	}
}

func TestNotificationFailure(t *testing.T) {

	inventory := &mockInventory{}
	payment := &mockPayment{}

	notification := &mockNotification{
		err: errors.New("notification failed"),
	}

	facade := NewOrderFacade(
		inventory,
		payment,
		notification,
	)

	err := facade.PlaceOrder(Order{})

	if err == nil {
		t.Fatal("expected error")
	}

	if !inventory.called {
		t.Fatal("inventory should execute")
	}

	if !payment.called {
		t.Fatal("payment should execute")
	}
}
