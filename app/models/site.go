package models

type NewSiteRequest struct {
	SiteName     string `json:"site_name" validate:"required"`
	SiteDesc     string `json:"site_desc" validate:"required"`
	SiteImage    string `json:"site_image" validate:"required"`
	SiteLatitude string    `json:"site_latitude" validate:"required"`
	SiteLongitude string    `json:"site_longitude" validate:"required"`
}

type GetSiteRequest struct {
	ID int `form:"id"`
}
