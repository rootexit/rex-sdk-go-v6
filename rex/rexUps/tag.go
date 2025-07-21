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
	TagService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelTag) (result *rexTypes.TagApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.TagApiFormIdReq) (result *rexTypes.TagApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.TagApiJsonIdsReq) (result *rexTypes.TagApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelTag) (result *rexTypes.TagApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelTag) (result *rexTypes.TagApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.TagApiFormIdReq) (result *rexTypes.ModelTag, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.TagApiJsonIdsReq) (result *rexTypes.TagCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.TagCommonSearchParams) (result *rexTypes.TagCommonQueryListResp, err error)
	}
	defaultTagService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewTagService(rexCtx *rexCtx.EngineCtx) TagService {
	return &defaultTagService{
		Svc:    "ups",
		rexCtx: rexCtx,
	}
}

func (m *defaultTagService) Create(ctx context.Context, params *rexTypes.AllowCreateModelTag) (result *rexTypes.TagApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiCreateResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/create", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:TagService:Create error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:TagService:Create fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultTagService) Delete(ctx context.Context, params *rexTypes.TagApiFormIdReq) (result *rexTypes.TagApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiOKResp]{}
	relativePath := "/ups/tag/delete"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/tag/query?id=%d", params.Id)
	}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:TagService:Delete error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:TagService:Delete fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultTagService) DeleteMany(ctx context.Context, params *rexTypes.TagApiJsonIdsReq) (result *rexTypes.TagApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/deleteMany", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:TagService:DeleteMany error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:TagService:DeleteMany fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultTagService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelTag) (result *rexTypes.TagApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/update", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:TagService:Update error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:TagService:Update fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultTagService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelTag) (result *rexTypes.TagApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/updateStatus", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:TagService:UpdateStatus error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:TagService:UpdateStatus fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultTagService) Query(ctx context.Context, params *rexTypes.TagApiFormIdReq) (result *rexTypes.ModelTag, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelTag]{}
	relativePath := "/ups/tag/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/tag/query?id=%d", params.Id)
	}
	logx.Infof("rex sdk: request path: %s", relativePath)
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
	if err != nil {
		logx.Errorf("rex sdk: request ups:TagService:Query error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:TagService:Query fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultTagService) QueryListWhereIds(ctx context.Context, params *rexTypes.TagApiJsonIdsReq) (result *rexTypes.TagCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/queryListWhereIds", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:TagService:QueryListWhereIds error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:TagService:QueryListWhereIds fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultTagService) QueryList(ctx context.Context, params *rexTypes.TagCommonSearchParams) (result *rexTypes.TagCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TagCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/queryList", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:TagService:QueryList error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:TagService:QueryList fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
