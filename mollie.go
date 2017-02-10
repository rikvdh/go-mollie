package mollie

import (
	"github.com/rikvdh/go-mollie-api/core"
	"github.com/rikvdh/go-mollie-api/issuers"
	"github.com/rikvdh/go-mollie-api/methods"
	"github.com/rikvdh/go-mollie-api/payments"
)

type Mollie struct {
	Issuers  *issuers.IssuerApi
	Methods  *methods.MethodApi
	Payments *payments.PaymentApi
}

func Get(apiKey string) Mollie {
	c := core.Core{ApiKey: apiKey}

	return Mollie{
		Issuers:  issuers.NewIssuers(c),
		Methods:  methods.NewMethods(c),
		Payments: payments.NewPayments(c),
	}
}
