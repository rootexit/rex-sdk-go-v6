package rexSas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	SasService struct {
		SasBaseService
	}
)

func NewSasService(rexCtx *rexCtx.QxCtx) SasService {
	return SasService{
		SasBaseService: NewSasBaseService(rexCtx),
	}
}
