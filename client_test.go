package iiq

import (
	"testing"

	"github.com/fmagana-fhps/incidentiq-api-go/models"
)

func TestClientFail(t *testing.T) {
	model, _ := do[models.ItemResponse[models.Asset]](Site{"fhps", "token", "token"},
		Options[models.ItemResponse[models.Asset]]{Method: "GET", Endpoint: ""})

	if model.StatusCode != 500 {
		t.Errorf("check = %d; want 500", model.StatusCode)
	}
}
