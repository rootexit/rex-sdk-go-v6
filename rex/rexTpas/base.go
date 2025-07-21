package rexTpas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	BaseService interface {
	}

	defaultBaseService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewBaseService(rexCtx *rexCtx.EngineCtx) BaseService {
	return &defaultBaseService{
		Svc:    "tpas",
		rexCtx: rexCtx,
	}
}
