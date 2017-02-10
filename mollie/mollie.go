package mollie

import (
	"localhost/he/go-mollie-api/mollie/core"
	"localhost/he/go-mollie-api/mollie/issuers"
	"localhost/he/go-mollie-api/mollie/methods"
	"localhost/he/go-mollie-api/mollie/payments"
)

type Mollie struct {
	Issuers *issuers.IssuerApi
	Methods *methods.MethodApi
	Payments *payments.PaymentApi
}

func Get(apiKey string) Mollie {
	c := core.Core{ApiKey : apiKey}

	return Mollie{
		Issuers: issuers.NewIssuers(c),
		Methods: methods.NewMethods(c),
		Payments: payments.NewPayments(c),
	}
}
