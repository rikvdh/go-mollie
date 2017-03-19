package mollie

type MethodApi struct {
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

type MethodListWrapper struct {
	TotalCount int `json:"totalCount"`
	Offset     int
	Count      int
	Data       []Method
}

func newMethods(co *core) *MethodApi {
	return &MethodApi{c: co}
}

func (a *MethodApi) List() ([]Method, error) {
	var methods MethodListWrapper
	err := a.c.Get("methods", &methods)

	if err != nil {
		return nil, err
	}

	return methods.Data, nil
}

func (a *MethodApi) Get(methodID string) (*Method, error) {
	var method Method
	err := a.c.Get("methods/"+methodID, &method)
	if err != nil {
		return nil, err
	}

	return &method, nil
}
