package iiq

import (
	"fmt"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func (c *Client) ModelById(id string) (m.Model, error) {
	url := fmt.Sprintf("/models/%s", id)
	model, err := c.get(url, &m.ItemResponse[m.Model]{})
	return model.(*m.ItemResponse[m.Model]).Item, err
}
