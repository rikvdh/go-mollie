package issuers

import (
	"github.com/rikvdh/go-mollie-api/core"
	"testing"
)

func TestIssuerList(t *testing.T) {
	is := NewIssuers(core.Core{ApiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	i, err := is.List()
	if err != nil {
		t.Errorf("Error is not nil\n")
	}
	if len(i) != 1 {
		t.Errorf("In testmode we expect exactly one issuer, got %d", len(i))
	}
}

func TestIssuerGet(t *testing.T) {
	is := NewIssuers(core.Core{ApiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	i, err := is.Get("non-existing-id")
	if i != nil {
		t.Errorf("This may not return an valid issuer ID\n")
	}
	if err == nil {
		t.Errorf("Error is nil")
	}

	i, err = is.Get("ideal_TESTNL99")
	if i.Id != "ideal_TESTNL99" {
		t.Errorf("We expect the ideal_TESTNL99 issuer\n")
	}
	if err != nil {
		t.Errorf("Error is not nil\n")
	}
}
