// Every method in this file returns Asset models

package iiq

import (
	"fmt"
	"strings"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func (c *Client) AssetById(id string) (m.Asset, error) {
	url := fmt.Sprintf("/assets/%s", id)
	model, err := c.get(url, &m.ItemResponse[m.Asset]{})
	return model.(*m.ItemResponse[m.Asset]).Item, err
}

func (c *Client) AssetsByAssetTag(params Parameters, assetTags ...string) ([]m.Asset, error) {
	query := params.encode()
	joined := strings.Join(assetTags, "|")
	url := fmt.Sprintf("/assets/assettag/search/%s?%s", joined, query)
	model, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	return model.(*m.ItemsResponse[m.Asset]).Items, err
}

func (c *Client) AssetsBySerialNumber(params Parameters, serialNumbers ...string) ([]m.Asset, error) {
	query := params.encode()
	joined := strings.Join(serialNumbers, "|")
	url := fmt.Sprintf("/assets/serial/search/%s?%s", joined, query)
	model, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	return model.(*m.ItemsResponse[m.Asset]).Items, err
}

func (c *Client) GetLinkedAssets(params Parameters, id string) ([]m.Asset, error) {
	query := params.encode()
	url := fmt.Sprintf("/assets/linked/to/%s?%s", id, query)
	model, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	return model.(*m.ItemsResponse[m.Asset]).Items, err
}

func (c *Client) GetAllAssets(params Parameters, data interface{}) ([]m.Asset, error) {
	query := params.encode()
	url := fmt.Sprintf("/assets/?%s", query)
	model, err := c.post(url, data, &m.ItemsResponse[m.Asset]{})
	return model.(*m.ItemsResponse[m.Asset]).Items, err
}
