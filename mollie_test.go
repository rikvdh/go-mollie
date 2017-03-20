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
