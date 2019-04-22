package dominos

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	APIURL = "https://order.dominos.com/power/"
)

var orderTypes = map[string]string{"delivery": "delivery", "carryout": "carryout"}

type Option func(*Client)

type Client struct {
	client    *http.Client
	url       string
	orderType string
	storeID   string
}

// New creates new client
func New(orderType, storeID string, opts ...Option) *Client {
	return &Client{
		orderType: orderType,
		url:       APIURL,
		storeID:   storeID,
	}
}

func (c *Client) Do(params url.Values, requestType string) (*http.Response, error) {
	if c.orderType == "" {
		return nil, errors.New("must provide order type")
	}

	// setup URL for action
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}

	u.Path = strings.TrimSuffix(u.Path, "/") + "/"
	u.RawQuery = params.Encode()

	req, err := http.NewRequest(requestType, u.String(), nil)
	if err != nil {
		return nil, err
	}

	cl := c.client
	if cl == nil {
		cl = http.DefaultClient
	}

	// do
	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		return nil, fmt.Errorf("recieved status code %d", res.StatusCode)
	}
	return res, nil
}
