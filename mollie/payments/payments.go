package payments

import (
	"localhost/he/go-mollie-api/mollie/core"
)

type PaymentApi struct {
	core core.Core
}

type PaymentData struct {
	Amount float64  `json:"amount,string"`
	Description string `json:"description"`
	RedirectUrl string `json:"redirectUrl"`
	WebhookUrl string `json:"webhookUrl"`
	Method string `json:"method,omitempty"`
	Metadata interface{} `json:"metadata"`
	// One of de_DE en_US es_ES fr_FR nl_BE fr_BE nl_NL
	Locale string `json:"locale,omitempty"`

	Issuer string `json:"issuer,omitempty"`
}

type PaymentReply struct {
	Id string
	Mode string
	CreatedDatetime string `json:"createdDatetime"`
	Status string
	ExpiryPeriod string `json:"expiryPeriod"`
	Amount float64 `json:",string"`
	Description string
	Method string
	Metadata interface{}
	Details interface{} `json:",omitempty"`
	ProfileId string `json:"profileId"`
	Links map[string]string
}

func NewPayments(c core.Core) *PaymentApi {
	return &PaymentApi{core: c}
}

func (a *PaymentApi) New(data PaymentData) (*PaymentReply, error) {
	p := PaymentReply{}

	err := a.core.Post("payments", &p, &data)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
