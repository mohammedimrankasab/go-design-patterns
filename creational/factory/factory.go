package factory

import "fmt"

type Notification interface {
	Send(message string) error
	Type() string
}

type EmailNotification struct {
	Recipient string
}

func (e *EmailNotification) Send(message string) error {
	// Logic to send email
	fmt.Printf("Sending Email to %s: %s\n", e.Recipient, message)
	return nil
}

func (e *EmailNotification) Type() string {
	return "email"
}

type SMSNotification struct {
	PhoneNumber string
}

func (s *SMSNotification) Send(message string) error {
	// Logic to send SMS
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
	return nil
}

func (s *SMSNotification) Type() string {
	return "sms"
}

type PushNotification struct {
	DeviceID string
}

func (p *PushNotification) Send(message string) error {
	// Logic to send push notification
	fmt.Printf("Sending Push Notification to %s: %s\n", p.DeviceID, message)
	return nil
}

func (p *PushNotification) Type() string {
	return "push"
}

type NotificationType string

const (
	Email NotificationType = "email"
	SMS   NotificationType = "sms"
	Push  NotificationType = "push"
)

func NewNotification(
	notificationType NotificationType,
	recipient string,
) (Notification, error) {

	switch notificationType {
	case Email:
		return &EmailNotification{Recipient: recipient}, nil
	case SMS:
		return &SMSNotification{PhoneNumber: recipient}, nil
	case Push:
		return &PushNotification{DeviceID: recipient}, nil
	default:
		return nil, fmt.Errorf("unsupported notification type: %s", notificationType)
	}
}
