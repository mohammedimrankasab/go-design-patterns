# Singleton Pattern

## Overview

The Singleton pattern ensures that only one instance of an object exists throughout the application lifecycle.

In Go, the preferred implementation uses `sync.Once` to guarantee thread-safe initialization.

---

## Problem

Some resources should not have multiple instances:

Examples:

- Configuration manager
- Metrics registry
- Connection pool

Creating multiple instances can lead to:

- inconsistent state
- duplicated resources
- unnecessary memory usage

---

## Solution

The singleton pattern provides a single shared instance.

Example:

```go
config := singleton.GetInstance()
```

Every caller receives the same object.

---

## Go Implementation

Go provides `sync.Once` for safe lazy initialization.

Example:

```go
once.Do(func() {
    instance = &ConfigManager{}
})
```

The initialization function executes exactly once.

---

## Advantages

- Thread-safe initialization
- Lazy loading
- Shared resource management
- Prevents duplicate instances

---

## Trade-offs

Singletons should be used carefully.

Problems:

- introduces global state
- makes testing harder
- creates hidden dependencies

Prefer dependency injection when possible.

---

## Real World Usage

Appropriate examples:

- application configuration
- metrics registry
- shared caches
- resource managers