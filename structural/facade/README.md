# Facade Pattern

## Overview

The **Facade Pattern** provides a simplified interface to a complex subsystem by hiding the interactions between multiple components behind a single entry point.

Instead of requiring clients to coordinate multiple services, a facade encapsulates the workflow and exposes a single high-level operation.

In this example, an `OrderFacade` orchestrates the complete order placement process by coordinating inventory, payment, and notification services.

---

# Problem

Placing an order involves multiple independent services.

- Reserve inventory
- Process payment
- Send an order confirmation

Without a facade, every client must know:

- Which services to call
- The correct execution order
- How to handle failures
- How to coordinate the workflow

```text
Client
 ├── Inventory Service
 ├── Payment Service
 └── Notification Service
```

As the workflow grows, orchestration logic becomes duplicated across the application, making it harder to maintain and test.

---

# Solution

Introduce an `OrderFacade` that hides the orchestration behind a single method.

```text
                Client
                   │
                   ▼
             OrderFacade
          ┌──────┼──────┐
          │      │      │
          ▼      ▼      ▼
   Inventory  Payment  Notification
```

The client simply invokes:

```go
facade := NewOrderFacade(
    inventory,
    payment,
    notification,
)

err := facade.PlaceOrder(order)
```

The facade is responsible for coordinating the workflow while the client remains unaware of the underlying implementation.

---

# Implementation

The facade coordinates three independent services.

## Inventory Service

```go
type InventoryService interface {
    Reserve(order Order) error
    Release(order Order) error
}
```

## Payment Service

```go
type PaymentService interface {
    Process(order Order) error
    Refund(order Order) error
}
```

## Notification Service

```go
type NotificationService interface {
    Send(order Order) error
}
```

The facade exposes a single high-level operation.

```go
func (o *OrderFacade) PlaceOrder(order Order) error
```

---

# Workflow

Successful execution follows a simple orchestration flow.

```text
Reserve Inventory
        │
        ▼
Process Payment
        │
        ▼
Send Notification
        │
        ▼
Success
```

If any operation fails, execution stops immediately and the error is returned to the caller.

---

# Error Handling

The facade wraps errors to provide additional context.

```go
return fmt.Errorf("process payment: %w", err)
```

This allows callers to understand which subsystem failed while still preserving the original error.

---

# Why Use a Facade?

Without a facade:

```go
inventory.Reserve(order)

payment.Process(order)

notification.Send(order)
```

Every caller needs to understand the orchestration.

With a facade:

```go
err := facade.PlaceOrder(order)
```

The orchestration logic exists in exactly one place.

---

# Advantages

- Simplifies client code
- Encapsulates workflow orchestration
- Reduces coupling between clients and subsystems
- Improves maintainability
- Centralizes error handling
- Easier to test through dependency injection

---

# Trade-offs

A facade should coordinate services, **not contain business logic**.

Avoid turning the facade into a large "god object" responsible for every operation in the system.

Keep it focused on providing a simplified interface to a related set of services.

---

# Production Considerations

This example intentionally stops at **workflow orchestration**.

The interfaces include rollback operations such as:

- `Release()`
- `Refund()`

These represent operations that may be required in a real production system, but they are **not invoked by this implementation**.

Why?

Because rollback and distributed transaction management belong to a different architectural pattern—the **Saga Pattern**.

Keeping these concerns separate allows the Facade pattern to remain focused on simplifying interactions, while the Saga pattern handles compensation and recovery.

---

# Real-World Usage

The Facade pattern is commonly used in backend systems.

Examples include:

- Checkout services
- Order processing
- User registration workflows
- Backend-for-Frontend (BFF) services
- Microservice orchestration
- Payment gateways
- Deployment orchestration
- Cloud provisioning services

---

# Running the Tests

Run the package tests.

```bash
go test ./structural/facade
```

Run with coverage.

```bash
go test ./structural/facade -cover
```

Run the race detector.

```bash
go test -race ./structural/facade
```

---

# Key Takeaways

- A Facade provides a simplified interface to multiple subsystems.
- Clients interact with a single entry point instead of coordinating multiple services.
- The pattern improves maintainability by centralizing orchestration logic.
- Error handling becomes consistent and easier to maintain.
- Long-running transactions and rollback logic are better handled by patterns such as **Saga**, keeping the Facade focused on orchestration.