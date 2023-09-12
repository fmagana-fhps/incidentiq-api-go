package models

import "time"

type BaseResponse struct {
	UserToken     string     `json:"UserToken"`
	RequestDate   time.Time  `json:"RequestDate"`
	ExecutionTime float64    `json:"ExecutionTime"`
	StatusCode    int        `json:"StatusCode"`
	ServerName    string     `json:"ServerName"`
	ProcessID     int        `json:"ProcessId"`
	OperationID   string     `json:"OperationId"`
	Properties    Properties `json:"Properties,omitempty"`
}

type ItemResponse[T any] struct {
	BaseResponse
	Item T `json:"Item"`
}

type ItemsResponse[T any] struct {
	BaseResponse
	ItemCount int      `json:"ItemCount"`
	Items     []T      `json:"Items"`
	Paging    Paging   `json:"Paging"`
	Metadata  Metadata `json:"Metadata,omitempty"`
}

type Paging struct {
	TotalRows int `json:"TotalRows"`
	PageCount int `json:"PageCount"`
	PageSize  int `json:"PageSize"`
	PageIndex int `json:"PageIndex"`
}

type LinkResponse struct {
	BaseResponse
	Message string `json:"Message"`
}

type Properties struct {
}

type Metadata struct {
	QueryEngineSource string `json:"QueryEngineSource,omitempty"`
}

type AssetVerificationResponse struct {
	AssetVerification AssetVerification `json:"Item"`
	Id                int               `json:"Id"`
	Uid               string            `json:"Uid"`
}
