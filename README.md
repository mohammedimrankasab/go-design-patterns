# Go Design Patterns

Production-oriented implementations of software design patterns using idiomatic Go.

This repository focuses on understanding **why and when to use a pattern**, along with practical implementation details, trade-offs, and testing strategies.

---

## Philosophy

Design patterns are not about adding unnecessary complexity. They are tools to solve recurring engineering problems.

This repository follows Go engineering principles:

- Composition over inheritance
- Interfaces over unnecessary coupling
- Explicit dependencies
- Small and focused abstractions
- Simple solutions over premature complexity
- Code that is easy to test and maintain

---

# Patterns Covered

## Creational Patterns

Patterns focused on object creation and dependency management.

| Pattern | Description |
|---|---|
| Singleton | Managing a single shared instance |
| Factory | Creating objects without exposing implementation details |
| Builder | Constructing complex objects step by step |
| Functional Options | Idiomatic Go configuration pattern |

---

## Structural Patterns

Patterns focused on composing objects and simplifying relationships.

| Pattern | Description |
|---|---|
| Adapter | Making incompatible interfaces work together |
| Decorator | Adding behavior dynamically |
| Facade | Providing a simplified interface over complex systems |
| Proxy | Controlling access to another object |

---

## Behavioral Patterns

Patterns focused on communication and responsibility handling.

| Pattern | Description |
|---|---|
| Strategy | Selecting algorithms dynamically |
| Observer | Event notification mechanism |
| Command | Encapsulating actions as objects |
| Chain of Responsibility | Middleware-style request processing |
| State | Changing behavior based on internal state |

---

## Concurrency Patterns

Go-specific patterns for building reliable concurrent systems.

| Pattern | Description |
|---|---|
| Worker Pool | Controlled parallel task execution |
| Pipeline | Multi-stage data processing |
| Fan-In / Fan-Out | Concurrent stream processing |
| Rate Limiter | Controlling request throughput |

---

# Repository Structure

```text
go-design-patterns
│
├── creational
├── structural
├── behavioral
├── concurrency
├── examples
├── docs
└── .github
```

Each pattern contains:

```text
pattern-name/

├── README.md
├── implementation.go
└── implementation_test.go
```

---

# Running Tests

Run all tests:

```bash
go test ./...
```

Run tests with race detection:

```bash
go test -race ./...
```

Run static analysis:

```bash
go vet ./...
```

---

# Documentation

Additional engineering notes are available under:

```text
docs/
```

Topics include:

- Design principles
- Go-specific patterns
- Pattern trade-offs
- Production usage considerations

---

# Goals

This repository demonstrates:

- Practical Go design skills
- Production-oriented thinking
- Clean architecture principles
- Concurrency design patterns
- Testing practices

---

# License

MIT License