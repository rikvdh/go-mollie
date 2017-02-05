package mollie

import (
	"localhost/he/go-mollie-api/mollie/core"
	"localhost/he/go-mollie-api/mollie/issuers"
)

type Mollie struct {
	Issuers *issuers.Issuers
}

func Get(apiKey string) Mollie {
	c := core.Core{ApiKey : apiKey}

	m := Mollie{}
	m.Issuers = issuers.NewIssuers(c)
	return m
}
