package iiq

import (
	"os"
	"testing"

	"github.com/fmagana-fhps/incidentiq-api-go/models"
	"github.com/joho/godotenv"
)

func TestClientResponses(t *testing.T) {
	err := godotenv.Load(".env.local")
	if err != nil {
		t.Error(err)
	}

	site := &IncidentIQ{
		SiteID: os.Getenv("SITEID"),
		Token:  os.Getenv("TOKEN"),
		Domain: os.Getenv("DOMAIN"),
	}

	model := models.ItemResponse{}
	site.doGet("/assets?$s=1", &model)

	if model.ItemCount != 1 {
		t.Errorf("'/assets?$s=1' = %d; want 1", model.ItemCount)
	}
}
