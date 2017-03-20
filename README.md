# Go Mollie API

[![Build Status](https://travis-ci.org/rikvdh/go-mollie.svg?branch=master)](https://travis-ci.org/rikvdh/go-mollie)
[![Go Report Card](https://goreportcard.com/badge/github.com/rikvdh/go-mollie)](https://goreportcard.com/report/github.com/rikvdh/go-mollie)
[![GoDoc](https://godoc.org/github.com/rikvdh/go-mollie?status.svg)](https://godoc.org/github.com/rikvdh/go-mollie)
[![codecov](https://codecov.io/gh/rikvdh/go-mollie/branch/master/graph/badge.svg)](https://codecov.io/gh/rikvdh/go-mollie)

The Go Mollie API is designed for use with the Mollie Payment Service Provider.
[Mollie](http://www.mollie.nl) is a Dutch Payment Service Provider.
This API is an interface to Mollie for [Golang](http://www.golang.org).

## Installation

Installation is very simple:

```sh
go get -u github.com/rikvdh/go-mollie
```

## Usage

A very basic example is shown below.
Just get an instance of the API with `mollie.Get("<API-KEY>")` and use one of the available API's.
Other examples are show in the `_examples` directory.

```golang
package main

import (
	"fmt"

	"github.com/rikvdh/go-mollie"
)

func main() {
	m := mollie.Get("<apikey>")

	methods, err := m.Methods().List()
	if err != nil {
		panic(err)
	}

	for _, method := range methods {
		fmt.Printf("method %s: %s\n", method.ID, method.Description)
	}
}
```

## Status

It works and I myself use it in production without issues.
Currently only authentication by API keys. See the [documentation of Mollie](https://www.mollie.com/en/docs/overview) for more information on the API's.
Currently implemented API's:

- [x] Payments API
- [x] Methods API
- [x] Issuers API
- [ ] Refunds API
- [x] Customers API
- [ ] Mandates API
- [ ] Subscriptions API

Other API's which require OAuth API authentication:

- [ ] Connect API
- [ ] Permissions API
- [ ] Organizations API
- [ ] Profiles API
- [ ] Settlements API
- [ ] Invoices API

## Contributing

Please read the [Contribution Guidelines](CONTRIBUTING.md). Furthermore: Fork -> Patch -> Push -> Pull Request

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/go-gitea/gitea/blob/master/LICENSE) file for the full license text.