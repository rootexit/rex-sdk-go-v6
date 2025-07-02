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
	UpsTagService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelTag) (result *rexTypes.TagApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.TagApiFormIdReq) (result *rexTypes.TagApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.TagApiJsonIdsReq) (result *rexTypes.TagApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelTag) (result *rexTypes.TagApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelTag) (result *rexTypes.TagApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.TagApiFormIdReq) (result *rexTypes.ModelTag, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.TagApiJsonIdsReq) (result *rexTypes.TagCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.TagCommonSearchParams) (result *rexTypes.TagCommonQueryListResp, err error)
	}
	defaultUpsTagService struct {
		Svc    string
		rexCtx *rexCtx.QxCtx
	}
)

func NewUpsTagService(rexCtx *rexCtx.QxCtx) UpsTagService {
	return &defaultUpsTagService{
		Svc:    "ups",
		rexCtx: rexCtx,
	}
}

func (m *defaultUpsTagService) Create(ctx context.Context, params *rexTypes.AllowCreateModelTag) (result *rexTypes.TagApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiCreateResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/create", http.MethodPost, &params)

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

func (m *defaultUpsTagService) Delete(ctx context.Context, params *rexTypes.TagApiFormIdReq) (result *rexTypes.TagApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiOKResp]{}
	relativePath := "/ups/tag/delete"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/tag/query?id=%d", params.Id)
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

func (m *defaultUpsTagService) DeleteMany(ctx context.Context, params *rexTypes.TagApiJsonIdsReq) (result *rexTypes.TagApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/deleteMany", http.MethodPost, &params)

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

func (m *defaultUpsTagService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelTag) (result *rexTypes.TagApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/update", http.MethodPost, &params)

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

func (m *defaultUpsTagService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelTag) (result *rexTypes.TagApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/updateStatus", http.MethodPost, &params)

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

func (m *defaultUpsTagService) Query(ctx context.Context, params *rexTypes.TagApiFormIdReq) (result *rexTypes.ModelTag, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelTag]{}
	relativePath := "/ups/tag/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/tag/query?id=%d", params.Id)
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

func (m *defaultUpsTagService) QueryListWhereIds(ctx context.Context, params *rexTypes.TagApiJsonIdsReq) (result *rexTypes.TagCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/queryListWhereIds", http.MethodPost, &params)

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

func (m *defaultUpsTagService) QueryList(ctx context.Context, params *rexTypes.TagCommonSearchParams) (result *rexTypes.TagCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/queryList", http.MethodPost, &params)

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
