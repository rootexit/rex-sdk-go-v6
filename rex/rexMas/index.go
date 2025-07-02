package rexMas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	MasService struct {
		MasBaseService
	}
)

func NewMasService(rexCtx *rexCtx.QxCtx) MasService {
	return MasService{
		MasBaseService: NewMsgBaseService(rexCtx),
	}
}
