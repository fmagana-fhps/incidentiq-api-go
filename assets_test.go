package iiq_test

import (
	"os"
	"testing"

	iiq "github.com/fmagana-fhps/incidentiq-api-go"
	"github.com/joho/godotenv"
)

var SiteModel iiq.Site

func init() {
	godotenv.Load(".env.local")
	SiteModel = iiq.Site{
		Id:     os.Getenv("SITEID"),
		Token:  os.Getenv("TOKEN"),
		Domain: os.Getenv("DOMAIN"),
	}
}

func TestAssetById(t *testing.T) {
	id := "4d2080b4-51dc-4b06-807b-21051c8e8fe7"
	result, err := iiq.AssetById(SiteModel, id)
	if err != nil {
		t.Error(err)
	}

	if result.AssetID != id {
		t.Errorf("AssetID = %s; want %s", result.AssetID, id)
	}
}

func TestAssetByAssetTag(t *testing.T) {
	assetTags := []string{"82888", "82889"}
	result, err := iiq.AssetsByAssetTag(SiteModel, assetTags...)
	if err != nil {
		t.Error(err)
	}
	for i := range result {
		if result[i].AssetTag != assetTags[i] {
			t.Errorf("AssetTag = %s; want %s", result[i].AssetTag, assetTags[i])
		}
	}
}

func TestAssetBySerialNumber(t *testing.T) {
	serialNumbers := []string{"476NTQ3", "9MFPTQ3"}
	result, err := iiq.AssetsBySerialNumber(SiteModel, serialNumbers...)
	if err != nil {
		t.Error(err)
	}
	for i := range result {
		if result[i].SerialNumber != serialNumbers[i] {
			t.Errorf("SerialNumber = %s; want %s", result[i].SerialNumber, serialNumbers[i])
		}
	}
}
