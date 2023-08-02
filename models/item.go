package models

type Item struct {
	AssetID                 string                  `json:"AssetId,omitempty"`
	SiteID                  string                  `json:"SiteId,omitempty"`
	ProductID               string                  `json:"ProductId,omitempty"`
	CreatedDate             string                  `json:"CreatedDate,omitempty"`
	ModifiedDate            string                  `json:"ModifiedDate,omitempty"`
	AssetTypeID             string                  `json:"AssetTypeId,omitempty"`
	AssetTypeName           string                  `json:"AssetTypeName,omitempty"`
	IsDeleted               string                  `json:"IsDeleted,omitempty"`
	IsTraining              string                  `json:"IsTraining,omitempty"`
	StatusTypeID            string                  `json:"StatusTypeId,omitempty"`
	Status                  Status                  `json:"Status,omitempty"`
	AssetTag                string                  `json:"AssetTag,omitempty"`
	SerialNumber            string                  `json:"SerialNumber,omitempty"`
	ExternalID              string                  `json:"ExternalId,omitempty"`
	Name                    string                  `json:"Name,omitempty"`
	CanOwnerManage          string                  `json:"CanOwnerManage,omitempty"`
	IsFavorite              string                  `json:"IsFavorite,omitempty"`
	ModelID                 string                  `json:"ModelId,omitempty"`
	Model                   Model                   `json:"Model,omitempty"`
	OwnerID                 string                  `json:"OwnerId,omitempty"`
	Owner                   Owner                   `json:"Owner,omitempty"`
	LocationID              string                  `json:"LocationId,omitempty"`
	Location                Location                `json:"Location,omitempty"`
	LocationDetails         string                  `json:"LocationDetails,omitempty"`
	LocationRoomID          string                  `json:"LocationRoomId,omitempty"`
	LocationRoom            LocationRoom            `json:"LocationRoom,omitempty"`
	Notes                   string                  `json:"Notes,omitempty"`
	HasOpenTickets          string                  `json:"HasOpenTickets,omitempty"`
	OpenTickets             string                  `json:"OpenTickets,omitempty"`
	UpdateCustomFields      string                  `json:"UpdateCustomFields,omitempty"`
	CustomFieldValues       []CustomFieldValues     `json:"CustomFieldValues,omitempty"`
	CanTakeOwnership        string                  `json:"CanTakeOwnership,omitempty"`
	PurchasedDate           PurchasedDate           `json:"PurchasedDate,omitempty"`
	DeployedDate            DeployedDate            `json:"DeployedDate,omitempty"`
	RetiredDate             RetiredDate             `json:"RetiredDate,omitempty"`
	PurchasePrice           string                  `json:"PurchasePrice,omitempty"`
	PurchasePoNumber        string                  `json:"PurchasePoNumber,omitempty"`
	WarrantyExpirationDate  WarrantyExpirationDate  `json:"WarrantyExpirationDate,omitempty"`
	WarrantyInfo            string                  `json:"WarrantyInfo,omitempty"`
	LastInventoryDate       LastInventoryDate       `json:"LastInventoryDate,omitempty"`
	InvoiceNumber           string                  `json:"InvoiceNumber,omitempty"`
	Vendor                  string                  `json:"Vendor,omitempty"`
	InsuranceExpirationDate InsuranceExpirationDate `json:"InsuranceExpirationDate,omitempty"`
	InsuranceInfo           string                  `json:"InsuranceInfo,omitempty"`
	FundingSourceID         string                  `json:"FundingSourceId,omitempty"`
	FundingSource           FundingSource           `json:"FundingSource,omitempty"`
	IsReadOnly              string                  `json:"IsReadOnly,omitempty"`
	SpareGroupID            string                  `json:"SpareGroupId,omitempty"`
	SpareGroupName          string                  `json:"SpareGroupName,omitempty"`
	SparePoolID             string                  `json:"SparePoolId,omitempty"`
	SparePoolName           string                  `json:"SparePoolName,omitempty"`
	StorageUnitNumber       string                  `json:"StorageUnitNumber,omitempty"`
	StorageSlotNumber       string                  `json:"StorageSlotNumber,omitempty"`
	StorageLocationID       string                  `json:"StorageLocationId,omitempty"`
	StorageLocationName     string                  `json:"StorageLocationName,omitempty"`
	StorageLocation         StorageLocation         `json:"StorageLocation,omitempty"`
	DataMappings            DataMappings            `json:"DataMappings,omitempty"`
}

type Status struct {
	ProductID         string `json:"ProductId,omitempty"`
	AssetStatusTypeID string `json:"AssetStatusTypeId,omitempty"`
	Name              string `json:"Name,omitempty"`
	IsRetired         string `json:"IsRetired,omitempty"`
}

type Owner struct {
	UserID         string `json:"UserId,omitempty"`
	SiteID         string `json:"SiteId,omitempty"`
	Name           string `json:"Name,omitempty"`
	FirstName      string `json:"FirstName,omitempty"`
	LastName       string `json:"LastName,omitempty"`
	LocationID     string `json:"LocationId,omitempty"`
	PhotoID        string `json:"PhotoId,omitempty"`
	RoleID         string `json:"RoleId,omitempty"`
	SchoolIDNumber string `json:"SchoolIdNumber,omitempty"`
}

type CustomFieldValues struct {
	AssetID string `json:"AssetId,omitempty"`
}

type FundingSource struct {
	SiteID             string `json:"SiteId,omitempty"`
	Scope              string `json:"Scope,omitempty"`
	ProductID          string `json:"ProductId,omitempty"`
	AssetFundingTypeID string `json:"AssetFundingTypeId,omitempty"`
	Name               string `json:"Name,omitempty"`
}
