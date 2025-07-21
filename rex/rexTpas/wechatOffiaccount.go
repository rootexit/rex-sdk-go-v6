package rexTpas

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	WechatOffiaccountService interface {
		// note: 为某个公众号强制刷新凭证
		ForceRefreshOffiaccountAccessToken(ctx context.Context, params *rexTypes.WechatForceRefreshOffiaccountAccessTokenReq) (result *rexTypes.WechatForceRefreshOffiaccountAccessTokenResp, err error)
		// note: 获取公众号普通AccessToken
		GetAccessToken(ctx context.Context, params *rexTypes.WechatOffiaccountGetAccessTokenReq) (result *rexTypes.WechatOffiaccountGetAccessTokenResp, err error)
		// note: 获取公众号JsApiTicket
		GetJsApiTicket(ctx context.Context, params *rexTypes.WechatOffiaccountGetJsApiTicketReq) (result *rexTypes.WechatOffiaccountGetJsApiTicketResp, err error)
		// note: 生成重定向授权链接
		GenShareConfig(ctx context.Context, params *rexTypes.WechatOffiaccountGenShareConfigReq) (result *rexTypes.WechatOffiaccountGenShareConfigResp, err error)
		// note: 生成分享配置注入
		GenRedirectUrl(ctx context.Context, params *rexTypes.WechatOffiaccountGenRedirectUrlReq) (result *rexTypes.WechatOffiaccountGenRedirectUrlResp, err error)
	}

	defaultWechatOffiaccountService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewWechatOffiaccountService(rexCtx *rexCtx.EngineCtx) WechatOffiaccountService {
	return &defaultWechatOffiaccountService{
		Svc:    "tpas",
		rexCtx: rexCtx,
	}
}

func (m *defaultWechatOffiaccountService) ForceRefreshOffiaccountAccessToken(ctx context.Context, params *rexTypes.WechatForceRefreshOffiaccountAccessTokenReq) (result *rexTypes.WechatForceRefreshOffiaccountAccessTokenResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatForceRefreshOffiaccountAccessTokenResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/tpas/wechatOffiaccount/forceRefreshOffiaccountAccessToken", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: tpas:WechatOffiaccountService:ForceRefreshOffiaccountAccessToken fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultWechatOffiaccountService) GetAccessToken(ctx context.Context, params *rexTypes.WechatOffiaccountGetAccessTokenReq) (result *rexTypes.WechatOffiaccountGetAccessTokenResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatOffiaccountGetAccessTokenResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/tpas/wechatOffiaccount/getAccessToken", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request tpas:WechatOffiaccountService:GetAccessToken error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: tpas:WechatOffiaccountService:GetAccessToken fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultWechatOffiaccountService) GetJsApiTicket(ctx context.Context, params *rexTypes.WechatOffiaccountGetJsApiTicketReq) (result *rexTypes.WechatOffiaccountGetJsApiTicketResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatOffiaccountGetJsApiTicketResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/tpas/wechatOffiaccount/getJsApiTicket", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request tpas:WechatOffiaccountService:GetJsApiTicket error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: tpas:WechatOffiaccountService:GetJsApiTicket fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultWechatOffiaccountService) GenShareConfig(ctx context.Context, params *rexTypes.WechatOffiaccountGenShareConfigReq) (result *rexTypes.WechatOffiaccountGenShareConfigResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatOffiaccountGenShareConfigResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/tpas/wechatOffiaccount/genShareConfig", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request tpas:WechatOffiaccountService:GenShareConfig error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: tpas:WechatOffiaccountService:GenShareConfig fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultWechatOffiaccountService) GenRedirectUrl(ctx context.Context, params *rexTypes.WechatOffiaccountGenRedirectUrlReq) (result *rexTypes.WechatOffiaccountGenRedirectUrlResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatOffiaccountGenRedirectUrlResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/tpas/wechatOffiaccount/genRedirectUrl", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request tpas:WechatOffiaccountService:GenRedirectUrl error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: tpas:WechatOffiaccountService:GenRedirectUrl fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
