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
