// Every method in this file returns Asset models

package iiq

import (
	"fmt"
	"strings"

	"github.com/fmagana-fhps/incidentiq-api-go/models"
)

// {
// 	"OnlyShowDeleted": false,
// 	"Filters": [
// 	  {
// 		"Facet": "AssetType",
// 		"Id": "2a1561e5-34ff-4fcf-87de-2a146f0e1c01"
// 	  }
// 	],
// 	"FilterByViewPermission": true
// }

// # Python code for get most pages
// 		url = "https://" + config.IIQ_INSTANCE + "/api/v1.0/assets/?$p=" + str(
// 		page) + "&$s=" + config.PAGE_SIZE + "&$d=Ascending&$o=AssetTag"
// # Cause an exception if for some reason the API returns nothing
// if response.json()['Paging']['PageSize'] <= 0:
// 	print("ERROR NO ELEMENTS WERE RETURNED")
// 	raise HTTPError("No elements were returned from a request")

// return response

// @staticmethod
// def get_num_pages():
// return Asset.get_data_request(0).json()['Paging']['PageCount']

// @classmethod
// def get_page(cls, page_number):
// return super().get_page(page_number)

func AssetById(site Site, id string) (models.Asset, error) {
	model, err := get[models.ItemResponse[models.Asset]](site, fmt.Sprintf("/assets/%s", id))
	return model.Item, err
}

func AssetsByAssetTag(site Site, assetTags ...string) ([]models.Asset, error) {
	joined := strings.Join(assetTags, "|")
	model, err := get[models.ItemsResponse[models.Asset]](site, fmt.Sprintf("/assets/assettag/search/%s", joined))
	return model.Items, err
}

func AssetsBySerialNumber(site Site, serialNumber ...string) ([]models.Asset, error) {
	joined := strings.Join(serialNumber, "|")
	model, err := get[models.ItemsResponse[models.Asset]](site, fmt.Sprintf("/assets/serial/search/%s", joined))
	return model.Items, err
}
