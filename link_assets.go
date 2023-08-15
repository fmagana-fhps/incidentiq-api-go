package iiq

import (
	"fmt"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

type AssetLinks struct {
	ChildAssetID string `json:"ChildAssetId"`
}

func createBody(childIds []string) []AssetLinks {
	links := []AssetLinks{}

	for _, child := range childIds {
		s := AssetLinks{child}

		links = append(links, s)
	}

	return links
}

func (c *Client) LinkAssets(parentId string, childIds ...string) (string, error) {
	url := fmt.Sprintf("/assets/linked/to/%s", parentId)
	res, err := c.post(url, createBody(childIds), &m.LinkResponse{})
	return res.(*m.LinkResponse).Message, err
}

func (c *Client) UnlinkAssets(parentId string, childIds ...string) (string, error) {
	url := fmt.Sprintf("/assets/linked/to/%s", parentId)
	res, err := c.delete(url, createBody(childIds), &m.LinkResponse{})
	return res.(*m.LinkResponse).Message, err
}
