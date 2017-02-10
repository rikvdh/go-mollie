package methods

import (
	"testing"
	"localhost/he/go-mollie-api/mollie/core"
)

func TestMethodList(t *testing.T) {
	is := NewMethods(core.Core{ApiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	i, err := is.List()
	if err != nil {
		t.Errorf("Error is not nil\n")
	}
	if len(i) != 11 {
		t.Errorf("In testmode we expect exactly one method, got %d", len(i))
	}
}

func TestMethodGet(t *testing.T) {
	is := NewMethods(core.Core{ApiKey: "test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU"})

	i, err := is.Get("non-existing-id")
	if i != nil {
		t.Errorf("This may not return an valid method ID\n")
	}
	if err == nil {
		t.Errorf("Error is nil")
	}

	// This is not tested, appearantly all methods are disabled.. -_-'
	/*i, err = is.Get("creditcard")
	if err != nil {
		t.Errorf("Error is not nil\n")
	}*/
}
