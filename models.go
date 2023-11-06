package iiq

import (
	"fmt"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func (c *Client) ModelById(id string) (m.Response[m.ItemResponse[m.Model]], error) {
	url := fmt.Sprintf("/models/%s", id)
	resp, err := c.get(url, &m.ItemResponse[m.Model]{})
	if err != nil {
		return m.Response[m.ItemResponse[m.Model]]{}, err
	}

	model := convertBody(resp, m.ItemResponse[m.Model]{})
	return model, err
}
