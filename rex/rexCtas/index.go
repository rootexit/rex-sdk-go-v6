package rexCtas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	CtasService struct {
		BaseService
		PeriodicJobService PeriodicJobService
	}
)

func NewCtasService(SdkCtx *sdkCtx.SdkCtx) CtasService {
	return CtasService{
		BaseService:        NewBaseService(SdkCtx),
		PeriodicJobService: NewPeriodicJobService(SdkCtx),
	}
}
