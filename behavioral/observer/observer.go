package observer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Event struct {
	ID        string
	Type      string
	Source    string
	Timestamp time.Time
	Payload   any
}

type Observer interface {
	Name() string
	Notify(ctx context.Context, event Event) error
}

type Publisher struct {
	mu        sync.RWMutex
	observers map[string]Observer
	order     []string
}

func New() *Publisher {
	return &Publisher{
		observers: make(map[string]Observer),
	}
}

func (p *Publisher) Subscribe(observer Observer) {
	p.mu.Lock()
	defer p.mu.Unlock()

	name := observer.Name()

	if _, exists := p.observers[name]; exists {
		return
	}

	p.observers[name] = observer
	p.order = append(p.order, name)
}

func (p *Publisher) Unsubscribe(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.observers[name]; !ok {
		return
	}

	delete(p.observers, name)

	for i, n := range p.order {
		if n == name {
			p.order = append(p.order[:i], p.order[i+1:]...)
			break
		}
	}
}

func (p *Publisher) Publish(ctx context.Context, event Event) error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var errs []error

	for _, name := range p.order {
		observer := p.observers[name]

		if err := observer.Notify(ctx, event); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
		}
	}

	return errors.Join(errs...)
}

// -----------------------------------------------------------------------------
// Email Observer Implementation
// -----------------------------------------------------------------------------
type EmailObserver struct{}

func (EmailObserver) Name() string {
	return "EmailObserver"
}

func (EmailObserver) Notify(ctx context.Context, event Event) error {
	fmt.Printf("Email observer got event %v\n", event)
	return nil
}

// -----------------------------------------------------------------------------
// Inventory Observer Implementation
// -----------------------------------------------------------------------------
type InventoryObserver struct{}

func (InventoryObserver) Name() string {
	return "InventoryObserver"
}

func (InventoryObserver) Notify(ctx context.Context, event Event) error {
	fmt.Printf("Inventory observer got event %v\n", event)
	return nil
}

// -----------------------------------------------------------------------------
// Analytics Observer Implementation
// -----------------------------------------------------------------------------
type AnalyticsObserver struct{}

func (AnalyticsObserver) Name() string {
	return "AnalyticsObserver"
}

func (AnalyticsObserver) Notify(ctx context.Context, event Event) error {
	fmt.Printf("Analytics observer got event %v\n", event)

	return nil
}

// -----------------------------------------------------------------------------
// Audit Observer Implementation
// -----------------------------------------------------------------------------
type AuditObserver struct{}

func (AuditObserver) Name() string {
	return "AuditObserver"
}

func (AuditObserver) Notify(ctx context.Context, event Event) error {
	fmt.Printf("Audit observer got event %v\n", event)

	return nil
}
