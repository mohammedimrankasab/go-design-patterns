# Chain of Responsibility Pattern

The **Chain of Responsibility** pattern allows multiple handlers to process a request sequentially. Each handler decides whether to process the request, stop the chain, or pass control to the next handler.

Unlike large conditional statements or tightly coupled request processing logic, this pattern keeps handlers independent, reusable, and easy to extend.

---

## Problem

As applications grow, request processing often involves multiple independent concerns:

- Logging
- Authentication
- Authorization
- Rate limiting
- Validation
- Business processing

Without the Chain of Responsibility pattern, these responsibilities become tightly coupled, making the code difficult to maintain and extend.

---

## Solution

Each handler performs one responsibility and decides whether the request should continue.

```
Incoming Request
        │
        ▼
    Logging
        │
  Authentication
        │
  Authorization
        │
  Rate Limiting
        │
  Validation
        │
  Business Handler
        │
    Response
```

If any handler returns an error, request processing stops immediately.

---

## Project Structure

```
chainofresponsibility/
├── chain.go
├── chain_test.go
├── README.md
└── examples/
    └── main.go
```

---

## Production Use Cases

This pattern is commonly used in:

- HTTP middleware pipelines (Gin, Chi, Echo)
- Authentication & Authorization
- API Gateways
- Reverse Proxies
- Kubernetes Admission Controllers
- Request Validation
- Logging
- Metrics Collection
- Distributed Tracing
- Retry Pipelines

---

## Advantages

- Single Responsibility Principle
- Open/Closed Principle
- Easily extensible
- Highly testable
- Reusable handlers
- Loose coupling
- Easy to reorder handlers

---

## Disadvantages

- Request flow becomes harder to debug in very long chains.
- Incorrect ordering may lead to security or business issues.
- Deep chains introduce additional execution overhead.

---

## When to Use

Use this pattern when:

- Multiple independent operations need to execute in sequence.
- The order of execution matters.
- Each step should be independently testable.
- Handlers should be reusable across different pipelines.

---

## When Not to Use

Avoid this pattern when:

- There is only one processing step.
- The execution order is fixed and unlikely to change.
- The processing logic is tightly coupled and cannot be meaningfully separated.

---

## Running the Example

```bash
go run ./examples
```

---

## Running Tests

```bash
go test ./...
```

---

## Related Patterns

- Strategy
- Command
- Decorator
- Observer