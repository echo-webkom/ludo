<div align="center">

<h1>Ludo</h1>

<img src=".github/mock.png" width="80%">

</div>

<br>

## Run

```sh
# Run board

cd board
go run cmd/main.go
```

```sh
# Run api

cd api
cenv fix
go run cmd/main.go
```

## Structure

- `/board`: Local web view
- `/api`: Ludo API
  - `/cmd`: Entry point
  - `/config`: API configuration
  - `/server`: Endpoint
  - `/ludo`: Service for Ludo actions
  - `/database`: Database repo for queries and models
  - `/github`: Service for interacting with the GitHub API
  - `/git`: Service for local git actions
