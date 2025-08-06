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
	BaseService interface {
		Bootstrap(ctx context.Context, params *rexTypes.UpsBaseBootstrapReq) (result *rexTypes.UpsBaseBootstrapResp, err error)
		Create(ctx context.Context, params *rexTypes.AllowCreateModelObject) (result *rexTypes.ObjectApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.ObjectApiFormIdReq) (result *rexTypes.ObjectApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.ObjectApiJsonIdsReq) (result *rexTypes.ObjectApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelObject) (result *rexTypes.ObjectApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelObject) (result *rexTypes.ObjectApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.ObjectApiFormIdReq) (result *rexTypes.ModelObject, err error)
		QueryWhereObjectId(ctx context.Context, params *rexTypes.ObjectApiQueryWhreObjectIdReq) (result *rexTypes.ModelObject, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.ObjectApiJsonIdsReq) (result *rexTypes.ObjectCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.ObjectCommonSearchParams) (result *rexTypes.ObjectCommonQueryListResp, err error)
	}
	defaultBaseService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewBaseService(rexCtx *rexCtx.EngineCtx) BaseService {
	return &defaultBaseService{
		Svc:    "ups",
		rexCtx: rexCtx,
	}
}

func (m *defaultBaseService) Bootstrap(ctx context.Context, params *rexTypes.UpsBaseBootstrapReq) (result *rexTypes.UpsBaseBootstrapResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.UpsBaseBootstrapResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/base/bootstrap", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:BaseService:Bootstrap fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) Create(ctx context.Context, params *rexTypes.AllowCreateModelObject) (result *rexTypes.ObjectApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ObjectApiCreateResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/object/create", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:Create fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) Delete(ctx context.Context, params *rexTypes.ObjectApiFormIdReq) (result *rexTypes.ObjectApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ObjectApiOKResp]{}
	relativePath := "/ups/object/delete"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/object/delete?id=%d", params.Id)
	}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:Delete fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) DeleteMany(ctx context.Context, params *rexTypes.ObjectApiJsonIdsReq) (result *rexTypes.ObjectApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ObjectApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/object/deleteMany", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:DeleteMany fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelObject) (result *rexTypes.ObjectApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ObjectApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/object/update", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:Update  fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelObject) (result *rexTypes.ObjectApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ObjectApiOKResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/object/updateStatus", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:UpdateStatus fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) Query(ctx context.Context, params *rexTypes.ObjectApiFormIdReq) (result *rexTypes.ModelObject, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelObject]{}
	relativePath := "/ups/object/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/object/query?id=%d", params.Id)
	}
	logx.Infof("rex sdk: request path: %s", relativePath)
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
	if err != nil {
		logx.Errorf("rex sdk: request ups:ObjectService:Query error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:Query fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) QueryWhereObjectId(ctx context.Context, params *rexTypes.ObjectApiQueryWhreObjectIdReq) (result *rexTypes.ModelObject, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelObject]{}
	relativePath := "/ups/object/query"
	if params.ObjectId != "" {
		relativePath = fmt.Sprintf("/ups/object/query?object_id=%s", params.ObjectId)
	}
	logx.Infof("rex sdk: request path: %s", relativePath)
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
	if err != nil {
		logx.Errorf("rex sdk: request ups:ObjectService:Query error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:Query fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) QueryListWhereIds(ctx context.Context, params *rexTypes.ObjectApiJsonIdsReq) (result *rexTypes.ObjectCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ObjectCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/object/queryListWhereIds", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ObjectService:QueryListWhereIds error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:QueryListWhereIds fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultBaseService) QueryList(ctx context.Context, params *rexTypes.ObjectCommonSearchParams) (result *rexTypes.ObjectCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ObjectCommonQueryListResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/object/queryList", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ups:ObjectService:QueryList error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:ObjectService:QueryList fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
