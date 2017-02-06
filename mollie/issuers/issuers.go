package issuers

import (
	"localhost/he/go-mollie-api/mollie/core"
)

type IssuerApi struct {
	core core.Core
}

type Issuer struct {
	Id string
	Name string
	Method string
	Resource string
}

type IssuerListWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []Issuer
}

func NewIssuers(c core.Core) *IssuerApi {
	return &IssuerApi{core : c}
}

func (a *IssuerApi) List() ([]Issuer, error) {
	var issuers IssuerListWrapper
	err := a.core.Request("issuers", &issuers)

	if err != nil {
		return nil, err
	}

	return issuers.Data, nil
}

func (a *IssuerApi) Get(issuerId string) (*Issuer, error) {
	var issuer Issuer
	err := a.core.Request("issuers/" + issuerId, &issuer)
	if err != nil {
		return nil, err
	}

	return &issuer, nil
}
