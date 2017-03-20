// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

import (
	"math/rand"
	"testing"
	"time"
)

func TestCustomersCreateAndGetError(t *testing.T) {
	mollieAPI := Get("test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU")

	rand.Seed(time.Now().Unix())
	id := rand.Intn(999999-1) + 1

	userEmail := "user_" + string(id)

	data := NewCustomerData{
		Name:   "test_user",
		Email:  userEmail,
		Locale: "nl_NL",
	}

	customer, err := mollieAPI.Customers().New(data)

	if err == nil {
		t.Errorf("Error is nil: %v", err)
	}

	if customer != nil {
		t.Errorf("Something may went wrong unmarchalling the customer\n")
	}

	// get that user
	q, err := mollieAPI.Customers().Get("user1")
	if err == nil {
		t.Errorf("Error is not nil\n")
	}
	if q != nil {
		t.Errorf("Something may went wrong unmarchalling the customer\n")
	}
}
