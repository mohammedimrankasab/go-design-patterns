# Contributing Guide

Thank you for your interest in contributing to Go Design Patterns.

This repository aims to provide clean, idiomatic, and production-oriented Go implementations of software design patterns.

---

## Development Setup

Requirements:

- Go 1.26+
- Make
- golangci-lint

Clone the repository:

```bash
git clone https://github.com/mohammedimrankasab/go-design-patterns.git

cd go-design-patterns
```

Verify installation:

```bash
go version
make help
```

---

# Development Workflow

Create a branch:

```bash
git checkout -b feature/pattern-name
```

Example:

```bash
git checkout -b feature/strategy-pattern
```

---

## Before Creating a Pull Request

Run:

```bash
make check
```

This validates:

- Formatting
- Static analysis
- Unit tests
- Race detection
- Linting

---

# Adding a New Pattern

Each pattern should follow this structure:

```
pattern-name/

├── README.md
├── implementation.go
└── implementation_test.go
```

---

## README Requirements

Each pattern documentation should include:

### Problem

What problem does this pattern solve?

### Solution

How does the pattern address the problem?

### Go Implementation

Explain the Go-specific approach.

### Example Usage

Show a practical example.

### Trade-offs

Explain when not to use the pattern.

---

# Code Guidelines

Follow Go best practices:

- Prefer simple designs
- Use interfaces carefully
- Avoid unnecessary abstractions
- Keep packages focused
- Add tests for all implementations
- Handle errors explicitly

---

# Commit Guidelines

Use meaningful commit messages.

Examples:

```
feat: add strategy pattern implementation

test: add worker pool tests

docs: improve factory pattern explanation
```

---

# Pull Requests

Pull requests should include:

- Clear description
- Motivation
- Implementation details
- Test coverage
- Documentation updates

Thank you for contributing.