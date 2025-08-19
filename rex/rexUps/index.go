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

func NewUpsService(SdkCtx *sdkCtx.SdkCtx) UpsService {
	return UpsService{
		BaseService:      NewBaseService(SdkCtx),
		TagService:       NewTagService(SdkCtx),
		IndustryService:  NewIndustryService(SdkCtx),
		ShortLinkService: NewShortLinkService(SdkCtx),
	}
}
