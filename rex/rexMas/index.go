package rexMas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	MasService struct {
		BaseService
		CaptchaConfigService CaptchaConfigService
		SmsConfigService     SmsConfigService
		EmsConfigService     EmsConfigService
	}
)

func NewMasService(SdkCtx *sdkCtx.SdkCtx) MasService {
	return MasService{
		BaseService:          NewBaseService(SdkCtx),
		CaptchaConfigService: NewCaptchaConfigService(SdkCtx),
		SmsConfigService:     NewSmsConfigService(SdkCtx),
		EmsConfigService:     NewEmsConfigService(SdkCtx),
	}
}
