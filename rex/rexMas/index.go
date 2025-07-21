package rexMas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	MasService struct {
		BaseService
	}
)

func NewMasService(rexCtx *rexCtx.EngineCtx) MasService {
	return MasService{
		BaseService: NewBaseService(rexCtx),
	}
}
