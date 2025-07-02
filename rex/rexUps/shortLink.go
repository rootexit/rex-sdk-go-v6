package rexUps

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	UpsShortLinkService interface {
		GetRedirectResult(ctx context.Context, params *rexTypes.GetRedirectResultReq) (result *rexTypes.GetRedirectResultResp, err error)
		Create(ctx context.Context, params *rexTypes.AllowCreateModelShortLink) (result *rexTypes.ShortLinkApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.ShortLinkApiFormIdReq) (result *rexTypes.ShortLinkApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.ShortLinkApiJsonIdsReq) (result *rexTypes.ShortLinkApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelShortLink) (result *rexTypes.ShortLinkApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelShortLink) (result *rexTypes.ShortLinkApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.ShortLinkApiFormIdReq) (result *rexTypes.ModelShortLink, err error)
		QueryWhereKey(ctx context.Context, params *rexTypes.ShortLinkApiFormKeyReq) (result *rexTypes.ModelShortLink, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.ShortLinkApiJsonIdsReq) (result *rexTypes.ShortLinkCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.ShortLinkCommonSearchParams) (result *rexTypes.ShortLinkCommonQueryListResp, err error)
	}
	defaultUpsShortLinkService struct {
		Svc    string
		rexCtx *rexCtx.QxCtx
	}
)

func NewUpsShortLinkService(rexCtx *rexCtx.QxCtx) UpsShortLinkService {
	return &defaultUpsShortLinkService{
		Svc:    "ups",
		rexCtx: rexCtx,
	}
}

func (m *defaultUpsShortLinkService) GetRedirectResult(ctx context.Context, params *rexTypes.GetRedirectResultReq) (result *rexTypes.GetRedirectResultResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.GetRedirectResultResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/getRedirectResult", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) Create(ctx context.Context, params *rexTypes.AllowCreateModelShortLink) (result *rexTypes.ShortLinkApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiCreateResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/create", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) Delete(ctx context.Context, params *rexTypes.ShortLinkApiFormIdReq) (result *rexTypes.ShortLinkApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiOKResp]{}
	relativePath := "/ups/shortLink/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/shortLink/delete?id=%d", params.Id)
	}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) DeleteMany(ctx context.Context, params *rexTypes.ShortLinkApiJsonIdsReq) (result *rexTypes.ShortLinkApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/deleteMany", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelShortLink) (result *rexTypes.ShortLinkApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/update", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelShortLink) (result *rexTypes.ShortLinkApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/updateStatus", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) Query(ctx context.Context, params *rexTypes.ShortLinkApiFormIdReq) (result *rexTypes.ModelShortLink, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelShortLink]{}
	relativePath := "/ups/shortLink/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/shortLink/query?id=%d", params.Id)
	}
	logx.Infof("rex sdk: request path: %s", relativePath)
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) QueryWhereKey(ctx context.Context, params *rexTypes.ShortLinkApiFormKeyReq) (result *rexTypes.ModelShortLink, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelShortLink]{}
	relativePath := "/ups/shortLink/queryWhereKey"
	if params.Key != "" {
		relativePath = fmt.Sprintf("/ups/shortLink/queryWhereKey?key=%d", params.Key)
	}
	logx.Infof("rex sdk: request path: %s", relativePath)
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) QueryListWhereIds(ctx context.Context, params *rexTypes.ShortLinkApiJsonIdsReq) (result *rexTypes.ShortLinkCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/queryListWhereIds", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultUpsShortLinkService) QueryList(ctx context.Context, params *rexTypes.ShortLinkCommonSearchParams) (result *rexTypes.ShortLinkCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/queryList", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
