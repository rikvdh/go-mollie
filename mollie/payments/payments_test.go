package payments

import (
	"testing"
	"localhost/he/go-mollie-api/mollie/core"
)

func TestMethodList(t *testing.T) {
	is := NewPayments(core.Core{ApiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	p, err := is.New(PaymentData{
		Amount: 100.99,
		Description: "test descr",
		RedirectUrl: "http://www.google.com",
		WebhookUrl: "https://www.google.com/mollieapihook",
		Method: "ideal",
		Metadata: map[string]string{
			"rikvdh" : "mollie-api",
		},
	})
	if err != nil {
		t.Errorf("Error is not nil\n")
	}
	if len(p.Id) == 0 {
		t.Errorf("ID must not be empty\n")
	}
}
