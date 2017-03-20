// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

import (
	"strconv"
	"time"
)

// CustomerAPI is the part of the API which handles Mollie customers
type CustomerAPI struct {
	c *core
}

type Links struct {
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
	First    string `json:"first,omitempty"`
	Last     string `json:"last,omitempty"`
}

type NewCustomerData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Locale   string `json:"locale,omitempty"` // One of de_DE en_US es_ES fr_FR nl_BE fr_BE nl_NL
	Metadata string `json:"metadata,omitempty"`
}

type ListCustomersResponse struct {
	Totalcount uint       `json:"totalcount"`
	Offset     uint       `json:"offset"`
	Count      uint       `json:"count"`
	Data       []Customer `json:"data"`
	Links      Links      `json:"links,omitempty"`
}

type Customer struct {
	Resource            string    `json:"resource"`
	ID                  string    `json:"id"`
	Mode                string    `json:"mode"`
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	Locale              string    `json:"locale"`
	Metadata            string    `json:"metadata,omitempty"`
	RecentlyUsedMethods []string  `json:"recentlyUsedMethods"`
	CreatedDatetime     time.Time `json:"createdDatetime"`
}

func newCustomers(co *core) *CustomerAPI {
	return &CustomerAPI{co}
}

// New creates a new customer
func (api *CustomerAPI) New(data NewCustomerData) (*Customer, error) {
	customer := &Customer{}

	err := api.c.Post("customers", customer, data)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// Get retrieves customer information with the specified ID
func (api *CustomerAPI) Get(id string) (*Customer, error) {
	customer := &Customer{}

	err := api.c.Get("customers/"+id, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// List returns a list of customers
func (api *CustomerAPI) List(offset, limit uint64) (*ListCustomersResponse, error) {
	uri := "customers?offset="
	uri += strconv.FormatUint(offset, 10)
	uri += "&count="
	uri += strconv.FormatUint(limit, 10)

	response := &ListCustomersResponse{}
	err := api.c.Get(uri, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// NewPayment creates a new payment for a customer with the specified ID
func (api *CustomerAPI) NewPayment(id string, data PaymentData) (*PaymentReply, error) {
	payment := &PaymentReply{}

	err := api.c.Post("customers/"+id+"/payments", payment, data)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// Payments returns a list of payments for a customer with the specified ID
func (api *CustomerAPI) Payments(id string, offset uint64, limit uint64) ([]PaymentReply, error) {
	uri := "customers/" + id + "/payments?offset="
	uri += strconv.FormatUint(offset, 10)
	uri += "&count="
	uri += strconv.FormatUint(limit, 10)

	response := &paymentReplyWrapper{}
	err := api.c.Get(uri, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
