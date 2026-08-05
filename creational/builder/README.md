# Builder Pattern

## Overview

The Builder pattern separates the construction of a complex object from its representation.

It allows objects to be created step by step while keeping construction readable.

---

## Problem

Objects with many configuration options often lead to:

- long constructors
- unclear parameters
- difficult maintenance

Example:

```go
NewHTTPClient(
    url,
    timeout,
    retries,
    headers,
    tracing,
)
```

This becomes difficult to understand as configuration grows.

---

## Solution

The Builder pattern provides a fluent API.

Example:

```go
client := NewClientBuilder().
    WithBaseURL("https://api.example.com").
    WithTimeout(30*time.Second).
    WithRetries(3).
    Build()
```

---

## Implementation

The builder contains:

### Builder Object

Maintains intermediate configuration.

```go
type ClientBuilder struct
```

---

### Fluent Methods

Each method modifies configuration and returns the builder.

```go
func (b *ClientBuilder) WithTimeout(...)
```

---

### Build Method

Creates the final object.

```go
func (b *ClientBuilder) Build() HTTPClient
```

---

## Advantages

- Readable object creation
- Avoids huge constructors
- Supports optional configuration
- Improves maintainability

---

## Trade-offs

Avoid using builders when:

- objects have only a few fields
- simple constructors are enough
- configuration is unlikely to grow

---

## Real World Usage

Common examples:

- HTTP clients
- Database clients
- Cloud SDK configuration
- Deployment configuration
- Complex request objects