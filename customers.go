package mollie

import "time"

type CustomerAPI struct {
	c *core
}

type NewCustomerData struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Locale string `json:"locale,omitempty"` // One of de_DE en_US es_ES fr_FR nl_BE fr_BE nl_NL
	Metadata string `json:"metadata,omitempty"`
}

type Customer struct {
	Resource string `json:"resource"`
	ID string `json:"id"`
	Mode string `json:"mode"`
	Name string `json:"name"`
	Email string `json:"email"`
	Locale string `json:"locale"`
	Metadata string `json:"metadata,omitempty"`
	RecentlyUsedMethods []string `json:"recentlyUsedMethods"`
	CreatedDatetime time.Time `json:"createdDatetime"`
}

func newCustomers(co *core) *CustomerAPI {
	return &CustomerAPI{co}
}

func(api *CustomerAPI) New(data NewCustomerData) (*Customer, error) {
	customer := &Customer{}

	err := api.c.Post("customers", customer, data)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (api *CustomerAPI) Get(id string) (*Customer, error) {
	customer := &Customer{}

	err := api.c.Get("customers/" + id, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}