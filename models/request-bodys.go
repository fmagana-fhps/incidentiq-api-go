package models

type FilterBody struct {
	OnlyShowDeleted        bool     `json:"OnlyShowDeleted,omitempty"`
	Filters                []Filter `json:"Filters,omitempty"`
	FilterByViewPermission bool     `json:"FilterByViewPermission,omitempty"`
}

type Filter struct {
	Facet string `json:"Facet,omitempty"`
	ID    string `json:"Id,omitempty"`
}

type AssetLinks struct {
	ChildAssetID string `json:"ChildAssetId"`
}
