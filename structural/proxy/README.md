# Proxy Pattern

## Overview

The **Proxy Pattern** provides a surrogate or placeholder for another object to control access to it.

Instead of clients interacting directly with the real service, they communicate through a proxy that can add additional behavior such as authorization, caching, logging, lazy initialization, or rate limiting.

In this example, an **Authorization Proxy** protects access to a document service by verifying user permissions before forwarding the request.

---

# Problem

A document service contains protected resources.

Without a proxy, every caller must implement its own authorization logic.

```text
Client
    │
    ▼
Document Service
```

This leads to duplicated security checks and increases the risk of inconsistent access control.

---

# Solution

Introduce a proxy that sits between the client and the real service.

```text
              Client
                 │
                 ▼
       Authorization Proxy
                 │
                 ▼
        Document Service
```

The proxy validates access before delegating the request to the underlying service.

```go
proxy := NewAuthorizationProxy(
    authService,
    documentService,
)

document, err := proxy.GetDocument(
    "user-1",
    "doc-1",
)
```

---

# Implementation

The proxy implements the same interface as the real service.

```go
type DocumentService interface {
    GetDocument(userID, documentID string) (*Document, error)
}
```

Clients remain unaware of whether they are communicating with the real service or a proxy.

---

# Workflow

Authorized request:

```text
Client
   │
   ▼
Authorization Proxy
   │
Permission Granted
   │
   ▼
Document Service
   │
   ▼
Document Returned
```

Unauthorized request:

```text
Client
   │
   ▼
Authorization Proxy
   │
Permission Denied
   │
   ▼
Error Returned
```

---

# Advantages

- Centralizes access control
- Reduces duplicated authorization logic
- Improves security
- Follows the Open/Closed Principle
- Clients remain unaware of the proxy

---

# Trade-offs

- Introduces an additional layer
- Can increase complexity if overused
- Multiple stacked proxies may impact readability

---

# Real-World Usage

Proxy is commonly used for:

- API Gateways
- Authentication middleware
- Authorization (RBAC/ABAC)
- Reverse proxies
- Service meshes
- Database connection proxies
- Lazy loading
- Caching layers
- Rate limiting

---

# Running the Tests

```bash
go test ./structural/proxy
```

```bash
go test ./structural/proxy -cover
```

```bash
go test -race ./structural/proxy
```

---

# Key Takeaways

- A Proxy controls access to another object.
- The proxy and the real service implement the same interface.
- Clients interact with the proxy transparently.
- Common responsibilities include authorization, caching, logging, rate limiting, and lazy initialization.