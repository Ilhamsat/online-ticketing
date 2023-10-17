package entities

import "time"

type Site struct {
	ID           int       `json:"id"`
	SiteName     string    `json:"site_name" validate:"required"`
	SiteDesc     string    `json:"site_desc" validate:"required"`
	SiteImage    string    `json:"site_image" validate:"required"`
	SiteLatitude string    `json:"site_latitude" validate:"required"`
	SiteLongitude string    `json:"site_longitude" validate:"required"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func (Site) TableName() string {
	return "site"
}
