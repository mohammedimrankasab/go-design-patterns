package builder

import (
	"testing"
	"time"
)

func TestHTTPClientBuilder_DefaultValues(t *testing.T) {

	client := NewClientBuilder().
		Build()

	if client.Timeout != 30*time.Second {
		t.Fatalf(
			"expected default timeout 30s, got %v",
			client.Timeout,
		)
	}

	if client.Retries != 3 {
		t.Fatalf(
			"expected default retries 3, got %d",
			client.Retries,
		)
	}
}

func TestHTTPClientBuilder_CustomConfiguration(t *testing.T) {

	client := NewClientBuilder().
		WithBaseURL("https://api.example.com").
		WithTimeout(60*time.Second).
		WithRetries(5).
		WithHeader("Authorization", "Bearer token").
		EnableTracing().
		Build()

	if client.BaseURL != "https://api.example.com" {
		t.Fatalf(
			"unexpected base URL %s",
			client.BaseURL,
		)
	}

	if client.Timeout != 60*time.Second {
		t.Fatalf(
			"unexpected timeout %v",
			client.Timeout,
		)
	}

	if client.Retries != 5 {
		t.Fatalf(
			"unexpected retries %d",
			client.Retries,
		)
	}

	if client.Headers["Authorization"] != "Bearer token" {
		t.Fatalf(
			"authorization header missing",
		)
	}

	if !client.EnableTrace {
		t.Fatal("expected tracing enabled")
	}
}
