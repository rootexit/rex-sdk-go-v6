package rexUps

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	UpsService struct {
		UpsBaseService
		UpsTagService       UpsTagService
		UpsIndustryService  UpsIndustryService
		UpsShortLinkService UpsShortLinkService
	}
)

func NewUpsService(rexCtx *rexCtx.QxCtx) UpsService {
	return UpsService{
		UpsBaseService:      NewUpsBaseService(rexCtx),
		UpsTagService:       NewUpsTagService(rexCtx),
		UpsIndustryService:  NewUpsIndustryService(rexCtx),
		UpsShortLinkService: NewUpsShortLinkService(rexCtx),
	}
}
