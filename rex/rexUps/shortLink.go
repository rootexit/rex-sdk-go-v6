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
	ShortLinkService interface {
		GetRedirectResult(ctx context.Context, params *rexTypes.GetRedirectResultReq) (code int32, result *rexTypes.GetRedirectResultResp, err error)
		Create(ctx context.Context, params *rexTypes.AllowCreateModelShortLink) (code int32, result *rexTypes.ShortLinkApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.ShortLinkApiFormIdReq) (code int32, result *rexTypes.ShortLinkApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.ShortLinkApiJsonIdsReq) (code int32, result *rexTypes.ShortLinkApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelShortLink) (code int32, result *rexTypes.ShortLinkApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelShortLink) (code int32, result *rexTypes.ShortLinkApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.ShortLinkApiFormIdReq) (code int32, result *rexTypes.ModelShortLink, err error)
		QueryWhereKey(ctx context.Context, params *rexTypes.ShortLinkApiFormKeyReq) (code int32, result *rexTypes.ModelShortLink, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.ShortLinkApiJsonIdsReq) (code int32, result *rexTypes.ShortLinkCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.ShortLinkCommonSearchParams) (code int32, result *rexTypes.ShortLinkCommonQueryListResp, err error)
	}
	defaultShortLinkService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewShortLinkService(rexCtx *rexCtx.EngineCtx) ShortLinkService {
	return &defaultShortLinkService{
		Svc:    "ups",
		rexCtx: rexCtx,
	}
}

func (m *defaultShortLinkService) GetRedirectResult(ctx context.Context, params *rexTypes.GetRedirectResultReq) (code int32, result *rexTypes.GetRedirectResultResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.GetRedirectResultResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/getRedirectResult", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request(code int32, resulterror: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request(code int32, result fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) Create(ctx context.Context, params *rexTypes.AllowCreateModelShortLink) (code int32, result *rexTypes.ShortLinkApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiCreateResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/create", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:Create  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:Create  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) Delete(ctx context.Context, params *rexTypes.ShortLinkApiFormIdReq) (code int32, result *rexTypes.ShortLinkApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiOKResp]{}
	relativePath := "/ups/shortLink/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/shortLink/delete?id=%d", params.Id)
	}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:Delete  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:Delete  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) DeleteMany(ctx context.Context, params *rexTypes.ShortLinkApiJsonIdsReq) (code int32, result *rexTypes.ShortLinkApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/deleteMany", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:DeleteMany  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:DeleteMany  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelShortLink) (code int32, result *rexTypes.ShortLinkApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/update", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:Update  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:Update  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelShortLink) (code int32, result *rexTypes.ShortLinkApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/updateStatus", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:UpdateStatus  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:UpdateStatus  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) Query(ctx context.Context, params *rexTypes.ShortLinkApiFormIdReq) (code int32, result *rexTypes.ModelShortLink, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelShortLink]{}
	relativePath := "/ups/shortLink/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/shortLink/query?id=%d", params.Id)
	}
	logx.Infof("rex sdk: request path: %s", relativePath)
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:Query  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:Query  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) QueryWhereKey(ctx context.Context, params *rexTypes.ShortLinkApiFormKeyReq) (code int32, result *rexTypes.ModelShortLink, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelShortLink]{}
	relativePath := "/ups/shortLink/queryWhereKey"
	if params.Key != "" {
		relativePath = fmt.Sprintf("/ups/shortLink/queryWhereKey?key=%d", params.Key)
	}
	logx.Infof("rex sdk: request path: %s", relativePath)
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:QueryWhereKey  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:QueryWhereKey  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) QueryListWhereIds(ctx context.Context, params *rexTypes.ShortLinkApiJsonIdsReq) (code int32, result *rexTypes.ShortLinkCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/queryListWhereIds", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:QueryListWhereIds  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:QueryListWhereIds  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultShortLinkService) QueryList(ctx context.Context, params *rexTypes.ShortLinkCommonSearchParams) (code int32, result *rexTypes.ShortLinkCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ShortLinkCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/queryList", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ShortLinkService:QueryList  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ups:ShortLinkService:QueryList  fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
