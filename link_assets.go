package iiq

import (
	"fmt"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func createBody(childIds []string) []m.AssetLinks {
	links := []m.AssetLinks{}

	for _, child := range childIds {
		s := m.AssetLinks{
			ChildAssetID: child,
		}

		links = append(links, s)
	}

	return links
}

func (c *Client) LinkAssets(parentId string, childIds ...string) (m.Response[m.LinkResponse], error) {
	url := fmt.Sprintf("/assets/linked/to/%s", parentId)
	resp, err := c.post(url, createBody(childIds), &m.LinkResponse{})
	if err != nil {
		return m.Response[m.LinkResponse]{}, err
	}

	model := convertBody(resp, m.LinkResponse{})
	return model, err
}

func (c *Client) UnlinkAssets(parentId string, childIds ...string) (m.Response[m.LinkResponse], error) {
	url := fmt.Sprintf("/assets/linked/to/%s", parentId)
	resp, err := c.delete(url, createBody(childIds), &m.LinkResponse{})
	if err != nil {
		return m.Response[m.LinkResponse]{}, err
	}

	model := convertBody(resp, m.LinkResponse{})
	return model, err
}
