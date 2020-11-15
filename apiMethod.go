package etsy

type ApiMethod struct {
	Name        string
	Description string
	Uri         string
	Params      map[string]string
	Defaults    map[string]interface{}
	HttpMethod  string `json:"http_method"`
	Type        string
}

func (client *Client) GetMethodTable() []ApiMethod {
	url := client.BuildUrl("", nil, nil)

	m, err := client.Get(url, true)
	if err != nil {
		panic(err)
	}
	return m.([]ApiMethod)
}
