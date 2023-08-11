// Every method in this file returns Asset models

package iiq

import (
	"fmt"
	"strings"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

// {
// 	"OnlyShowDeleted": false,
// 	"Filters": [
// 	  {
// 		"Facet": "AssetType",
// 		"Id": "2a1561e5-34ff-4fcf-87de-2a146f0e1c01"
// 	  }
// 	],
// 	"FilterByViewPermission": true
// }

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

func (c *Client) AssetsBySerialNumber(serialNumber ...string) ([]m.Asset, error) {
	joined := strings.Join(serialNumber, "|")
	url := fmt.Sprintf("/assets/serial/search/%s", joined)
	model, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	return model.(*m.ItemsResponse[m.Asset]).Items, err
}

func (c *Client) GetLinkedAssets(id string) ([]m.Asset, error) {
	url := fmt.Sprintf("/assets/linked/to/%s", id)
	model, err := c.get(url, &m.ItemsResponse[m.Asset]{})
	return model.(*m.ItemsResponse[m.Asset]).Items, err
}
