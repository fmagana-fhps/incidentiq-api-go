package iiq

import (
	"fmt"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

type AssetLinks []struct {
	ParentAssetID string `json:"ParentAssetId"`
	ChildAssetID  string `json:"ChildAssetId"`
}

func createBody(parentId string, childIds ...string) AssetLinks {
	links := AssetLinks{}

	for _, child := range childIds {
		s := struct {
			ParentAssetID string `json:"ParentAssetId"`
			ChildAssetID  string `json:"ChildAssetId"`
		}{parentId, child}

		links = append(links, s)
	}

	return links
}

func (c *Client) LinkAssets(parentId string, childIds ...string) (string, error) {
	url := fmt.Sprintf("/assets/linked/to/%s", parentId)
	res, err := c.post(url, createBody(parentId, childIds...), &m.BaseResponse{})
	return res.(m.BaseResponse).Message, err
}
