# Functional Options Pattern

## Overview

Functional Options is an idiomatic Go pattern used to configure objects while keeping APIs clean and extensible.

---

## Problem

Constructor parameters become difficult to maintain as configuration grows.

Example:

```go
NewServer(
    "localhost",
    8080,
    30,
    true,
)
```

Problems:

- unclear arguments
- difficult maintenance
- breaking changes when adding options

---

## Solution

Functional options use functions that modify configuration.

Example:

```go
server := NewServer(
    WithPort(8080),
    WithTimeout(time.Second*30),
)
```

---

## Implementation

The pattern consists of:

### Option type

```go
type Option func(*Server)
```

An option receives the object and modifies it.

---

### Constructor

```go
func NewServer(options ...Option)
```

The constructor applies provided options.

---

## Advantages

- Clean API
- Optional configuration
- Backward compatible
- Easy extension

---

## Trade-offs

Use carefully when:

- configuration is very small
- every option is mandatory
- validation becomes complicated

---

## Real World Usage

Commonly used in:

- HTTP clients
- database libraries
- cloud SDKs
- Kubernetes controllers
- logging frameworks