package rexSas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	SasService struct {
		BaseService
	}
)

func NewSasService(SdkCtx *sdkCtx.SdkCtx) SasService {
	return SasService{
		BaseService: NewBaseService(SdkCtx),
	}
}
