package main

import (
	"fmt"

	"github.com/mohammedimrankasab/go-design-patterns/behavioral/chain"
)

func main() {

	chainOfResponsibility := chain.New(
		chain.LoggingMiddleware{},
		chain.AuthenticationMiddleware{},
		chain.AuthorizationMiddleware{
			RequiredRole: "admin",
		},
		chain.RateLimitMiddleware{},
		chain.ValidationMiddleware{},
		chain.MetadataHandler{},
	)

	ctx := &chain.Context{
		RequestID:          "req-001",
		Token:              "valid-token",
		Role:               "admin",
		RateLimitRemaining: 5,
		Payload: map[string]any{
			"type": "metadata",
		},
	}

	if err := chainOfResponsibility.Next(ctx); err != nil {
		panic(err)
	}

	fmt.Println(ctx.Response)
}
