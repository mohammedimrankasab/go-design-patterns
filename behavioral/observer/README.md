# Observer Pattern

The **Observer Pattern** defines a one-to-many relationship between objects. When the publisher emits an event, all registered observers are notified automatically without the publisher needing to know their implementation.

---

# Problem

As systems grow, a single business event often requires multiple independent actions.

For example, when an order is created:

- Send a confirmation email
- Reserve inventory
- Record analytics
- Write an audit log

Without the Observer pattern, the order service becomes tightly coupled to every downstream component.

---

# Solution

The publisher only emits an event.

Observers subscribe to events and react independently.

```
                Order Service
                      │
          Publish(OrderCreated)
                      │
        ┌─────────────┼─────────────┐
        ▼             ▼             ▼
 Email Observer  Inventory   Analytics
        │
        ▼
 Audit Observer
```

---

# Project Structure

```
observer/
├── observer.go
├── observer_test.go
├── README.md
└── examples/
    └── main.go
```

---

# Production Use Cases

The Observer pattern is commonly used in:

- Event-driven architectures
- Kafka consumers
- RabbitMQ subscribers
- NATS messaging
- Redis Pub/Sub
- Kubernetes Informers
- Webhooks
- Notification systems
- Monitoring and alerting
- Domain Events

---

# Advantages

- Loose coupling
- Easy to extend
- Open/Closed Principle
- Independent subscribers
- Improved maintainability
- Easy to test

---

# Disadvantages

- Notification order may matter
- Large numbers of observers can increase latency
- Error handling across observers requires careful consideration
- Debugging event flows may become more difficult

---

# When to Use

Use this pattern when:

- Multiple components react to the same event.
- New subscribers should be added without modifying existing code.
- The publisher should not know about downstream business logic.

---

# When Not to Use

Avoid this pattern when:

- There is only one consumer.
- The processing order is tightly coupled.
- Strong transactional consistency across all handlers is required.

---

# Running the Example

```bash
go run ./examples
```

---

# Running Tests

```bash
go test ./...
```

---

# Related Patterns

- Chain of Responsibility
- Command
- Mediator
- Publish/Subscribe (architectural pattern)