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
	QxBaseService interface {
		Codes(ctx context.Context, params *rexTypes.CodesReq) (result *rexTypes.CodesResp, err error)
		Zones(ctx context.Context, params *rexTypes.ZonesReq) (result *rexTypes.ZonesResp, err error)
	}
	defaultQxBaseService struct {
		Svc    string
		rexCtx *rexCtx.QxCtx
	}
)

func NewQxBaseService(rexCtx *rexCtx.QxCtx) QxBaseService {
	return &defaultQxBaseService{
		Svc:    "base",
		rexCtx: rexCtx,
	}
}

func (m *defaultQxBaseService) Codes(ctx context.Context, params *rexTypes.CodesReq) (result *rexTypes.CodesResp, err error) {
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
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: codes fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultQxBaseService) Zones(ctx context.Context, params *rexTypes.ZonesReq) (result *rexTypes.ZonesResp, err error) {
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
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: zones fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}
