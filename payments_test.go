// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

import (
	"testing"
)

func TestPaymentAddGet(t *testing.T) {
	is := newPayments(&core{apiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	p, err := is.New(PaymentData{
		Amount:      100.99,
		Description: "test descr",
		RedirectURL: "http://www.google.com",
		WebhookURL:  "https://www.google.com/mollieapihook",
		Method:      "ideal",
		Metadata: map[string]string{
			"rikvdh": "mollie-api",
		},
	})
	if err != nil {
		t.Errorf("Error is not nil\n")
	}
	if len(p.ID) == 0 {
		t.Errorf("ID must not be empty\n")
	}
	if p.Amount != 100.99 {
		t.Errorf("Mollie may not change our amount! :O")
	}

	q, err := is.Get(p.ID)
	if err != nil {
		t.Errorf("Error is not nil\n")
	}
	if len(q.ID) == 0 {
		t.Errorf("ID must not be empty\n")
	}
	if q.Amount != 100.99 {
		t.Errorf("Mollie may not change our amount! :O")
	}

}

func TestPaymentGetError(t *testing.T) {
	is := newPayments(&core{apiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	q, err := is.Get("foo_bar")
	if err == nil {
		t.Errorf("Error is not nil\n")
	}
	if q != nil {
		t.Errorf("No payment may be retrieved\n")
	}
}

func TestPaymentList(t *testing.T) {
	is := newPayments(&core{apiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	q, err := is.List(0, 2)
	if err != nil {
		t.Errorf("Error is not nil: %s\n", err)
	}
	if len(q) != 2 {
		t.Errorf("ID must not be empty\n")
	}
	w, err := is.List(0, 1)
	if err != nil {
		t.Errorf("Error is not nil: %s\n", err)
	}
	if len(w) != 1 {
		t.Errorf("ID must not be empty\n")
	}
	e, err := is.List(1, 1)
	if err != nil {
		t.Errorf("Error is not nil: %s\n", err)
	}
	if len(e) != 1 {
		t.Errorf("ID must not be empty\n")
	}
}
