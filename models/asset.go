package models

type Asset struct {
	AssetID                         string              `json:"AssetId,omitempty"`
	SiteID                          string              `json:"SiteId,omitempty"`
	Site                            Site                `json:"Site,omitempty"`
	ProductID                       string              `json:"ProductId,omitempty"`
	CreatedDate                     string              `json:"CreatedDate,omitempty"`
	ModifiedDate                    string              `json:"ModifiedDate,omitempty"`
	AssetTypeID                     string              `json:"AssetTypeId,omitempty"`
	AssetTypeName                   string              `json:"AssetTypeName,omitempty"`
	IsDeleted                       bool                `json:"IsDeleted,omitempty"`
	IsTraining                      bool                `json:"IsTraining,omitempty"`
	StatusTypeID                    string              `json:"StatusTypeId,omitempty"`
	Status                          Status              `json:"Status,omitempty"`
	AssetTag                        string              `json:"AssetTag,omitempty"`
	SerialNumber                    string              `json:"SerialNumber,omitempty"`
	ExternalID                      string              `json:"ExternalId,omitempty"`
	Name                            string              `json:"Name,omitempty"`
	CanOwnerManage                  bool                `json:"CanOwnerManage,omitempty"`
	IsFavorite                      bool                `json:"IsFavorite,omitempty"`
	ModelID                         string              `json:"ModelId,omitempty"`
	Model                           Model               `json:"Model,omitempty"`
	OwnerID                         string              `json:"OwnerId,omitempty"`
	Owner                           Owner               `json:"Owner,omitempty"`
	LocationID                      string              `json:"LocationId,omitempty"`
	Location                        Location            `json:"Location,omitempty"`
	LocationDetails                 string              `json:"LocationDetails,omitempty"`
	LocationRoomID                  string              `json:"LocationRoomId,omitempty"`
	LocationRoom                    LocationRoom        `json:"LocationRoom,omitempty"`
	Notes                           string              `json:"Notes,omitempty"`
	HasOpenTickets                  bool                `json:"HasOpenTickets,omitempty"`
	OpenTickets                     int                 `json:"OpenTickets,omitempty"`
	UpdateCustomFields              bool                `json:"UpdateCustomFields,omitempty"`
	CustomFieldValues               []CustomFieldValues `json:"CustomFieldValues,omitempty"`
	CanTakeOwnership                bool                `json:"CanTakeOwnership,omitempty"`
	PurchasedDate                   string              `json:"PurchasedDate,omitempty"`
	DeployedDate                    string              `json:"DeployedDate,omitempty"`
	RetiredDate                     string              `json:"RetiredDate,omitempty"`
	PurchasePrice                   float32             `json:"PurchasePrice,omitempty"`
	PurchasePoNumber                string              `json:"PurchasePoNumber,omitempty"`
	WarrantyExpirationDate          string              `json:"WarrantyExpirationDate,omitempty"`
	WarrantyInfo                    string              `json:"WarrantyInfo,omitempty"`
	LastInventoryDate               string              `json:"LastInventoryDate,omitempty"`
	InvoiceNumber                   string              `json:"InvoiceNumber,omitempty"`
	Vendor                          string              `json:"Vendor,omitempty"`
	InsuranceExpirationDate         string              `json:"InsuranceExpirationDate,omitempty"`
	InsuranceInfo                   string              `json:"InsuranceInfo,omitempty"`
	FundingSourceID                 string              `json:"FundingSourceId,omitempty"`
	FundingSource                   FundingSource       `json:"FundingSource,omitempty"`
	IsReadOnly                      bool                `json:"IsReadOnly,omitempty"`
	IsExternallyManaged             bool                `json:"IsExternallyManaged,omitempty"`
	SpareGroupID                    string              `json:"SpareGroupId,omitempty"`
	SpareGroupName                  string              `json:"SpareGroupName,omitempty"`
	SparePoolID                     string              `json:"SparePoolId,omitempty"`
	SparePoolName                   string              `json:"SparePoolName,omitempty"`
	StorageUnitNumber               string              `json:"StorageUnitNumber,omitempty"`
	StorageSlotNumber               int                 `json:"StorageSlotNumber,omitempty"`
	StorageLocationID               string              `json:"StorageLocationId,omitempty"`
	StorageLocationName             string              `json:"StorageLocationName,omitempty"`
	AssetAuditPolicyStatusName      string              `json:"AssetAuditPolicyStatusName,omitempty"`
	AssetAuditPolicyScheduleName    string              `json:"AssetAuditPolicyScheduleName,omitempty"`
	AssetAuditPolicyName            string              `json:"AssetAuditPolicyName,omitempty"`
	AssetAuditPolicyStatusSortOrder int                 `json:"AssetAuditPolicyStatusSortOrder,omitempty"`
	LastVerificationDateTime        string              `json:"LastVerificationDateTime,omitempty"`
	LastVerificationUserID          string              `json:"LastVerificationUserId,omitempty"`
	LastVerificationTypeID          string              `json:"LastVerificationTypeId,omitempty"`
	LastVerificationTypeName        string              `json:"LastVerificationTypeName,omitempty"`
	LastVerificationSuccessful      bool                `json:"LastVerificationSuccessful,omitempty"`
	LastVerificationComments        string              `json:"LastVerificationComments,omitempty"`
	LastVerificationLocationID      string              `json:"LastVerificationLocationId,omitempty"`
	LastVerificationLocationName    string              `json:"LastVerificationLocationName,omitempty"`
	StorageLocation                 StorageLocation     `json:"StorageLocation,omitempty"`
	DataMappings                    DataMappings        `json:"DataMappings,omitempty"`
}

type Site struct {
	SiteID                 string `json:"SiteId"`
	ProductID              string `json:"ProductId"`
	CreatedDate            string `json:"CreatedDate"`
	Name                   string `json:"Name"`
	OffsetInMinutesFromUtc int    `json:"OffsetInMinutesFromUtc"`
	WorkMinutesPerDay      int    `json:"WorkMinutesPerDay"`
	DbClusterIndex         int    `json:"DbClusterIndex"`
	StatusTypeID           string `json:"StatusTypeId"`
	WorkHoursStartUtc      string `json:"WorkHoursStartUtc"`
	WorkHoursEndUtc        string `json:"WorkHoursEndUtc"`
	SystemUserID           string `json:"SystemUserId"`
	AnonymousUserID        string `json:"AnonymousUserId"`
	IsSandbox              bool   `json:"IsSandbox"`
}

type Status struct {
	ProductID         string `json:"ProductId,omitempty"`
	AssetStatusTypeID string `json:"AssetStatusTypeId,omitempty"`
	Name              string `json:"Name,omitempty"`
	IsRetired         bool   `json:"IsRetired,omitempty"`
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
