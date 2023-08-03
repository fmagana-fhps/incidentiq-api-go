package models

import "time"

type ItemResponse[T any] struct {
	Item          T          `json:"Item,omitempty"`
	UserToken     string     `json:"UserToken,omitempty"`
	RequestDate   time.Time  `json:"RequestDate,omitempty"`
	ExecutionTime float64    `json:"ExecutionTime,omitempty"`
	StatusCode    int        `json:"StatusCode,omitempty"`
	ServerName    string     `json:"ServerName,omitempty"`
	ProcessID     int        `json:"ProcessId,omitempty"`
	OperationID   string     `json:"OperationId.omitempty"`
	Properties    Properties `json:"Properties,omitempty"`
}

type ItemsResponse[T any] struct {
	ItemCount     int        `json:"ItemCount,omitempty"`
	Items         []T        `json:"Items,omitempty"`
	Paging        Paging     `json:"Paging,omitempty"`
	UserToken     string     `json:"UserToken,omitempty"`
	RequestDate   time.Time  `json:"RequestDate,omitempty"`
	ExecutionTime float64    `json:"ExecutionTime,omitempty"`
	StatusCode    int        `json:"StatusCode,omitempty"`
	ServerName    string     `json:"ServerName,omitempty"`
	ProcessID     int        `json:"ProcessId,omitempty"`
	OperationID   string     `json:"OperationId.omitempty"`
	Properties    Properties `json:"Properties,omitempty"`
	Metadata      Metadata   `json:"Metadata,omitempty"`
}

type Paging struct {
	TotalRows int `json:"TotalRows,omitempty"`
	PageCount int `json:"PageCount,omitempty"`
	PageSize  int `json:"PageSize,omitempty"`
	PageIndex int `json:"PageIndex,omitempty"`
}
type Properties struct {
}
type Metadata struct {
	QueryEngineSource string `json:"QueryEngineSource,omitempty"`
}
