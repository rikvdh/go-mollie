package mollie

type IssuerApi struct {
	c *core
}

type Issuer struct {
	Id       string
	Name     string
	Method   string
	Resource string
}

type IssuerListWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []Issuer
}

func NewIssuers(co *core) *IssuerApi {
	return &IssuerApi{c: co}
}

func (a *IssuerApi) List() ([]Issuer, error) {
	var issuers IssuerListWrapper
	err := a.c.Get("issuers", &issuers)

	if err != nil {
		return nil, err
	}

	return issuers.Data, nil
}

func (a *IssuerApi) Get(issuerId string) (*Issuer, error) {
	var issuer Issuer
	err := a.c.Get("issuers/"+issuerId, &issuer)
	if err != nil {
		return nil, err
	}

	return &issuer, nil
}
