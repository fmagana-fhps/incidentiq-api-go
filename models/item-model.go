package models

type Model struct {
	ModelID                string       `json:"ModelId,omitempty"`
	SiteID                 string       `json:"SiteId,omitempty"`
	Scope                  string       `json:"Scope,omitempty"`
	Name                   string       `json:"Name,omitempty"`
	ModelName              string       `json:"ModelName,omitempty"`
	IsEnabled              string       `json:"IsEnabled,omitempty"`
	AssetTypeID            string       `json:"AssetTypeId,omitempty"`
	AssetTypeName          string       `json:"AssetTypeName,omitempty"`
	ManufacturerID         string       `json:"ManufacturerId,omitempty"`
	Manufacturer           Manufacturer `json:"Manufacturer,omitempty"`
	CategoryID             string       `json:"CategoryId,omitempty"`
	Category               Category     `json:"Category,omitempty"`
	CategoryNameWithParent string       `json:"CategoryNameWithParent,omitempty"`
	ImageID                string       `json:"ImageId,omitempty"`
	Icon                   string       `json:"Icon,omitempty"`
	CreatedDate            string       `json:"CreatedDate,omitempty"`
	ModifiedDate           string       `json:"ModifiedDate,omitempty"`
}

type Manufacturer struct {
	ManufacturerID   string `json:"ManufacturerId,omitempty"`
	Name             string `json:"Name,omitempty"`
	IsEnabled        string `json:"IsEnabled,omitempty"`
	LogoID           string `json:"LogoId,omitempty"`
	Icon             string `json:"Icon,omitempty"`
	TotalModels      string `json:"TotalModels,omitempty"`
	AvailableModels  string `json:"AvailableModels,omitempty"`
	IncludeAllModels string `json:"IncludeAllModels,omitempty"`
}

type Category struct {
	SiteID                         string                           `json:"SiteId,omitempty"`
	Scope                          string                           `json:"Scope,omitempty"`
	ProductID                      string                           `json:"ProductId,omitempty"`
	CategoryID                     string                           `json:"CategoryId,omitempty"`
	Level                          string                           `json:"Level,omitempty"`
	SortOrder                      string                           `json:"SortOrder,omitempty"`
	ParentCategoryID               string                           `json:"ParentCategoryId,omitempty"`
	ParentCategoryName             string                           `json:"ParentCategoryName,omitempty"`
	CategoryTypeID                 string                           `json:"CategoryTypeId,omitempty"`
	CategoryTypeName               string                           `json:"CategoryTypeName,omitempty"`
	Name                           string                           `json:"Name,omitempty"`
	Description                    string                           `json:"Description,omitempty"`
	ImageID                        string                           `json:"ImageId,omitempty"`
	Icon                           string                           `json:"Icon,omitempty"`
	ModelsCount                    string                           `json:"ModelsCount,omitempty"`
	CategoryIssues                 []string                         `json:"CategoryIssues,omitempty"`
	CategoryCustomFieldTypes       []CategoryCustomFieldTypes       `json:"CategoryCustomFieldTypes,omitempty"`
	TicketCategoryCustomFieldTypes []TicketCategoryCustomFieldTypes `json:"TicketCategoryCustomFieldTypes,omitempty"`
	NameWithParent                 string                           `json:"NameWithParent,omitempty"`
}

type CategoryCustomFieldTypes struct {
	SiteID                   string `json:"SiteId,omitempty"`
	ProductID                string `json:"ProductId,omitempty"`
	CustomFieldTypeID        string `json:"CustomFieldTypeId,omitempty"`
	Scope                    string `json:"Scope,omitempty"`
	Name                     string `json:"Name,omitempty"`
	OwnerAppID               string `json:"OwnerAppId,omitempty"`
	EditorType               string `json:"EditorType,omitempty"`
	Options                  string `json:"Options,omitempty"`
	HasSensitiveInformation  string `json:"HasSensitiveInformation,omitempty"`
	Configuration            string `json:"Configuration,omitempty"`
	IsRequired               string `json:"IsRequired,omitempty"`
	IsReadOnly               string `json:"IsReadOnly,omitempty"`
	IsHidden                 string `json:"IsHidden,omitempty"`
	MaxNumberOfInstances     string `json:"MaxNumberOfInstances,omitempty"`
	DefaultSiteRowVisibility string `json:"DefaultSiteRowVisibility,omitempty"`
	DisplayOrder             string `json:"DisplayOrder,omitempty"`
}

type TicketCategoryCustomFieldTypes struct {
	SiteID                   string `json:"SiteId,omitempty"`
	ProductID                string `json:"ProductId,omitempty"`
	CustomFieldTypeID        string `json:"CustomFieldTypeId,omitempty"`
	Scope                    string `json:"Scope,omitempty"`
	Name                     string `json:"Name,omitempty"`
	OwnerAppID               string `json:"OwnerAppId,omitempty"`
	EditorType               string `json:"EditorType,omitempty"`
	Options                  string `json:"Options,omitempty"`
	HasSensitiveInformation  string `json:"HasSensitiveInformation,omitempty"`
	Configuration            string `json:"Configuration,omitempty"`
	IsRequired               string `json:"IsRequired,omitempty"`
	IsReadOnly               string `json:"IsReadOnly,omitempty"`
	IsHidden                 string `json:"IsHidden,omitempty"`
	MaxNumberOfInstances     string `json:"MaxNumberOfInstances,omitempty"`
	DefaultSiteRowVisibility string `json:"DefaultSiteRowVisibility,omitempty"`
	DisplayOrder             string `json:"DisplayOrder,omitempty"`
}
