# Decorator Pattern

## Overview

The Decorator pattern allows additional behavior to be added to an object dynamically without modifying its original implementation.

In Go, this pattern is commonly implemented by wrapping interfaces.

---

## Problem

A service often needs additional capabilities:

- Logging
- Metrics
- Tracing
- Authentication
- Caching
- Retry handling

Adding all these responsibilities directly into the service creates:

- tightly coupled code
- difficult testing
- poor separation of concerns

Example:

```go
type UserService struct {
    // business logic
    // logging
    // metrics
    // tracing
}
```

---

## Solution

The Decorator pattern wraps the original implementation with additional behavior.

Each decorator implements the same interface as the original service.

Example:

```
MetricsDecorator
        |
LoggingDecorator
        |
UserService
```

The caller continues using the same interface.

---

## Implementation

Base interface:

```go
type UserService interface {
    GetUser(id string) (*User, error)
}
```

The concrete implementation:

```go
type userService struct {
    users map[string]*User
}
```

Decorators also implement:

```go
UserService
```

This allows multiple decorators to be composed together.

---

## Example Usage

```go
service := NewUserService()

service = NewLoggingDecorator(
    service,
    logger,
)

service = NewMetricsDecorator(
    service,
    metrics,
)

user, err := service.GetUser("1")
```

Execution flow:

```
Request
   |
   v
MetricsDecorator
   |
   v
LoggingDecorator
   |
   v
UserService
```

---

## Advantages

- Follows Open/Closed Principle
- Adds behavior without modifying existing code
- Improves testability
- Enables flexible composition
- Keeps business logic separated from cross-cutting concerns

---

## Trade-offs

Avoid using decorators when:

- additional behavior is very small
- the number of wrappers becomes difficult to understand
- debugging the execution chain becomes complicated

---

## Real World Usage

Common examples:

- HTTP middleware
- Logging wrappers
- Metrics instrumentation
- Distributed tracing
- Authentication middleware
- Retry mechanisms
- Cache layers

---

## Go Perspective

Decorator-style designs are common in Go because of interfaces and composition.

Examples:

- `http.Handler` middleware
- gRPC interceptors
- Database client wrappers
- Cloud SDK wrappers
- Service-layer instrumentation