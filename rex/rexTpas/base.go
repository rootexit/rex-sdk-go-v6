package rexTpas

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	BaseService interface {
		WechatJobWebhook(ctx context.Context, params *rexTypes.TpasWechatJobWebhookReq) (code int32, result *rexTypes.TpasWechatJobWebhookResp, err error)
	}

	defaultBaseService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewBaseService(SdkCtx *sdkCtx.SdkCtx) BaseService {
	return &defaultBaseService{
		Svc:    "tpas",
		SdkCtx: SdkCtx,
	}
}

func (m *defaultBaseService) WechatJobWebhook(ctx context.Context, params *rexTypes.TpasWechatJobWebhookReq) (code int32, result *rexTypes.TpasWechatJobWebhookResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TpasWechatJobWebhookResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/tpas/wechatOffiaccount/job/webhook", http.MethodPost, &params)
	if err != nil {
		logx.Errorf("rex sdk: request tpas:BaseService:WechatJobWebhook error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request tpas:BaseService:WechatJobWebhook fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
