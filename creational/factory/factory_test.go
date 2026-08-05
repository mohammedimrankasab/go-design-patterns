package factory

import (
	"testing"
)

func TestNotificationFactoryEmail(t *testing.T) {

	notification, err := NewNotification(
		Email,
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

	notification, err := NewNotification(
		SMS,
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

	notification, err := NewNotification(
		Push,
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

	_, err := NewNotification(
		"whatsapp",
		"user",
	)

	if err == nil {
		t.Fatal(
			"expected error for unsupported notification type",
		)
	}
}
func TestNotificationSend(t *testing.T) {

	tests := []struct {
		name         string
		notification Notification
	}{
		{
			name: "email",
			notification: &EmailNotification{
				Recipient: "user@example.com",
			},
		},
		{
			name: "sms",
			notification: &SMSNotification{
				PhoneNumber: "+911234567890",
			},
		},
		{
			name: "push",
			notification: &PushNotification{
				DeviceID: "device-123",
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			err := tt.notification.Send(
				"hello",
			)

			if err != nil {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}
		})
	}
}
