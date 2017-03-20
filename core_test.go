package mollie

import (
	"errors"
	"testing"
)

type MarshalFail struct {
}

func (m MarshalFail) MarshalJSON() ([]byte, error) {
	return nil, errors.New("Marshal fail")
}

func TestPostFail(t *testing.T) {
	c := core{apiKey: ""}
	err := c.Post("", nil, MarshalFail{})
	if err == nil {
		t.Errorf("Error is nil")
	}
}
