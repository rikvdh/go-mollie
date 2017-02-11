package mollie

import (
	"strconv"
	"time"
)

type PaymentApi struct {
	c *core
}

type PaymentData struct {
	Amount      float64     `json:"amount,string"`
	Description string      `json:"description"`
	RedirectUrl string      `json:"redirectUrl"`
	WebhookUrl  string      `json:"webhookUrl"`
	Method      string      `json:"method,omitempty"`
	Metadata    interface{} `json:"metadata"`
	// One of de_DE en_US es_ES fr_FR nl_BE fr_BE nl_NL
	Locale string `json:"locale,omitempty"`

	Issuer string `json:"issuer,omitempty"`
}

type PaymentStatus string

const (
	STATUS_OPEN         PaymentStatus = "open"
	STATUS_CANCELLED    PaymentStatus = "cancelled"
	STATUS_EXPIRED      PaymentStatus = "expired"
	STATUS_FAILED       PaymentStatus = "failed"
	STATUS_PENDING      PaymentStatus = "pending"
	STATUS_PAID         PaymentStatus = "paid"
	STATUS_PAIDOUT      PaymentStatus = "paidout"
	STATUS_REFUNDED     PaymentStatus = "refunded"
	STATUS_CHARGED_BACK PaymentStatus = "charged_back"
)

type PaymentReply struct {
	Id                string
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
	ProfileId         string      `json:"profileId"`
	Links             map[string]string
}

type PaymentReplyWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []PaymentReply
}

func NewPayments(co *core) *PaymentApi {
	return &PaymentApi{c: co}
}

func (a *PaymentApi) New(data PaymentData) (*PaymentReply, error) {
	p := PaymentReply{}

	err := a.c.Post("payments", &p, &data)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (a *PaymentApi) Get(id string) (*PaymentReply, error) {
	p := PaymentReply{}

	err := a.c.Get("payments/"+id, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (a *PaymentApi) List(offset, limit uint64) ([]PaymentReply, error) {
	p := PaymentReplyWrapper{}

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
