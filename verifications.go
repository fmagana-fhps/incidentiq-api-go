package iiq

import (
	"fmt"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func (c *Client) VerificationById(id string) (m.AssetVerification, error) {
	url := fmt.Sprintf("/assets/verifications/%s", id)
	model, err := c.get(url, m.ItemResponse[m.AssetVerification]{})
	return model.(*m.ItemResponse[m.AssetVerification]).Item, err
}

func (c *Client) VerificationsByAssetId(params Parameters, assetId string) ([]m.AssetVerification, error) {
	query := params.encode()
	url := fmt.Sprintf("/assets/%s/verifications/%s", assetId, query)
	model, err := c.get(url, m.ItemsResponse[m.AssetVerification]{})
	return model.(*m.ItemsResponse[m.AssetVerification]).Items, err
}

func (c *Client) NewVerification(assetId string) (m.AssetVerification, error) {
	url := fmt.Sprintf("/assets/%s/verifications/new", assetId)
	model, err := c.post(url, m.AssetVerification{}, m.AssetVerificationResponse{})
	return model.(*m.AssetVerificationResponse).AssetVerification, err
}
