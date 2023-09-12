package models

type AssetVerification struct {
	AssetVerificationID     string `json:"AssetVerificationId,omitempty"`
	AssetID                 string `json:"AssetId,omitempty"`
	CreatedDate             string `json:"CreatedDate,omitempty"`
	VerifiedByUserID        string `json:"VerifiedByUserId,omitempty"`
	AssetVerificationTypeID string `json:"AssetVerificationTypeId,omitempty"`
	IsSuccessful            string `json:"IsSuccessful,omitempty"`
	Comments                string `json:"Comments,omitempty"`
	Latitude                string `json:"Latitude,omitempty"`
	Longitude               string `json:"Longitude,omitempty"`
	LocationID              string `json:"LocationId,omitempty"`
}
