package persistence

import (
	"context"
	"fmt"
	"online-ticketing/commons/utils"
	"online-ticketing/domain/entities"
	"online-ticketing/domain/repositories"

	"gorm.io/gorm"
)

type SitePersistence struct {
	dbConn *gorm.DB
}

func NewSitePersistence(dbConn *gorm.DB) repositories.SiteRepository {
	return &SitePersistence{dbConn: dbConn}
}

func (p *SitePersistence) WithContext(ctx context.Context) repositories.SiteRepository {
	return &SitePersistence{dbConn: p.dbConn.WithContext(ctx)}
}

func (p *SitePersistence) Store(data *entities.Site) {
	if err := p.dbConn.Create(data).Error; err != nil {
		panic(utils.InvariantError("site", err))
	}
}

func (p *SitePersistence) GetList() []*entities.Site {
	var sites []*entities.Site
	if err := p.dbConn.Find(&sites).Error; err != nil {
		fmt.Println("ERR SITE")
		panic(utils.InvariantError("site", err))
	}

	return sites
}

func (p *SitePersistence) GetByColumn(data *entities.Site) *entities.Site {
	var site entities.Site
	if err := p.dbConn.Where(data).First(&site).Error; err != nil {
		panic(utils.InvariantError("site", err))
	}

	return &site
}
