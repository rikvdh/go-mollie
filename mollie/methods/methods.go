package methods

import (
	"localhost/he/go-mollie-api/mollie/core"
)

type MethodApi struct {
	core core.Core
}

type MethodAmount struct {
	Minimum string
	Maximum string
}

type MethodImage struct {
	Normal string
	Buffer string
}

type Method struct {
	Id string
	Description string
	Amount MethodAmount
	Image MethodImage
	Resource string
}

type MethodListWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []Method
}

func NewMethods(c core.Core) *MethodApi {
	return &MethodApi{core : c}
}

func (a *MethodApi) List() ([]Method, error) {
	var methods MethodListWrapper
	err := a.core.Request("methods", &methods)

	if err != nil {
		return nil, err
	}

	return methods.Data, nil
}

func (a *MethodApi) Get(methodId string) (*Method, error) {
	var method Method
	err := a.core.Request("methods/" + methodId, &method)
	if err != nil {
		return nil, err
	}

	return &method, nil
}
