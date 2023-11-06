// Every method in this file returns Asset models

package iiq

import (
	"fmt"
	"strings"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func (c *Client) AssetById(id string) (m.Response[m.ItemResponse[m.Asset]], error) {
	url := fmt.Sprintf("/assets/%s", id)
	resp, err := c.get(url, &m.ItemResponse[m.Asset]{})
	if err != nil {
		return m.Response[m.ItemResponse[m.Asset]]{}, err
	}

	model := convertBody(resp, m.ItemResponse[m.Asset]{})
	return model, err
}

func (c *Client) AssetsByAssetTag(params Parameters, assetTags ...string) (m.Response[m.ItemsResponse[m.Asset]], error) {
	query := params.encode()
	joined := strings.Join(assetTags, "|")
	url := fmt.Sprintf("/assets/assettag/search/%s?%s", joined, query)
	resp, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	if err != nil {
		return m.Response[m.ItemsResponse[m.Asset]]{}, err
	}

	model := convertBody(resp, m.ItemsResponse[m.Asset]{})
	return model, err
}

func (c *Client) AssetsBySerialNumber(params Parameters, serialNumbers ...string) (m.Response[m.ItemsResponse[m.Asset]], error) {
	query := params.encode()
	joined := strings.Join(serialNumbers, "|")
	url := fmt.Sprintf("/assets/serial/search/%s?%s", joined, query)
	resp, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	if err != nil {
		return m.Response[m.ItemsResponse[m.Asset]]{}, err
	}

	model := convertBody(resp, m.ItemsResponse[m.Asset]{})
	return model, err
}

func (c *Client) LinkedAssets(params Parameters, id string) (m.Response[m.ItemsResponse[m.Asset]], error) {
	query := params.encode()
	url := fmt.Sprintf("/assets/linked/to/%s?%s", id, query)
	resp, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	if err != nil {
		return m.Response[m.ItemsResponse[m.Asset]]{}, err
	}

	model := convertBody(resp, m.ItemsResponse[m.Asset]{})
	return model, err
}

func (c *Client) AllAssets(params Parameters, data interface{}) (m.Response[m.ItemsResponse[m.Asset]], error) {
	query := params.encode()
	url := fmt.Sprintf("/assets/?%s", query)
	resp, err := c.post(url, data, &m.ItemsResponse[m.Asset]{})
	if err != nil {
		return m.Response[m.ItemsResponse[m.Asset]]{}, err
	}

	model := convertBody(resp, m.ItemsResponse[m.Asset]{})
	return model, err
}
