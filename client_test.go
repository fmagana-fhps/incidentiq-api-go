package iiq

import (
	"testing"

	m "github.com/fmagana-fhps/incidentiq-api-go/models"
)

func TestClientFail(t *testing.T) {
	// model, _ := do[models.ItemResponse[models.Asset]](Client{"fhps", "token", "token"},
	// 	Options[models.ItemResponse[models.Asset]]{Method: "GET", Endpoint: ""})

	client, _ := NewClient(&Options{
		Domain: "fhps",
		Token:  "token",
	})

	model, _ := client.do("GET", "/", nil, &m.ItemResponse[m.Asset]{})

	if model.StatusCode != 500 {
		t.Errorf("check = %d; want 500", model.StatusCode)
	}
}
