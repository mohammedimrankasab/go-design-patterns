# Strategy Pattern in Go

The **Strategy Pattern** is a behavioral design pattern that allows an application to choose an algorithm or implementation at runtime.

Instead of embedding multiple conditional branches inside a class or service, each implementation is encapsulated behind a common interface.

---

## Problem

Imagine building an enterprise metadata ingestion platform.

The platform needs to ingest metadata from multiple external systems such as:

- Power BI
- Tableau
- MLflow
- Snowflake
- Databricks

Each system exposes completely different APIs, authentication mechanisms, response formats, pagination strategies, and rate limits.

Without the Strategy Pattern the code usually becomes:

```go
switch connector {
case "powerbi":
    ...
case "tableau":
    ...
case "mlflow":
    ...
}
```