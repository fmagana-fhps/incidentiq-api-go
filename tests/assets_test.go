package iiq_test

import (
	"os"
	"testing"

	iiq "github.com/fmagana-fhps/incidentiq-api-go"
	m "github.com/fmagana-fhps/incidentiq-api-go/models"
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
	assetTags := []string{"41708", "43023", "43914", "43921", "43922", "43923",
		"44259", "45014", "45017", "45020", "45027", "45037", "45075", "45408",
		"45861", "82880", "82881", "82882", "82883", "82884", "82885", "82886",
		"82887", "82888", "82889",
	}

	params := iiq.Parameters{PageSize: len(assetTags)}
	result, err := client.AssetsByAssetTag(params, assetTags...)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(assetTags); i++ {
		if result[i].AssetTag != assetTags[i] {
			t.Errorf("AssetTag = %s; expected %s", result[i].AssetTag, assetTags[i])
		}
	}
}

func TestAssetBySerialNumber(t *testing.T) {
	serialNumbers := []string{"476NTQ3", "9MFPTQ3"}

	params := iiq.Parameters{PageSize: len(serialNumbers)}
	result, err := client.AssetsBySerialNumber(params, serialNumbers...)
	if err != nil {
		t.Error(err)
	}
	for i := range result {
		if result[i].SerialNumber != serialNumbers[i] {
			t.Errorf("SerialNumber = %s; expected %s", result[i].SerialNumber, serialNumbers[i])
		}
	}
}

func TestAllAssets(t *testing.T) {
	params := iiq.Parameters{
		Filter:  "CategoryId eq '[36c6be98-4f39-e611-bf3b-005056bb000e]'",
		OrderBy: "AssetTag",
	}

	data := m.FilterBody{
		OnlyShowDeleted:        false,
		FilterByViewPermission: true,
		Filters: []m.Filter{
			{
				Facet: "Manufacturer",
				ID:    "518000c0-4dff-e511-a789-005056bb000e",
			},
			{
				Facet: "Location",
				ID:    "ca2c34ac-7f5f-481b-b563-a53f90165444",
			},
		},
	}

	result, err := client.GetAllAssets(params, data)
	if err != nil {
		t.Error(err)
	}

	if result[0].AssetTag != "200000" {
		t.Errorf("Length of returned assets = %s, expected %s", result[0].AssetTag, "20000")
	}
}
