package models

import "time"

type User struct {
	UserID                                string              `json:"UserId,omitempty"`
	IsDeleted                             bool                `json:"IsDeleted,omitempty"`
	SiteID                                string              `json:"SiteId,omitempty"`
	Site                                  Site                `json:"Site,omitempty"`
	ProductID                             string              `json:"ProductId,omitempty"`
	CreatedDate                           string              `json:"CreatedDate,omitempty"`
	ModifiedDate                          string              `json:"ModifiedDate,omitempty"`
	LocationID                            string              `json:"LocationId,omitempty"`
	LocationName                          string              `json:"LocationName,omitempty"`
	Location                              Location            `json:"Location,omitempty"`
	IsActive                              bool                `json:"IsActive,omitempty"`
	IsOnline                              bool                `json:"IsOnline,omitempty"`
	IsOnlineLastUpdated                   time.Time           `json:"IsOnlineLastUpdated,omitempty"`
	Name                                  string              `json:"Name,omitempty"`
	FirstName                             string              `json:"FirstName,omitempty"`
	LastName                              string              `json:"LastName,omitempty"`
	Email                                 string              `json:"Email,omitempty"`
	Username                              string              `json:"Username,omitempty"`
	SchoolIDNumber                        string              `json:"SchoolIdNumber,omitempty"`
	Grade                                 string              `json:"Grade,omitempty"`
	ExternalID                            string              `json:"ExternalId,omitempty"`
	RoleID                                string              `json:"RoleId,omitempty"`
	Role                                  Role                `json:"Role,omitempty"`
	Department                            Department          `json:"Department,omitempty"`
	EmploymentStatus                      EmploymentStatus    `json:"EmploymentStatus,omitempty"`
	AuthenticatedBy                       string              `json:"AuthenticatedBy,omitempty"`
	AccountSetupProgress                  int                 `json:"AccountSetupProgress,omitempty"`
	IsEmailVerified                       bool                `json:"IsEmailVerified,omitempty"`
	IsWelcomeEmailSent                    bool                `json:"IsWelcomeEmailSent,omitempty"`
	PreventProviderUpdates                bool                `json:"PreventProviderUpdates,omitempty"`
	IsOutOfOffice                         bool                `json:"IsOutOfOffice,omitempty"`
	Portal                                int                 `json:"Portal,omitempty"`
	Options                               Options             `json:"Options,omitempty"`
	DataMappings                          DataMappings        `json:"DataMappings,omitempty"`
	CustomFieldValues                     []CustomFieldValues `json:"CustomFieldValues,omitempty"`
	IsExternalUser                        bool                `json:"IsExternalUser,omitempty"`
	ShouldOnlySeeExternallyAvailableRooms bool                `json:"ShouldOnlySeeExternallyAvailableRooms,omitempty"`
}

type Role struct {
	RoleID     string `json:"RoleId,omitempty"`
	Name       string `json:"Name,omitempty"`
	Portal     int    `json:"Portal,omitempty"`
	Visibility int    `json:"Visibility,omitempty"`
	CategoryID string `json:"CategoryId,omitempty"`
	Users      int    `json:"Users,omitempty"`
	Scope      string `json:"Scope,omitempty"`
}

type Department struct {
	DepartmentID string `json:"DepartmentId,omitempty"`
	SiteID       string `json:"SiteId,omitempty"`
	CreatedDate  string `json:"CreatedDate,omitempty"`
	ModifiedDate string `json:"ModifiedDate,omitempty"`
}

type EmploymentStatus struct {
	EmploymentStatusID string `json:"EmploymentStatusId,omitempty"`
	SiteID             string `json:"SiteId,omitempty"`
	CreatedDate        string `json:"CreatedDate,omitempty"`
	ModifiedDate       string `json:"ModifiedDate,omitempty"`
}

type Notifications struct {
	TicketCreated                      bool `json:"TicketCreated,omitempty"`
	TicketUpdated                      bool `json:"TicketUpdated,omitempty"`
	TicketAssignedToYou                bool `json:"TicketAssignedToYou,omitempty"`
	TicketAssignedToTeam               bool `json:"TicketAssignedToTeam,omitempty"`
	TicketDisableFollowerNotifications bool `json:"TicketDisableFollowerNotifications,omitempty"`
	AssetChanged                       bool `json:"AssetChanged,omitempty"`
}

type Options struct {
	Notifications Notifications `json:"Notifications,omitempty"`
}
