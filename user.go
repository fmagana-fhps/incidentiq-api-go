package iiq

import (
	"fmt"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func (c *Client) UserById(id string) (m.Response[m.ItemResponse[m.User]], error) {
	url := fmt.Sprintf("/users/%s", id)
	resp, err := c.get(url, &m.ItemResponse[m.User]{})
	if err != nil {
		return m.Response[m.ItemResponse[m.User]]{}, err
	}

	model := convertBody(resp, m.ItemResponse[m.User]{})
	return model, err
}

func (c *Client) AllUsers(params Parameters) (m.Response[m.ItemsResponse[m.User]], error) {
	query := params.encode()
	url := fmt.Sprintf("/users/?%s", query)
	resp, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	if err != nil {
		return m.Response[m.ItemsResponse[m.User]]{}, err
	}

	model := convertBody(resp, m.ItemsResponse[m.User]{})
	return model, err
}
