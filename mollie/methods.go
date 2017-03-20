package mollie

type MethodAPI struct {
	c *core
}

type MethodAmount struct {
	Minimum float64 `json:",string"`
	Maximum float64 `json:",string"`
}

type MethodImage struct {
	Normal string
	Buffer string
}

type Method struct {
	ID          string
	Description string
	Amount      MethodAmount
	Image       MethodImage
	Resource    string
}

type methodListWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []Method
}

func newMethods(co *core) *MethodAPI {
	return &MethodAPI{c: co}
}

func (a *MethodAPI) List() ([]Method, error) {
	var methods methodListWrapper
	err := a.c.Get("methods", &methods)

	if err != nil {
		return nil, err
	}

	return methods.Data, nil
}

func (a *MethodAPI) Get(methodID string) (*Method, error) {
	var method Method
	err := a.c.Get("methods/"+methodID, &method)
	if err != nil {
		return nil, err
	}

	return &method, nil
}
