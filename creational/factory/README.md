# Factory Pattern

## Overview

The Factory pattern provides a way to create objects without exposing the creation logic to the caller.

The caller depends on an abstraction instead of concrete implementations.

---

## Problem

Without a factory:

```go
notification := EmailNotification{}
```

The application becomes tightly coupled to a specific implementation.

Adding new notification providers requires changes throughout the codebase.

---

## Solution

The factory centralizes object creation.

Example:

```go
notification, err := factory.NewNotification(
    factory.Email,
    "user@example.com",
)
```

The caller only knows about the interface:

```go
type Notification interface {
    Send(message string) error
}
```

---

## Implementation

The factory:

1. Receives the requested type
2. Creates the correct implementation
3. Returns the interface

Flow:

```
Request
   |
   v
Factory
   |
   +---- EmailNotification
   |
   +---- SMSNotification
   |
   +---- PushNotification
```

---

## Advantages

- Reduces coupling
- Centralizes creation logic
- Makes adding implementations easier
- Improves testability

---

## Trade-offs

Avoid factories when:

- creation logic is trivial
- there are very few implementations
- abstraction provides no value

---

## Real World Usage

Common examples:

- Notification providers
- Database drivers
- Cloud storage clients
- Payment providers
- Message queue clients