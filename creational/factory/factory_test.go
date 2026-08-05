package factory_test

import (
	"testing"

	"github.com/mohammedimrankasab/go-design-patterns/creational/factory"
)

func TestNotificationFactoryEmail(t *testing.T) {

	notification, err := factory.NewNotification(
		factory.Email,
		"user@example.com",
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if notification.Type() != "email" {
		t.Fatalf(
			"expected email notification, got %s",
			notification.Type(),
		)
	}
}

func TestNotificationFactorySMS(t *testing.T) {

	notification, err := factory.NewNotification(
		factory.SMS,
		"+911234567890",
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if notification.Type() != "sms" {
		t.Fatalf(
			"expected sms notification, got %s",
			notification.Type(),
		)
	}
}

func TestNotificationFactoryPush(t *testing.T) {

	notification, err := factory.NewNotification(
		factory.Push,
		"device123",
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if notification.Type() != "push" {
		t.Fatalf(
			"expected push notification, got %s",
			notification.Type(),
		)
	}
}
func TestNotificationFactoryInvalidType(t *testing.T) {

	_, err := factory.NewNotification(
		"whatsapp",
		"user",
	)

	if err == nil {
		t.Fatal(
			"expected error for unsupported notification type",
		)
	}
}
