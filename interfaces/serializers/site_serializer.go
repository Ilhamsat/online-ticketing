package serializers

import "online-ticketing/domain/entities"

type SiteDTO struct {
	ID           int    `json:"id"`
	SiteName     string `json:"site_name"`
	SiteDesc     string `json:"site_desc"`
	SiteImage    string `json:"site_image"`
	SiteLatitude string `json:"site_latitude"`
	SiteLongitude string `json:"site_longitude"`
}

func SerializeEvent(event *entities.Site) *SiteDTO {
	return &SiteDTO{
		ID:           event.ID,
		SiteName:     event.SiteName,
		SiteDesc:     event.SiteDesc,
		SiteImage:    event.SiteImage,
		SiteLatitude: event.SiteLatitude,
		SiteLongitude: event.SiteLongitude,
	}
}

func SerializeEvents(events []*entities.Site) []*SiteDTO {
	data := make([]*SiteDTO, 0)
	for _, event := range events {
		data = append(data, SerializeEvent(event))
	}

	return data
}
