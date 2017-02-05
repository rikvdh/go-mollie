package issuers

import (
	"fmt"
	"localhost/he/go-mollie-api/mollie/core"
)

type Issuers struct {
	core core.Core
}
type Issuer struct {
	Id string
	Name string
	Method string
}


func NewIssuers(c core.Core) *Issuers {
	return &Issuers{core : c}
}

func (a *Issuers) List() {
	var issuers []Issuer
	a.core.Request("issuers", &issuers)
	fmt.Println(issuers)
}
