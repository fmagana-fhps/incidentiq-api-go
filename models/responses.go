package models

import (
	"net/http"
	"time"
)

type Response[T any] struct {
	StatusCode int
	Header     http.Header
	Body       T
}

type BodyResponse struct {
	UserToken     string     `json:"UserToken"`
	RequestDate   time.Time  `json:"RequestDate"`
	ExecutionTime float64    `json:"ExecutionTime"`
	StatusCode    int        `json:"StatusCode"`
	ServerName    string     `json:"ServerName"`
	ProcessID     int        `json:"ProcessId"`
	OperationID   string     `json:"OperationId"`
	Properties    Properties `json:"Properties,omitempty"`
	Uid           string     `json:"Uid,omitempty"`
	Id            int        `json:"Id,omitempty"`
}

type ItemResponse[T any] struct {
	BodyResponse
	Item T `json:"Item"`
}

type ItemsResponse[T any] struct {
	BodyResponse
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
	BodyResponse
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
