package services

import (
	"context"
	"online-ticketing/app/models"
	"online-ticketing/domain/entities"
	"online-ticketing/domain/repositories"
	"github.com/gin-gonic/gin"
)

type SiteService struct {
	ctx context.Context

	siteRepository repositories.SiteRepository
}

func NewSiteService(siteService repositories.SiteRepository) *SiteService {
	return &SiteService{siteRepository: siteService}
}

func (s *SiteService) WithContext(ctx *gin.Context) *SiteService {
	s.ctx = ctx.Request.Context()
	return s
}

func (s *SiteService) Store(data *models.NewSiteRequest) {
	newSite := &entities.Site{

		SiteName:  data.SiteName,
		SiteDesc:  data.SiteDesc,
		SiteImage: data.SiteImage,
		SiteLatitude: data.SiteLatitude,
		SiteLongitude: data.SiteLongitude,
	}

	s.siteRepository.WithContext(s.ctx).Store(newSite)
}

func (s *SiteService) GetSites(query *models.GetSiteRequest) []*entities.Site {
	if query.ID != 0 {
		var sites []*entities.Site

		user := s.siteRepository.WithContext(s.ctx).GetByColumn(&entities.Site{ID: query.ID})
		sites = append(sites, user)

		return sites
	} else {
		return s.siteRepository.WithContext(s.ctx).GetList()
	}
}
