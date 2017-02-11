package mollie

type Mollie struct {
	issuers  *IssuerApi
	methods  *MethodApi
	payments *PaymentApi
}

func Get(apikey string) Mollie {
	c := core{apiKey: apikey}

	return Mollie{
		issuers:  NewIssuers(&c),
		methods:  NewMethods(&c),
		payments: NewPayments(&c),
	}
}

func (m Mollie) Issuers() *IssuerApi {
	return m.issuers
}

func (m Mollie) Methods() *MethodApi {
	return m.methods
}

func (m Mollie) Payments() *PaymentApi {
	return m.payments
}
