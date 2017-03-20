// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

import (
	"testing"
)

func TestMollie(t *testing.T) {
	m := Get("")

	if m.Issuers() == nil {
		t.Errorf("Issuers is nil")
	}
	if m.Methods() == nil {
		t.Errorf("Issuers is nil")
	}
	if m.Payments() == nil {
		t.Errorf("Issuers is nil")
	}
	if m.Customers() == nil {
		t.Errorf("Customers is nil")
	}
}
