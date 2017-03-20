// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
