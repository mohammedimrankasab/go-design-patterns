package functionaloptions

import (
	"testing"
	"time"
)

func TestNewServerDefaultConfiguration(t *testing.T) {

	server := NewServer()

	if server.Host != "localhost" {
		t.Errorf("Expected default host to be 'localhost', got '%s'", server.Host)
	}

	if server.Port != 8080 {
		t.Errorf("Expected default port to be 8080, got %d", server.Port)
	}

	if server.Timeout != 30*time.Second {
		t.Errorf("Expected default timeout to be 30 seconds, got %v", server.Timeout)
	}
	if server.Logging {
		t.Fatal("expected logging disabled")
	}

}

func TestNewServer_WithOptions(t *testing.T) {

	server := NewServer(
		WithHost("api.example.com"),
		WithPort(9000),
		WithTimeout(30*time.Second),
		WithLogging(true),
	)

	if server.Host != "api.example.com" {
		t.Fatalf("unexpected host %s", server.Host)
	}

	if server.Port != 9000 {
		t.Fatalf("unexpected port %d", server.Port)
	}

	if server.Timeout != 30*time.Second {
		t.Fatalf("unexpected timeout %v", server.Timeout)
	}

	if !server.Logging {
		t.Fatal("expected logging enabled")
	}
}
