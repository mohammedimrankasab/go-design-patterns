package main

import (
	"context"
	"log"
	"time"

	"github.com/mohammedimrankasab/go-design-patterns/behavioral/observer"
)

func main() {
	publisher := observer.New()

	emailObserver := observer.EmailObserver{}
	inventoryObserver := observer.InventoryObserver{}
	analyticsObserver := observer.AnalyticsObserver{}
	auditObserver := observer.AuditObserver{}

	publisher.Subscribe(emailObserver)
	publisher.Subscribe(inventoryObserver)
	publisher.Subscribe(analyticsObserver)
	publisher.Subscribe(auditObserver)

	event := observer.Event{
		ID:        "evt-001",
		Type:      "OrderCreated",
		Timestamp: time.Now(),
		Source:    "OrderService",
		Payload:   map[string]any{"ID": "ORD-1001"},
	}

	if err := publisher.Publish(context.Background(), event); err != nil {
		log.Fatal(err)
	}
}
