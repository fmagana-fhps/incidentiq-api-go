package iiq

import (
	"fmt"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func (c *Client) VerificationById(id string) (m.Response[m.ItemResponse[m.AssetVerification]], error) {
	url := fmt.Sprintf("/assets/verifications/%s", id)
	resp, err := c.get(url, m.ItemResponse[m.AssetVerification]{})
	if err != nil {
		return m.Response[m.ItemResponse[m.AssetVerification]]{}, err
	}

	model := convertBody(resp, m.ItemResponse[m.AssetVerification]{})
	return model, err
}

func (c *Client) VerificationsByAssetId(params Parameters, assetId string) (m.Response[m.ItemsResponse[m.AssetVerification]], error) {
	query := params.encode()
	url := fmt.Sprintf("/assets/%s/verifications/%s", assetId, query)
	resp, err := c.get(url, m.ItemsResponse[m.AssetVerification]{})
	if err != nil {
		return m.Response[m.ItemsResponse[m.AssetVerification]]{}, err
	}

	model := convertBody(resp, m.ItemsResponse[m.AssetVerification]{})
	return model, err
}

func (c *Client) NewVerification(assetId string) (m.Response[m.ItemResponse[m.AssetVerification]], error) {
	url := fmt.Sprintf("/assets/%s/verifications/new", assetId)
	resp, err := c.post(url, m.AssetVerification{}, m.AssetVerificationResponse{})
	if err != nil {
		return m.Response[m.ItemResponse[m.AssetVerification]]{}, err
	}

	model := convertBody(resp, m.ItemResponse[m.AssetVerification]{})
	return model, err
}
