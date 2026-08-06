package chain

import (
	"errors"
	"log"
	"time"
)

var (
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrRateLimited     = errors.New("rate limit exceeded")
	ErrInvalidPayload  = errors.New("invalid payload")
)

type Context struct {
	RequestID string

	UserID string
	Token  string
	Role   string

	RateLimitRemaining int

	Payload  any
	Response any
}

type Handler interface {
	Handle(ctx *Context, chain Chain) error
}

type Chain interface {
	Next(ctx *Context) error
}

type middlewareChain struct {
	handlers []Handler
	index    int
}

func New(handlers ...Handler) Chain {
	return &middlewareChain{
		handlers: handlers,
	}
}

func (c *middlewareChain) Next(ctx *Context) error {
	if c.index >= len(c.handlers) {
		return nil
	}

	handler := c.handlers[c.index]
	c.index++

	return handler.Handle(ctx, c)
}

type LoggingMiddleware struct{}

func (l LoggingMiddleware) Handle(ctx *Context, chain Chain) error {
	start := time.Now()

	log.Printf("[%s] request started", ctx.RequestID)

	err := chain.Next(ctx)

	log.Printf("[%s] completed in %v", ctx.RequestID, time.Since(start))

	return err
}

type AuthenticationMiddleware struct{}

func (a AuthenticationMiddleware) Handle(ctx *Context, chain Chain) error {

	if ctx.Token == "" {
		return ErrUnauthenticated
	}

	ctx.UserID = "user-123"

	return chain.Next(ctx)
}

type AuthorizationMiddleware struct {
	RequiredRole string
}

func (a AuthorizationMiddleware) Handle(ctx *Context, chain Chain) error {

	if ctx.Role != a.RequiredRole {
		return ErrUnauthorized
	}

	return chain.Next(ctx)
}

type RateLimitMiddleware struct{}

func (r RateLimitMiddleware) Handle(ctx *Context, chain Chain) error {

	if ctx.RateLimitRemaining <= 0 {
		return ErrRateLimited
	}

	ctx.RateLimitRemaining--

	return chain.Next(ctx)
}

type ValidationMiddleware struct{}

func (v ValidationMiddleware) Handle(ctx *Context, chain Chain) error {

	if ctx.Payload == nil {
		return ErrInvalidPayload
	}

	return chain.Next(ctx)
}

type MetadataHandler struct{}

func (m MetadataHandler) Handle(ctx *Context, chain Chain) error {

	ctx.Response = "metadata processed"

	return chain.Next(ctx)
}
