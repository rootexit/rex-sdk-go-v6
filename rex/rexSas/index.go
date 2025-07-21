package rexSas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	SasService struct {
		BaseService
	}
)

func NewSasService(rexCtx *rexCtx.EngineCtx) SasService {
	return SasService{
		BaseService: NewBaseService(rexCtx),
	}
}
