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

func NewTpasService(SdkCtx *sdkCtx.SdkCtx) TpasService {
	return TpasService{
		BaseService:              NewBaseService(SdkCtx),
		WechatOffiaccountService: NewWechatOffiaccountService(SdkCtx),
	}
}
