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
- Docker-ready for deployment

---

## Tech Stack

- **Language:** Go (Golang)
- **Architecture:** Domain → Service → Handler → Main
- **Testing:** Go `testing` package with table-driven tests
- **Deployment:** Can run locally or via Docker
- **Optional UI:** Can integrate with front-end via `/calculate` endpoint

---

## Project Structure

