package models

type Location struct {
	LocationID     string       `json:"LocationId,omitempty"`
	SiteID         string       `json:"SiteId,omitempty"`
	Name           string       `json:"Name,omitempty"`
	Abbreviation   string       `json:"Abbreviation,omitempty"`
	CreatedDate    string       `json:"CreatedDate,omitempty"`
	ModifiedDate   string       `json:"ModifiedDate,omitempty"`
	AddressID      string       `json:"AddressId,omitempty"`
	Address        Address      `json:"Address,omitempty"`
	LocationTypeID string       `json:"LocationTypeId,omitempty"`
	LocationType   LocationType `json:"LocationType,omitempty"`
	LogoID         string       `json:"LogoId,omitempty"`
	FloorplanID    string       `json:"FloorplanId,omitempty"`
}

type Address struct {
	AddressID    string `json:"AddressId,omitempty"`
	SiteID       string `json:"SiteId,omitempty"`
	CreatedDate  string `json:"CreatedDate,omitempty"`
	ModifiedDate string `json:"ModifiedDate,omitempty"`
	Attn         string `json:"Attn,omitempty"`
	Street1      string `json:"Street1,omitempty"`
	Street2      string `json:"Street2,omitempty"`
	City         string `json:"City,omitempty"`
	State        string `json:"State,omitempty"`
	Zip          string `json:"Zip,omitempty"`
	Country      string `json:"Country,omitempty"`
}

type LocationType struct {
	LocationTypeID string `json:"LocationTypeId,omitempty"`
	Name           string `json:"Name,omitempty"`
}

type LocationRoom struct {
	LocationRoomID string `json:"LocationRoomId,omitempty"`
	SiteID         string `json:"SiteId,omitempty"`
	LocationID     string `json:"LocationId,omitempty"`
	Name           string `json:"Name,omitempty"`
	LocationName   string `json:"LocationName,omitempty"`
}

type StorageLocation struct {
	LocationID     string       `json:"LocationId,omitempty"`
	SiteID         string       `json:"SiteId,omitempty"`
	Name           string       `json:"Name,omitempty"`
	Abbreviation   string       `json:"Abbreviation,omitempty"`
	CreatedDate    string       `json:"CreatedDate,omitempty"`
	ModifiedDate   string       `json:"ModifiedDate,omitempty"`
	AddressID      string       `json:"AddressId,omitempty"`
	Address        Address      `json:"Address,omitempty"`
	LocationTypeID string       `json:"LocationTypeId,omitempty"`
	LocationType   LocationType `json:"LocationType,omitempty"`
	LogoID         string       `json:"LogoId,omitempty"`
	FloorplanID    string       `json:"FloorplanId,omitempty"`
}
