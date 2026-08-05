package proxy

import (
	"errors"
	"fmt"
)

var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrDocumentNotFound = errors.New("document not found")
)

// Document represents a protected resource.
type Document struct {
	ID      string
	Title   string
	Content string
}

// DocumentService defines document operations.
type DocumentService interface {
	GetDocument(userID, documentID string) (*Document, error)
}

// AuthorizationService validates permissions.
type AuthorizationService interface {
	CanAccess(userID, documentID string) bool
}

// documentService is the real implementation.
type documentService struct {
	documents map[string]*Document
}

// NewDocumentService creates the real service.
func NewDocumentService() DocumentService {

	return &documentService{
		documents: map[string]*Document{
			"doc-1": {
				ID:      "doc-1",
				Title:   "System Design",
				Content: "Scaling distributed systems.",
			},
		},
	}
}

func (d *documentService) GetDocument(
	userID,
	documentID string,
) (*Document, error) {

	document, ok := d.documents[documentID]
	if !ok {
		return nil, ErrDocumentNotFound
	}

	return document, nil
}

// AuthorizationProxy controls access to the service.
type AuthorizationProxy struct {
	auth    AuthorizationService
	service DocumentService
}

// NewAuthorizationProxy creates a proxy.
func NewAuthorizationProxy(
	auth AuthorizationService,
	service DocumentService,
) DocumentService {

	return &AuthorizationProxy{
		auth:    auth,
		service: service,
	}
}

func (p *AuthorizationProxy) GetDocument(
	userID,
	documentID string,
) (*Document, error) {

	if !p.auth.CanAccess(userID, documentID) {
		return nil, fmt.Errorf(
			"%w: user %q cannot access %q",
			ErrUnauthorized,
			userID,
			documentID,
		)
	}

	document, err := p.service.GetDocument(
		userID,
		documentID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"get document: %w",
			err,
		)
	}

	return document, nil
}
