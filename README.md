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

## Run Locally

Show available commands:

go run . --help

### Fetch current conditions

Defaults to Dublin coordinates if no flags are provided.

go run . forecast --lat 53.3498 --lon -6.2603

Example Output (truncated)

```json
{
  "latitude": 53.35,
  "longitude": -6.26,
  "current": {
    "time": "2024-01-01T00:00",
    "temperature_2m": 12.5,
    "wind_speed_10m": 5.4
  }
}
```

## Run tests

go test ./...

### Run tests with coverage

go test ./... -cover

## Lint (optional locally)

If you have golangci-lint installed:

golangci-lint run

Linting and tests are automatically run on every push and pull request via GitHub Actions.

## Project Structure

```text
.
├── cmd/ # Cobra commands (CLI entrypoints)
├── internal/api/ # API client and unit tests
├── main.go # CLI root
└── .github/workflows/ # GitHub Actions CI configuration
```

## Notes

- Data is fetched from the Open-Meteo public API.

- This project is intended as a learning and portfolio exercise for Go.
