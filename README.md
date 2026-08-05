# Go Design Patterns: Production-Ready Examples

> A curated collection of Go design patterns implemented using real-world backend engineering examples, idiomatic Go practices, comprehensive tests, and CI/CD.

![Go Version](https://img.shields.io/badge/Go-1.26-blue)
![License](https://img.shields.io/badge/License-MIT-green)

---

## Why This Repository?

Most Go design pattern repositories demonstrate concepts using toy examples such as **Animals**, **Shapes**, or **Cars**.

While these examples explain the pattern, they rarely show how it is applied in production software.

This repository focuses on **real-world backend engineering scenarios**, making it useful for:

- Backend Engineers
- Go Developers
- Interview Preparation
- System Design Learning
- Production Software Development

Every implementation is designed to be:

- Idiomatic Go
- Dependency Injection friendly
- Production-inspired
- Well documented
- Fully tested
- CI verified

---

# Repository Structure

```text
go-design-patterns
│
├── creational
│   ├── builder
│   ├── factory
│   ├── functional-options
│   └── singleton
│
├── structural
│   ├── adapter
│   ├── decorator
│   ├── facade
│   └── proxy
│
├── behavioral
│   ├── chain
│   ├── command
│   ├── observer
│   └── strategy
│
├── concurrency
│   ├── worker-pool
│   ├── pipeline
│   ├── fan-in-fan-out
│   └── rate-limiter
│
└── docs
```

---

# Implemented Patterns

## Creational

| Pattern | Status | Production Example |
|----------|--------|--------------------|
| Functional Options | ✅ | Configurable server initialization |
| Builder | ✅ | HTTP client construction |
| Factory | ✅ | Notification providers |
| Singleton | ✅ | Configuration manager |

---

## Structural

| Pattern | Status | Production Example |
|----------|--------|--------------------|
| Adapter | ✅ | Legacy payment gateway |
| Decorator | ✅ | Logging & Metrics wrappers |
| Facade | ✅ | Order orchestration |
| Proxy | 🚧 | Authorization proxy |

---

## Behavioral

Coming soon.

- Strategy
- Observer
- Command
- Chain of Responsibility

---

## Concurrency

Coming soon.

- Worker Pool
- Pipeline
- Fan-In / Fan-Out
- Rate Limiter

---

# Running Examples

Clone the repository.

```bash
git clone https://github.com/mohammedimrankasab/go-design-patterns.git

cd go-design-patterns
```

Run all tests.

```bash
go test ./...
```

Run the race detector.

```bash
go test -race ./...
```

Run code quality checks.

```bash
make check
```

---

# Learning Path

If you're new to design patterns, follow this order.

### Beginner

- Functional Options
- Factory
- Builder

### Intermediate

- Singleton
- Adapter
- Decorator
- Facade
- Proxy

### Advanced

- Strategy
- Observer
- Command
- Chain of Responsibility

### Go Concurrency

- Worker Pool
- Pipeline
- Fan-In / Fan-Out
- Rate Limiter

---

# Why These Examples?

Every implementation uses a realistic backend engineering scenario.

Examples include:

- Payment gateways
- Notification services
- Configuration managers
- HTTP middleware
- Logging wrappers
- Metrics instrumentation
- Order orchestration

These are patterns you are likely to encounter in production Go applications.

---

# Repository Goals

This project aims to provide:

- Clear explanations
- Idiomatic Go implementations
- Production-inspired examples
- Comprehensive unit tests
- High code quality
- Interview-ready reference implementations

---

# Contributing

Contributions are welcome.

Please read:

- CONTRIBUTING.md
- CODE_OF_CONDUCT.md

before opening an issue or pull request.

---

# Roadmap

## Phase 1

- [x] Creational Patterns
- [x] Structural Patterns

## Phase 2

- [ ] Behavioral Patterns
- [ ] Concurrency Patterns

## Phase 3

Production Engineering Patterns

- [ ] Retry
- [ ] Circuit Breaker
- [ ] Saga
- [ ] Outbox
- [ ] Idempotency
- [ ] Cache Aside
- [ ] CQRS

---

# License

This project is licensed under the MIT License.