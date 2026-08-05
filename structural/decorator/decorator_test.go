package decorator

import (
	"testing"
)

type mockLogger struct {
	messages []string
}

func (m *mockLogger) Info(
	message string,
) {
	m.messages = append(
		m.messages,
		message,
	)
}

type mockMetrics struct {
	counters map[string]int
}

func (m *mockMetrics) Increment(
	metric string,
) {

	if m.counters == nil {
		m.counters = make(map[string]int)
	}

	m.counters[metric]++
}

func TestUserServiceReturnsUser(t *testing.T) {

	service := NewUserService()

	user, err := service.GetUser("1")

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if user.Name != "Imran" {
		t.Fatalf(
			"unexpected user %s",
			user.Name,
		)
	}
}

func TestUserServiceUserNotFound(t *testing.T) {

	service := NewUserService()

	_, err := service.GetUser("unknown")

	if err == nil {
		t.Fatal(
			"expected error",
		)
	}
}

func TestLoggingDecorator(t *testing.T) {

	logger := &mockLogger{}

	service :=
		NewLoggingDecorator(
			NewUserService(),
			logger,
		)

	_, err := service.GetUser("1")

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if len(logger.messages) != 1 {
		t.Fatal(
			"expected logging call",
		)
	}
}

func TestMetricsDecorator(t *testing.T) {

	metrics := &mockMetrics{
		counters: make(map[string]int),
	}

	service :=
		NewMetricsDecorator(
			NewUserService(),
			metrics,
		)

	_, err := service.GetUser("1")

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if metrics.counters["user_fetches"] != 1 {
		t.Fatal(
			"expected metric increment",
		)
	}
}

func TestDecoratorComposition(t *testing.T) {

	logger := &mockLogger{}

	metrics := &mockMetrics{
		counters: make(map[string]int),
	}

	service := NewUserService()

	service =
		NewLoggingDecorator(
			service,
			logger,
		)

	service =
		NewMetricsDecorator(
			service,
			metrics,
		)

	user, err := service.GetUser("1")

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if user.ID != "1" {
		t.Fatal(
			"expected user",
		)
	}

	if len(logger.messages) != 1 {
		t.Fatal(
			"logging decorator not executed",
		)
	}

	if metrics.counters["user_fetches"] != 1 {
		t.Fatal(
			"metrics decorator not executed",
		)
	}
}
