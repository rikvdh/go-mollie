// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

// IssuerAPI type, holds reference to the core
type IssuerAPI struct {
	c *core
}

// Issuer for, for example, the iDeal payment method
type Issuer struct {
	ID       string
	Name     string
	Method   string
	Resource string
}

type issuerListWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []Issuer
}

func newIssuers(co *core) *IssuerAPI {
	return &IssuerAPI{c: co}
}

// List returns a full listing of all issuers
func (a *IssuerAPI) List() ([]Issuer, error) {
	var issuers issuerListWrapper
	err := a.c.Get("issuers", &issuers)

	if err != nil {
		return nil, err
	}

	return issuers.Data, nil
}

// Get returns a single issuer with the ID given
func (a *IssuerAPI) Get(issuerID string) (*Issuer, error) {
	var issuer Issuer
	err := a.c.Get("issuers/"+issuerID, &issuer)
	if err != nil {
		return nil, err
	}

	return &issuer, nil
}
