package facade

import "fmt"

// Order represents a customer order.
type Order struct {
	ID     string
	Amount float64
}

// InventoryService reserves inventory.
type InventoryService interface {
	Reserve(order Order) error
	Release(order Order) error
}

// PaymentService processes payments.
type PaymentService interface {
	Process(order Order) error
	Refund(order Order) error
}

// NotificationService sends confirmations.
type NotificationService interface {
	Send(order Order) error
}

// OrderFacade orchestrates order placement.
type OrderFacade struct {
	inventoryService    InventoryService
	paymentService      PaymentService
	notificationService NotificationService
}

// NewOrderFacade creates a new facade.
func NewOrderFacade(
	inventory InventoryService,
	payment PaymentService,
	notification NotificationService,
) *OrderFacade {

	return &OrderFacade{
		inventoryService:    inventory,
		paymentService:      payment,
		notificationService: notification,
	}
}

// PlaceOrder executes the order workflow.
func (o *OrderFacade) PlaceOrder(
	order Order,
) error {

	if err := o.inventoryService.Reserve(order); err != nil {
		return fmt.Errorf(
			"reserve inventory: %w",
			err,
		)
	}

	if err := o.paymentService.Process(order); err != nil {
		return fmt.Errorf(
			"process payment: %w",
			err,
		)
	}

	if err := o.notificationService.Send(order); err != nil {
		return fmt.Errorf(
			"send notification: %w",
			err,
		)
	}

	return nil
}
