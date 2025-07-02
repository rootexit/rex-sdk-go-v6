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
	UpsIndustryService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelIndustry) (result *rexTypes.IndustryApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.IndustryApiFormIdReq) (result *rexTypes.IndustryApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.IndustryApiJsonIdsReq) (result *rexTypes.IndustryApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelIndustry) (result *rexTypes.IndustryApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelIndustry) (result *rexTypes.IndustryApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.IndustryApiFormIdReq) (result *rexTypes.ModelIndustry, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.IndustryApiJsonIdsReq) (result *rexTypes.IndustryCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.IndustryCommonSearchParams) (result *rexTypes.IndustryCommonQueryListResp, err error)
	}
	defaultUpsIndustryService struct {
		Svc    string
		rexCtx *rexCtx.QxCtx
	}
)

func NewUpsIndustryService(rexCtx *rexCtx.QxCtx) UpsIndustryService {
	return &defaultUpsIndustryService{
		Svc:    "ups",
		rexCtx: rexCtx,
	}
}

func (m *defaultUpsIndustryService) Create(ctx context.Context, params *rexTypes.AllowCreateModelIndustry) (result *rexTypes.IndustryApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.IndustryApiCreateResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/create", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) Delete(ctx context.Context, params *rexTypes.IndustryApiFormIdReq) (result *rexTypes.IndustryApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.IndustryApiOKResp]{}
	relativePath := "/ups/industry/delete"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/industry/delete?id=%d", params.Id)
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

func (m *defaultUpsIndustryService) DeleteMany(ctx context.Context, params *rexTypes.IndustryApiJsonIdsReq) (result *rexTypes.IndustryApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.IndustryApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/deleteMany", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelIndustry) (result *rexTypes.IndustryApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.IndustryApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/update", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelIndustry) (result *rexTypes.IndustryApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.IndustryApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/updateStatus", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) Query(ctx context.Context, params *rexTypes.IndustryApiFormIdReq) (result *rexTypes.ModelIndustry, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelIndustry]{}
	relativePath := "/ups/industry/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/industry/query?id=%d", params.Id)
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

func (m *defaultUpsIndustryService) QueryListWhereIds(ctx context.Context, params *rexTypes.IndustryApiJsonIdsReq) (result *rexTypes.IndustryCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.IndustryCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/queryListWhereIds", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) QueryList(ctx context.Context, params *rexTypes.IndustryCommonSearchParams) (result *rexTypes.IndustryCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.IndustryCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/queryList", http.MethodPost, &params)

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
