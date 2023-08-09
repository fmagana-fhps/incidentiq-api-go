package iiq_test

import (
	"os"
	"testing"

	iiq "github.com/fmagana-fhps/incidentiq-api-go"
	"github.com/joho/godotenv"
)

var client *iiq.Client

func init() {
	godotenv.Load("../.env.local")
	client, _ = iiq.NewClient(&iiq.Options{
		SiteId: os.Getenv("SITEID"),
		Token:  os.Getenv("TOKEN"),
		Domain: os.Getenv("DOMAIN"),
	})
}

func TestAssetById(t *testing.T) {
	id := "4d2080b4-51dc-4b06-807b-21051c8e8fe7"
	result, err := client.AssetById(id)
	if err != nil {
		t.Error(err)
	}

	if result.AssetID != id {
		t.Errorf("AssetID = %s; expected %s", result.AssetID, id)
	}
}

func TestAssetByAssetTag(t *testing.T) {
	assetTags := []string{"82888", "82889"}
	result, err := client.AssetsByAssetTag(assetTags...)
	if err != nil {
		t.Error(err)
	}
	for i := range result {
		if result[i].AssetTag != assetTags[i] {
			t.Errorf("AssetTag = %s; expected %s", result[i].AssetTag, assetTags[i])
		}
	}
}

func TestAssetBySerialNumber(t *testing.T) {
	serialNumbers := []string{"476NTQ3", "9MFPTQ3"}
	result, err := client.AssetsBySerialNumber(serialNumbers...)
	if err != nil {
		t.Error(err)
	}
	for i := range result {
		if result[i].SerialNumber != serialNumbers[i] {
			t.Errorf("SerialNumber = %s; expected %s", result[i].SerialNumber, serialNumbers[i])
		}
	}
}
