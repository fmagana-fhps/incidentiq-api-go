package models

type DataMappings struct {
	AssetTag               AssetTagMappings               `json:"AssetTag,omitempty"`
	SerialNumber           SerialNumberMappings           `json:"SerialNumber,omitempty"`
	LocationName           LocationNameMappings           `json:"LocationName,omitempty"`
	LocationDetails        LocationDetailsMappings        `json:"LocationDetails,omitempty"`
	Room                   RoomMappings                   `json:"Room,omitempty"`
	Model                  ModelMappings                  `json:"Model,omitempty"`
	Owner                  OwnerMappings                  `json:"Owner,omitempty"`
	FundingSource          FundingSourceMappings          `json:"FundingSource,omitempty"`
	PurchasedDate          PurchasedDateMappings          `json:"PurchasedDate,omitempty"`
	PurchasedPrice         PurchasedPriceMappings         `json:"PurchasedPrice,omitempty"`
	PONumber               PONumberMappings               `json:"PONumber,omitempty"`
	DeployedDate           DeployedDateMappings           `json:"DeployedDate,omitempty"`
	WarrantyExpirationDate WarrantyExpirationDateMappings `json:"WarrantyExpirationDate,omitempty"`
	WarrantyInfo           WarrantyInfoMappings           `json:"WarrantyInfo,omitempty"`
	RetiredDate            RetiredDateMappings            `json:"RetiredDate,omitempty"`
	LastInventoryDate      LastInventoryDateMappings      `json:"LastInventoryDate,omitempty"`
	Notes                  NotesMappings                  `json:"Notes,omitempty"`
	Status                 StatusMappings                 `json:"Status,omitempty"`
	Lookups                []LookupsMappings              `json:"Lookups,omitempty"`
}

type AssetTagMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type SerialNumberMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type LocationNameMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type LocationDetailsMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type RoomMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type ModelMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type OwnerMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type FundingSourceMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type PurchasedDateMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type PurchasedPriceMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type PONumberMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type DeployedDateMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type WarrantyExpirationDateMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type WarrantyInfoMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type RetiredDateMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type LastInventoryDateMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type NotesMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type StatusMappings struct {
	AppID string `json:"AppId,omitempty"`
}
type LookupsMappings struct {
	Key   string `json:"Key,omitempty"`
	AppID string `json:"AppId,omitempty"`
	Value string `json:"Value,omitempty"`
}
