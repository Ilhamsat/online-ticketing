package repositories

import (
	"context"
	"online-ticketing/domain/entities"
)

type SiteRepository interface {
	WithContext(context.Context) SiteRepository

	Store(*entities.Site)
	GetList() []*entities.Site
	GetByColumn(*entities.Site) *entities.Site
}
