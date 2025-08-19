package rexBase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	BaseService interface {
		Codes(ctx context.Context, params *rexTypes.CodesReq) (code int32, result *rexTypes.CodesResp, err error)
		Zones(ctx context.Context, params *rexTypes.ZonesReq) (code int32, result *rexTypes.ZonesResp, err error)
	}
	defaultBaseService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewQxBaseService(SdkCtx *sdkCtx.SdkCtx) BaseService {
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
