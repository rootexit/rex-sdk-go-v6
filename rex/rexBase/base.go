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
		Codes(ctx context.Context, params *rexTypes.CodesReq) (result *rexTypes.CodesResp, err error)
		Zones(ctx context.Context, params *rexTypes.ZonesReq) (result *rexTypes.ZonesResp, err error)
	}
	defaultBaseService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewQxBaseService(rexCtx *rexCtx.EngineCtx) BaseService {
	return &defaultBaseService{
		Svc:    "base",
		rexCtx: rexCtx,
	}
}

func (m *defaultBaseService) Codes(ctx context.Context, params *rexTypes.CodesReq) (result *rexTypes.CodesResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CodesResp]{}
	relativePath := ""
	if params.Lang != "" || params.Svc != "" {
		v, _ := query.Values(params)
		relativePath = fmt.Sprintf("/base/codes?%s", v.Encode())
	} else {
		relativePath = "/base/codes"
	}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, nil)
	if err != nil {
		logx.Errorf("rex sdk: request base:BaseService:Codes error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request base:BaseService:Codes fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) Zones(ctx context.Context, params *rexTypes.ZonesReq) (result *rexTypes.ZonesResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ZonesResp]{}
	relativePath := ""
	if params.Lang != "" {
		v, _ := query.Values(params)
		relativePath = fmt.Sprintf("/base/zones?%s", v.Encode())
	} else {
		relativePath = "/base/zones"
	}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, nil)
	if err != nil {
		logx.Errorf("rex sdk: request base:BaseService:Zones error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk:request base:BaseService:Zones fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}
