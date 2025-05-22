# go-skydio

A [golang](https://go.dev/) SDK for the [Skydio Cloud API](https://apidocs.skydio.com/reference/introduction).

Skydio Cloud allows you to manage your Skydio vehicles using Fleet Manager and Media Sync to provide automatic media and flight data uploads.

## Installation

go-skydio is compatible with Go releases in module mode:

```sh
go get -u github.com/justinsimmons/go-skydio
```

## Usage

First initialize a Skydio API client. This will allow you to interface with the API in an easy manner.

```go
import "github.com/justinsimmons/go-coinbase"

func main() {
    client := skydio.NewAuthenticatedClient(token)
}
```

Once you have the client you can perform any of the various API operations:

```go
// TODO
```

The services of the client divide the API into logical chunks and correspond to the structure of the Skydio Cloud API [documentation](https://apidocs.skydio.com/reference/introduction).

NOTE: Using the [context](https://godoc.org/context) package, one can easily pass cancellation signals and deadlines to various services of the client for handling a request. In case there is no context available, then `context.Background()` can be used as a starting point.

## Authentication

Interacting with the Skydio API requires an API Token for both authenticating and authorizing requests.

To create an API token follow Skydio's instructions located [here](https://apidocs.skydio.com/reference/authentication#creating-an-api-token).

Once you have the API token you can supply it to the Go Skydio client one of two ways. Both are equivalent.

```go
// This is your API token.
token := "super secret"

// You can pass the API token in directly as an option to the constructor.
client := skydio.NewClient(skydio.WithApiKey(token))

// Or you can use the provided convenience function:
client = skydio.NewAuthenticatedClient(token)
```

## Pagination

## Liscense

Copyright 2025 Justin Simmons.

This program is released under the [TBD]() or later.
