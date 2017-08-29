// Copyright 2017 The Go-Mollie Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mollie

import (
	"strconv"
	"time"
)

// PaymentAPI structure holds a pointer to the core
type PaymentAPI struct {
	c *core
}

// PaymentData needed for a payment request
type PaymentData struct {
	Amount      float64     `json:"amount,string"`
	Description string      `json:"description"`
	RedirectURL string      `json:"redirectUrl"`
	WebhookURL  string      `json:"webhookUrl"`
	Method      string      `json:"method,omitempty"`
	Metadata    interface{} `json:"metadata"`
	// One of de_DE en_US es_ES fr_FR nl_BE fr_BE nl_NL
	Locale string `json:"locale,omitempty"`

	Issuer string `json:"issuer,omitempty"`
}

// PaymentStatus for a payment
type PaymentStatus string

const (
	// StatusOpen means a payment is open for payment
	StatusOpen PaymentStatus = "open"
	// StatusCancelled means a payment is cancelled by a user
	StatusCancelled PaymentStatus = "cancelled"
	// StatusExpired means a payment is expired and not accepting payments anymore
	StatusExpired PaymentStatus = "expired"
	// StatusFailed means a payment has failed
	StatusFailed PaymentStatus = "failed"
	// StatusPending means a payment is on hold and not confirmed (yet)
	StatusPending PaymentStatus = "pending"
	// StatusPaid informs you that a payment is marked as paid
	StatusPaid PaymentStatus = "paid"
	// StatusPaidout means ?
	StatusPaidout PaymentStatus = "paidout"
	// StatusRefunded means this payment is refunded
	StatusRefunded PaymentStatus = "refunded"
	// StatusChargedBack means a payment is charged back by the customer
	StatusChargedBack PaymentStatus = "charged_back"
)

// PaymentReply for a payment
type PaymentReply struct {
	ID                string
	Mode              string
	CreatedDatetime   time.Time `json:"createdDatetime"`
	ExpiredDatetime   time.Time `json:"expiredDatetime"`
	CancelledDatetime time.Time `json:"cancelledDatetime"`
	PaidDatetime      time.Time `json:"paidDatetime"`
	Status            PaymentStatus
	ExpiryPeriod      string  `json:"expiryPeriod"`
	Amount            float64 `json:",string"`
	AmountRefunded    float64 `json:"amountRefunded,string,omitempty"`
	AmountRemaining   float64 `json:"amountRemaining,string,omitempty"`
	Description       string
	Method            string
	Metadata          interface{}
	Details           interface{} `json:",omitempty"`
	ProfileID         string      `json:"profileId"`
	Links             map[string]string
}

type paymentReplyWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []PaymentReply
}

func newPayments(co *core) *PaymentAPI {
	return &PaymentAPI{c: co}
}

// New Creates a new payment
func (a *PaymentAPI) New(data PaymentData) (*PaymentReply, error) {
	p := PaymentReply{}

	err := a.c.Post("payments", &p, &data)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Get retrieves a payment with the given ID
func (a *PaymentAPI) Get(id string) (*PaymentReply, error) {
	p := PaymentReply{}

	err := a.c.Get("payments/"+id, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Delete a payment with the given ID
func (a *PaymentAPI) Delete(id string) (*PaymentReply, error) {
	p := PaymentReply{}

	err := a.c.Delete("payments/"+id, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// List returns a list of payments
func (a *PaymentAPI) List(offset, limit uint64) ([]PaymentReply, error) {
	p := paymentReplyWrapper{}

	uri := "payments?offset="
	uri += strconv.FormatUint(offset, 10)
	uri += "&count="
	uri += strconv.FormatUint(limit, 10)

	err := a.c.Get(uri, &p)
	if err != nil {
		return nil, err
	}

	return p.Data, nil
}
