package rexTpas

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	TpasService struct {
		BaseService
		WechatOffiaccountService WechatOffiaccountService
	}
)

func NewTpasService(rexCtx *rexCtx.EngineCtx) TpasService {
	return TpasService{
		BaseService:              NewBaseService(rexCtx),
		WechatOffiaccountService: NewWechatOffiaccountService(rexCtx),
	}
}
