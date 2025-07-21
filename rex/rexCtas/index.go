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

func NewCtasService(rexCtx *rexCtx.EngineCtx) CtasService {
	return CtasService{
		BaseService:        NewBaseService(rexCtx),
		PeriodicJobService: NewPeriodicJobService(rexCtx),
	}
}
