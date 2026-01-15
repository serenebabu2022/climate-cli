# Climate CLI (Go)

A small Go (Golang) command-line tool that fetches climate/weather data from a public HTTP API (Open-Meteo).  
Built to practice Go backend fundamentals: HTTP clients, JSON parsing, testing, and CI.

## Features

- Fetch current conditions for a given latitude/longitude
- Simple, flag-based CLI commands
- Unit tests using `httptest`
- GitHub Actions CI: `go test` + `golangci-lint`

## Tech Stack

- Go (Golang)
- Cobra (CLI framework)
- net/http (HTTP client)
- GitHub Actions (CI)
- golangci-lint (linting)

## Getting Started

### Prerequisites

- Go installed (recommended: Go 1.22+)

### Clone

```bash
git clone https://github.com/serenebabu2022/climate-cli.git
cd climate-cli
```
