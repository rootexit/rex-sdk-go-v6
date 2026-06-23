package rexSas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	SasService struct {
		BaseService
		BucketConfigService BucketConfigService
	}
)

func NewSasService(SdkCtx *sdkCtx.SdkCtx) SasService {
	return SasService{
		BaseService:         NewBaseService(SdkCtx),
		BucketConfigService: NewBucketConfigService(SdkCtx),
	}
}
