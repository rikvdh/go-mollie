// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

import (
	"testing"
)

func TestIssuerList(t *testing.T) {
	is := newIssuers(&core{apiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	i, err := is.List()
	if err != nil {
		t.Errorf("Error is not nil\n")
	}
	if len(i) != 1 {
		t.Errorf("In testmode we expect exactly one issuer, got %d", len(i))
	}
}

func TestIssuerGet(t *testing.T) {
	is := newIssuers(&core{apiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	i, err := is.Get("non-existing-id")
	if i != nil {
		t.Errorf("This may not return an valid issuer ID\n")
	}
	if err == nil {
		t.Errorf("Error is nil")
	}

	i, err = is.Get("ideal_TESTNL99")
	if i.ID != "ideal_TESTNL99" {
		t.Errorf("We expect the ideal_TESTNL99 issuer, got: %s\n", i.ID)
	}
	if err != nil {
		t.Errorf("Error is not nil\n")
	}
}
