package rexBase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	sdkCtx "github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	BaseService interface {
		Codes(ctx context.Context, params *rexTypes.CodesReq) (code int32, result *rexTypes.CodesResp, err error)
		Zones(ctx context.Context, params *rexTypes.ZonesReq) (code int32, result *rexTypes.ZonesResp, err error)
		GetIpMy(ctx context.Context, params *rexTypes.GetIpMyReq) (code int32, result *rexTypes.ModelClient, err error)
		GetIpQuery(ctx context.Context, params *rexTypes.GetIpQueryReq) (code int32, result *rexTypes.GetIpQueryResp, err error)
		AddSomePolicy(ctx context.Context, params *rexTypes.AddSomePolicyReq) (code int32, result *rexTypes.OkResp, err error)
		RemoveSomePolicy(ctx context.Context, params *rexTypes.RemoveSomePolicyReq) (code int32, result *rexTypes.OkResp, err error)
	}
	defaultBaseService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewBaseService(SdkCtx *sdkCtx.SdkCtx) BaseService {
	return &defaultBaseService{
		Svc:    "base",
		SdkCtx: SdkCtx,
	}
}

func (m *defaultBaseService) Codes(ctx context.Context, params *rexTypes.CodesReq) (code int32, result *rexTypes.CodesResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CodesResp]{}
	relativePath := "/base/codes"
	if params.Lang != "" && params.Svc != "" {
		relativePath = fmt.Sprintf("/base/codes?lang=%s&svc=%s", params.Lang, params.Svc)
	} else if params.Lang != "" {
		relativePath = fmt.Sprintf("/base/codes?lang=%s", params.Lang)
	} else if params.Svc != "" {
		relativePath = fmt.Sprintf("/base/codes?svc=%s", params.Svc)
	}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, nil)
	if err != nil {
		logx.Errorf("rex sdk: request base:BaseService:Codes error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request base:BaseService:Codes fail: %v", tmp)
		return rexCodes.OK, &tmp.Data, nil
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) Zones(ctx context.Context, params *rexTypes.ZonesReq) (code int32, result *rexTypes.ZonesResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ZonesResp]{}
	relativePath := ""
	if params.Lang != "" {
		v, _ := query.Values(params)
		relativePath = fmt.Sprintf("/base/zones?%s", v.Encode())
	} else {
		relativePath = "/base/zones"
	}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, nil)
	if err != nil {
		logx.Errorf("rex sdk: request base:BaseService:Zones error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk:request base:BaseService:Zones fail: %v", tmp)
		return rexCodes.OK, &tmp.Data, nil
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) GetIpMy(ctx context.Context, params *rexTypes.GetIpMyReq) (code int32, result *rexTypes.ModelClient, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelClient]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/base/ip/my", http.MethodGet, nil)
	if err != nil {
		logx.Errorf("rex sdk: request base:BaseService:GetIpMy error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request base:BaseService:GetIpMy fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) GetIpQuery(ctx context.Context, params *rexTypes.GetIpQueryReq) (code int32, result *rexTypes.GetIpQueryResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.GetIpQueryResp]{}
	v, _ := query.Values(params)
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, fmt.Sprintf("/base/ip/query?%s", v.Encode()), http.MethodGet, nil)
	if err != nil {
		logx.Errorf("rex sdk: request base:BaseService:GetIpQuery error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request base:BaseService:GetIpQuery fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) AddSomePolicy(ctx context.Context, params *rexTypes.AddSomePolicyReq) (code int32, result *rexTypes.OkResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.OkResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/base/policy/add", http.MethodPost, &params)
	if err != nil {
		logx.Errorf("rex sdk: request base:BaseService:AddSomePolicy error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request base:BaseService:AddSomePolicy fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) RemoveSomePolicy(ctx context.Context, params *rexTypes.RemoveSomePolicyReq) (code int32, result *rexTypes.OkResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.OkResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/base/policy/remove", http.MethodPost, &params)
	if err != nil {
		logx.Errorf("rex sdk: request base:BaseService:RemoveSomePolicy error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request base:BaseService:RemoveSomePolicy fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
