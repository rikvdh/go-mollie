package mollie

import "time"

type CustomerAPI struct {
	c *core
}

type Customer struct {
	Resource string `json:"resource"`
	ID string `json:"id"`
	Mode string `json:"mode"`
	Name string `json:"name"`
	Email string `json:"email"`
	Locale string `json:"locale"`
	Metadata string `json:"metadata"`
	RecentlyUsedMethods []string `json:"recentlyUsedMethods"`
	CreatedDatetime time.Time `json:"createdDatetime"`
}

func newCustomerAPI(co *core) *CustomerAPI {
	return &CustomerAPI{co}
}

func (cAPI *CustomerAPI) Get(id string) (*Customer, error) {
	customer := &Customer{}

	err := cAPI.c.Get("/customers" + id, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}