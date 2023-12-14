package amaze

// Client -
type Client struct {
	*Config
	*Request
}

// NewClient -
func NewClient() (*Client, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	return &Client{
		Config:  config,
		Request: &Request{Config: config},
	}, nil
}
