package proxy

import (
	"errors"
	"testing"
)

type mockAuthorization struct {
	allowed bool
	called  bool
}

func (m *mockAuthorization) CanAccess(
	userID,
	documentID string,
) bool {

	m.called = true
	return m.allowed
}

type mockDocumentService struct {
	called bool
	err    error
}

func (m *mockDocumentService) GetDocument(
	userID,
	documentID string,
) (*Document, error) {

	m.called = true

	if m.err != nil {
		return nil, m.err
	}

	return &Document{
		ID:    documentID,
		Title: "Test",
	}, nil
}

func TestAuthorizationProxy(t *testing.T) {

	tests := []struct {
		name          string
		allowed       bool
		serviceErr    error
		expectErr     error
		serviceCalled bool
	}{
		{
			name:          "authorized",
			allowed:       true,
			serviceCalled: true,
		},
		{
			name:      "unauthorized",
			allowed:   false,
			expectErr: ErrUnauthorized,
		},
		{
			name:          "document not found",
			allowed:       true,
			serviceErr:    ErrDocumentNotFound,
			expectErr:     ErrDocumentNotFound,
			serviceCalled: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			auth := &mockAuthorization{
				allowed: tt.allowed,
			}

			service := &mockDocumentService{
				err: tt.serviceErr,
			}

			proxy := NewAuthorizationProxy(
				auth,
				service,
			)

			_, err := proxy.GetDocument(
				"user-1",
				"doc-1",
			)

			if !auth.called {
				t.Fatal("authorization service should be called")
			}

			if service.called != tt.serviceCalled {
				t.Fatal("unexpected service invocation")
			}

			if tt.expectErr == nil {

				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				return
			}

			if !errors.Is(err, tt.expectErr) {
				t.Fatalf("expected %v got %v", tt.expectErr, err)
			}
		})
	}
}

func TestDocumentService(t *testing.T) {

	service := NewDocumentService()

	document, err := service.GetDocument(
		"user-1",
		"doc-1",
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if document.ID != "doc-1" {
		t.Fatal("unexpected document")
	}
}

func TestDocumentServiceNotFound(t *testing.T) {

	service := NewDocumentService()

	_, err := service.GetDocument(
		"user-1",
		"missing",
	)

	if !errors.Is(err, ErrDocumentNotFound) {
		t.Fatal("expected document not found")
	}
}
