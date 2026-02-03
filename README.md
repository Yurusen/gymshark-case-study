# Gymshark Pack Calculator

A Go-based HTTP API that calculates the optimal combination of product packs for customer orders, following Gymshark’s shipping rules:

1. Only whole packs can be sent. Packs cannot be broken open.
2. Send no more items than necessary to fulfill the order.
3. Use as few packs as possible when multiple options exist.

This project is designed as a **backend service** with clean separation of concerns, ready for deployment and testing.

---

## Features

- Configurable pack sizes (default: 250, 500, 1000, 2000, 5000)
- Calculates minimal oversupply and minimal number of packs
- HTTP API endpoint for integration with front-end or other services
- Fully testable service layer
- Easy to extend for additional pack sizes or business rules
  
---

## Tech Stack

- **Language:** Go (Golang)
- **Architecture:** Domain → Service → Handler → Main
- **Testing:** Go `testing` package with table-driven tests
- **Deployment:** Can run locally or is currently hosted via Railway
- **Optional UI:** Can integrate with front-end via `/calculate` endpoint

---

## Project Structure

- `cmd/api/main.go` – the main entry point of the application.
- `internal/config/config.go` – handles application configuration.
- `internal/domain/pack.go` – defines the core domain models related to packs.
- `internal/service/pack_calculator.go` – contains the main business logic for calculating the optimal packs.
- `internal/handler/calculate_handler.go` – exposes the HTTP endpoint for pack calculations.
- `tests/pack_calculator_test.go` – unit tests for the service layer.
- `go.mod` – Go module file.
- `README.md` – this file.

---

## API

This project exposes a single HTTP endpoint:

### `POST /calculate`

Calculates the optimal combination of product packs to fulfill an order according to Gymshark’s shipping rules.

**Request Body:**

- `items` (integer, required) – the total number of items in the order.  
- `packSizes` (array of integers, optional) – a custom list of pack sizes to use instead of the default `[250, 500, 1000, 2000, 5000]`.

**Example Requests:**

```json
// Using default pack sizes
{
  "items": 12
}

// Using custom pack sizes
{
  "items": 12,
  "packSizes": [12, 145]
}
