# Praemunio Go SDK

Go client for the Praemunio public API.

## Installation

```bash
go get github.com/idcoopp-praemunio/praemunio-go-sdk
```

## Quick start

```go
package main

import (
    "fmt"
    "log"

    praemunio "github.com/idcoopp-praemunio/praemunio-go-sdk"
)

func main() {
    client := praemunio.NewPraemunioClient("pk_test_...", "sk_test_...")

    if err := client.Authenticate(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Access token:", client.GetAccessToken())

    // Create an API client
    api := client.CreateAPI()
    // use api to call endpoints...
}
```

## Environments

```go
sandbox := praemunio.NewPraemunioClientSandbox(publicKey, secretKey)
live := praemunio.NewPraemunioClientLive(publicKey, secretKey)
```

## Authentication

The SDK exchanges `publicKey` and `secretKey` for an access token and refresh token via `POST /auth/auth-token`.

```go
if err := client.Authenticate(); err != nil {
    log.Fatal(err)
}
```

## Build

```bash
go build ./...
```

## Notes

- Generated from `sdks/shared/openapi/praemunio.openapi.json`
- Base URLs:
  - Sandbox: `https://sandbox-api.praemunio.org/api/v1`
  - Live: `https://api.praemunio.org/api/v1`