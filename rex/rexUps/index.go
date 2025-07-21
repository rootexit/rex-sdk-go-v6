package rexUps

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	UpsService struct {
		BaseService
		TagService       TagService
		IndustryService  IndustryService
		ShortLinkService ShortLinkService
	}
)

func NewUpsService(rexCtx *rexCtx.EngineCtx) UpsService {
	return UpsService{
		BaseService:      NewBaseService(rexCtx),
		TagService:       NewTagService(rexCtx),
		IndustryService:  NewIndustryService(rexCtx),
		ShortLinkService: NewShortLinkService(rexCtx),
	}
}
