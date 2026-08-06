package chain

import (
	"errors"
	"testing"
)

type testHandler struct {
	called bool
	err    error
}

func (t *testHandler) Handle(ctx *Context, chain Chain) error {
	t.called = true

	if t.err != nil {
		return t.err
	}

	return chain.Next(ctx)
}

func TestChainExecutesHandlersInOrder(t *testing.T) {
	first := &testHandler{}
	second := &testHandler{}
	third := &testHandler{}

	ctx := &Context{}

	chain := New(first, second, third)

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !first.called {
		t.Fatal("first handler was not called")
	}

	if !second.called {
		t.Fatal("second handler was not called")
	}

	if !third.called {
		t.Fatal("third handler was not called")
	}
}

func TestChainStopsOnError(t *testing.T) {
	expected := errors.New("boom")

	first := &testHandler{}
	second := &testHandler{
		err: expected,
	}
	third := &testHandler{}

	ctx := &Context{}

	chain := New(first, second, third)

	err := chain.Next(ctx)

	if !errors.Is(err, expected) {
		t.Fatalf("expected %v got %v", expected, err)
	}

	if !first.called {
		t.Fatal("first handler not executed")
	}

	if !second.called {
		t.Fatal("second handler not executed")
	}

	if third.called {
		t.Fatal("third handler should not execute")
	}
}

func TestAuthenticationMiddleware(t *testing.T) {
	ctx := &Context{
		Token: "valid-token",
	}

	handler := AuthenticationMiddleware{}

	chain := New(handler)

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ctx.UserID == "" {
		t.Fatal("user id should be populated")
	}
}

func TestAuthenticationMiddlewareFailsWithoutToken(t *testing.T) {
	ctx := &Context{}

	handler := AuthenticationMiddleware{}

	chain := New(handler)

	err := chain.Next(ctx)

	if !errors.Is(err, ErrUnauthenticated) {
		t.Fatalf("expected ErrUnauthenticated got %v", err)
	}
}

func TestAuthorizationMiddleware(t *testing.T) {
	ctx := &Context{
		Role: "admin",
	}

	handler := AuthorizationMiddleware{
		RequiredRole: "admin",
	}

	chain := New(handler)

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAuthorizationMiddlewareFails(t *testing.T) {
	ctx := &Context{
		Role: "viewer",
	}

	handler := AuthorizationMiddleware{
		RequiredRole: "admin",
	}

	chain := New(handler)

	err := chain.Next(ctx)

	if !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized got %v", err)
	}
}

func TestRateLimitMiddleware(t *testing.T) {
	ctx := &Context{
		RateLimitRemaining: 5,
	}

	handler := RateLimitMiddleware{}

	chain := New(handler)

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ctx.RateLimitRemaining != 4 {
		t.Fatalf("expected 4 got %d", ctx.RateLimitRemaining)
	}
}

func TestRateLimitMiddlewareFails(t *testing.T) {
	ctx := &Context{}

	handler := RateLimitMiddleware{}

	chain := New(handler)

	err := chain.Next(ctx)

	if !errors.Is(err, ErrRateLimited) {
		t.Fatalf("expected ErrRateLimited got %v", err)
	}
}

func TestValidationMiddleware(t *testing.T) {
	ctx := &Context{
		Payload: struct{}{},
	}

	handler := ValidationMiddleware{}

	chain := New(handler)

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidationMiddlewareFails(t *testing.T) {
	ctx := &Context{}

	handler := ValidationMiddleware{}

	chain := New(handler)

	err := chain.Next(ctx)

	if !errors.Is(err, ErrInvalidPayload) {
		t.Fatalf("expected ErrInvalidPayload got %v", err)
	}
}

func TestBusinessHandler(t *testing.T) {
	ctx := &Context{}

	handler := MetadataHandler{}

	chain := New(handler)

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ctx.Response == nil {
		t.Fatal("expected response")
	}
}
func TestChainSuccessfulRequest(t *testing.T) {
	ctx := &Context{
		RequestID:          "req-1",
		Token:              "valid-token",
		Role:               "admin",
		RateLimitRemaining: 5,
		Payload:            struct{}{},
	}

	chain := New(
		LoggingMiddleware{},
		AuthenticationMiddleware{},
		AuthorizationMiddleware{
			RequiredRole: "admin",
		},
		RateLimitMiddleware{},
		ValidationMiddleware{},
		MetadataHandler{},
	)

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ctx.UserID == "" {
		t.Fatal("expected user id to be populated")
	}

	if ctx.Response == nil {
		t.Fatal("expected response")
	}

	if ctx.RateLimitRemaining != 4 {
		t.Fatalf("expected rate limit to decrement to 4, got %d", ctx.RateLimitRemaining)
	}
}

func TestChainStopsAfterAuthenticationFailure(t *testing.T) {
	ctx := &Context{
		Role:               "admin",
		RateLimitRemaining: 5,
		Payload:            struct{}{},
	}

	chain := New(
		AuthenticationMiddleware{},
		AuthorizationMiddleware{
			RequiredRole: "admin",
		},
		RateLimitMiddleware{},
		ValidationMiddleware{},
		MetadataHandler{},
	)

	err := chain.Next(ctx)

	if !errors.Is(err, ErrUnauthenticated) {
		t.Fatalf("expected ErrUnauthenticated got %v", err)
	}

	if ctx.Response != nil {
		t.Fatal("business handler should not execute")
	}
}

func TestChainStopsAfterAuthorizationFailure(t *testing.T) {
	ctx := &Context{
		Token:              "valid-token",
		Role:               "viewer",
		RateLimitRemaining: 5,
		Payload:            struct{}{},
	}

	chain := New(
		AuthenticationMiddleware{},
		AuthorizationMiddleware{
			RequiredRole: "admin",
		},
		RateLimitMiddleware{},
		ValidationMiddleware{},
		MetadataHandler{},
	)

	err := chain.Next(ctx)

	if !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized got %v", err)
	}

	if ctx.Response != nil {
		t.Fatal("business handler should not execute")
	}
}

func TestChainStopsAfterRateLimitFailure(t *testing.T) {
	ctx := &Context{
		Token:              "valid-token",
		Role:               "admin",
		RateLimitRemaining: 0,
		Payload:            struct{}{},
	}

	chain := New(
		AuthenticationMiddleware{},
		AuthorizationMiddleware{
			RequiredRole: "admin",
		},
		RateLimitMiddleware{},
		ValidationMiddleware{},
		MetadataHandler{},
	)

	err := chain.Next(ctx)

	if !errors.Is(err, ErrRateLimited) {
		t.Fatalf("expected ErrRateLimited got %v", err)
	}

	if ctx.Response != nil {
		t.Fatal("business handler should not execute")
	}
}

func TestChainStopsAfterValidationFailure(t *testing.T) {
	ctx := &Context{
		Token:              "valid-token",
		Role:               "admin",
		RateLimitRemaining: 5,
	}

	chain := New(
		AuthenticationMiddleware{},
		AuthorizationMiddleware{
			RequiredRole: "admin",
		},
		RateLimitMiddleware{},
		ValidationMiddleware{},
		MetadataHandler{},
	)

	err := chain.Next(ctx)

	if !errors.Is(err, ErrInvalidPayload) {
		t.Fatalf("expected ErrInvalidPayload got %v", err)
	}

	if ctx.Response != nil {
		t.Fatal("business handler should not execute")
	}
}

func TestChainWithoutHandlers(t *testing.T) {
	ctx := &Context{}

	chain := New()

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestChainCallsEveryHandlerOnce(t *testing.T) {
	var count int

	handler := &countingHandler{
		count: &count,
	}

	chain := New(handler, handler, handler)

	if err := chain.Next(&Context{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if count != 3 {
		t.Fatalf("expected handler to execute 3 times, got %d", count)
	}
}

type countingHandler struct {
	count *int
}

func (c *countingHandler) Handle(ctx *Context, chain Chain) error {
	*c.count++
	return chain.Next(ctx)
}

func TestChainPropagatesContext(t *testing.T) {
	first := &contextHandler{}
	second := &verifyContextHandler{}

	ctx := &Context{}

	chain := New(first, second)

	if err := chain.Next(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

type contextHandler struct{}

func (c *contextHandler) Handle(ctx *Context, chain Chain) error {
	ctx.UserID = "user-123"
	return chain.Next(ctx)
}

type verifyContextHandler struct{}

func (v *verifyContextHandler) Handle(ctx *Context, chain Chain) error {
	if ctx.UserID != "user-123" {
		return errors.New("context not propagated")
	}
	return chain.Next(ctx)
}
