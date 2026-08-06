package observer

import (
	"context"
	"errors"
	"testing"
	"time"
)

type testObserver struct {
	name      string
	called    bool
	callOrder *[]string
	err       error
}

func (t *testObserver) Name() string {
	return t.name
}

func (t *testObserver) Notify(ctx context.Context, event Event) error {
	t.called = true

	if t.callOrder != nil {
		*t.callOrder = append(*t.callOrder, t.name)
	}

	return t.err
}

func TestSubscribe(t *testing.T) {
	publisher := New()

	observer := &testObserver{name: "Email"}

	publisher.Subscribe(observer)

	if len(publisher.observers) != 1 {
		t.Fatalf("expected 1 observer, got %d", len(publisher.observers))
	}
}

func TestDuplicateSubscribe(t *testing.T) {
	publisher := New()

	observer := &testObserver{name: "Email"}

	publisher.Subscribe(observer)
	publisher.Subscribe(observer)

	if len(publisher.observers) != 1 {
		t.Fatal("duplicate observer should not be added")
	}
}

func TestUnsubscribe(t *testing.T) {
	publisher := New()

	observer := &testObserver{name: "Email"}

	publisher.Subscribe(observer)
	publisher.Unsubscribe(observer.Name())

	if len(publisher.observers) != 0 {
		t.Fatal("observer should be removed")
	}
}

func TestPublishSingleObserver(t *testing.T) {
	publisher := New()

	observer := &testObserver{name: "Email"}

	publisher.Subscribe(observer)

	event := Event{
		ID:        "evt-1",
		Type:      "OrderCreated",
		Source:    "OrderService",
		Timestamp: time.Now(),
	}

	if err := publisher.Publish(context.Background(), event); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !observer.called {
		t.Fatal("observer was not notified")
	}
}

func TestPublishMultipleObservers(t *testing.T) {
	publisher := New()

	email := &testObserver{name: "Email"}
	inventory := &testObserver{name: "Inventory"}
	analytics := &testObserver{name: "Analytics"}

	publisher.Subscribe(email)
	publisher.Subscribe(inventory)
	publisher.Subscribe(analytics)

	event := Event{
		ID:        "evt-1",
		Type:      "OrderCreated",
		Source:    "OrderService",
		Timestamp: time.Now(),
	}

	if err := publisher.Publish(context.Background(), event); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !email.called {
		t.Fatal("email observer not called")
	}

	if !inventory.called {
		t.Fatal("inventory observer not called")
	}

	if !analytics.called {
		t.Fatal("analytics observer not called")
	}
}

func TestPublishWithoutObservers(t *testing.T) {
	publisher := New()

	event := Event{
		ID:        "evt-1",
		Type:      "OrderCreated",
		Source:    "OrderService",
		Timestamp: time.Now(),
	}

	if err := publisher.Publish(context.Background(), event); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPublishReturnsJoinedErrors(t *testing.T) {
	publisher := New()

	expected := errors.New("email failed")

	email := &testObserver{
		name: "Email",
		err:  expected,
	}

	inventory := &testObserver{
		name: "Inventory",
	}

	publisher.Subscribe(email)
	publisher.Subscribe(inventory)

	event := Event{
		ID:        "evt-1",
		Type:      "OrderCreated",
		Source:    "OrderService",
		Timestamp: time.Now(),
	}

	err := publisher.Publish(context.Background(), event)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, expected) {
		t.Fatal("joined error should wrap original error")
	}

	if !inventory.called {
		t.Fatal("subsequent observers should still execute")
	}
}

func TestPublishMaintainsSubscriptionOrder(t *testing.T) {
	publisher := New()

	order := []string{}

	publisher.Subscribe(&testObserver{
		name:      "Email",
		callOrder: &order,
	})

	publisher.Subscribe(&testObserver{
		name:      "Inventory",
		callOrder: &order,
	})

	publisher.Subscribe(&testObserver{
		name:      "Analytics",
		callOrder: &order,
	})

	event := Event{
		ID:        "evt-1",
		Type:      "OrderCreated",
		Source:    "OrderService",
		Timestamp: time.Now(),
	}

	if err := publisher.Publish(context.Background(), event); err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"Email",
		"Inventory",
		"Analytics",
	}

	for i := range expected {
		if order[i] != expected[i] {
			t.Fatalf("expected %s got %s", expected[i], order[i])
		}
	}
}

func TestUnsubscribeUnknownObserver(t *testing.T) {
	publisher := New()

	publisher.Unsubscribe("unknown")

	if len(publisher.observers) != 0 {
		t.Fatal("publisher should remain unchanged")
	}
}
