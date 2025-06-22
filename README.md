# ludo

## Run

```sh
# Run board

cd board
go run cmd/main.go
```

```sh
# Run dice

cd dice
cenv fix
go run cmd/main.go
```

## Structure

- `/board`: Local web view
- `/dice`: Ludo API
  - `/cmd`: Entry point
  - `/config`: API configuration
  - `/server`: Endpoint
  - `/ludo`: Service for Ludo actions
  - `/database`: Database repo for queries and models
  - `/github`: Service for interacting with the GitHub API
  - `/git`: Service for local git actions
