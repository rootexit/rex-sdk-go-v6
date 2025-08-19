package rexMas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	MasService struct {
		BaseService
	}
)

func NewMasService(SdkCtx *sdkCtx.SdkCtx) MasService {
	return MasService{
		BaseService: NewBaseService(SdkCtx),
	}
}
