package rexMas

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	sdkCtx "github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	EmsConfigService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelEmsConfig) (code int32, result *rexTypes.EmsConfigApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.EmsConfigApiFormIdReq) (code int32, result *rexTypes.EmsConfigApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.EmsConfigApiJsonIdsReq) (code int32, result *rexTypes.EmsConfigApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelEmsConfig) (code int32, result *rexTypes.EmsConfigApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelEmsConfig) (code int32, result *rexTypes.EmsConfigApiOKResp, err error)
		UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelEmsConfig) (code int32, result *rexTypes.EmsConfigApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.EmsConfigApiFormIdReq) (code int32, result *rexTypes.ModelEmsConfig, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.EmsConfigApiJsonIdsReq) (code int32, result *rexTypes.EmsConfigCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.EmsConfigCommonSearchParams) (code int32, result *rexTypes.EmsConfigCommonQueryListResp, err error)
	}

	defaultEmsConfigService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewEmsConfigService(SdkCtx *sdkCtx.SdkCtx) EmsConfigService {
	return &defaultEmsConfigService{Svc: "mas", SdkCtx: SdkCtx}
}

func (m *defaultEmsConfigService) do(ctx context.Context, path string, method string, body interface{}, out interface{}) (int32, error) {
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, path, method, body)
	if err != nil {
		logx.Errorf("rex sdk: request mas:EmsConfigService error: %v", err)
		return rexCodes.FAIL, err
	}
	if err = json.Unmarshal(res, out); err != nil {
		return rexCodes.FAIL, err
	}
	return rexCodes.OK, nil
}

func (m *defaultEmsConfigService) Create(ctx context.Context, params *rexTypes.AllowCreateModelEmsConfig) (code int32, result *rexTypes.EmsConfigApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.EmsConfigApiCreateResp]{}
	code, err = m.do(ctx, "/mas/emsConfig/create", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultEmsConfigService) Delete(ctx context.Context, params *rexTypes.EmsConfigApiFormIdReq) (code int32, result *rexTypes.EmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.EmsConfigApiOKResp]{}
	code, err = m.do(ctx, fmt.Sprintf("/mas/emsConfig/delete?id=%d", params.Id), http.MethodPost, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultEmsConfigService) DeleteMany(ctx context.Context, params *rexTypes.EmsConfigApiJsonIdsReq) (code int32, result *rexTypes.EmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.EmsConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/emsConfig/deleteMany", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultEmsConfigService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelEmsConfig) (code int32, result *rexTypes.EmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.EmsConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/emsConfig/update", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultEmsConfigService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelEmsConfig) (code int32, result *rexTypes.EmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.EmsConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/emsConfig/updateStatus", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultEmsConfigService) UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelEmsConfig) (code int32, result *rexTypes.EmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.EmsConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/emsConfig/updateDefault", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultEmsConfigService) Query(ctx context.Context, params *rexTypes.EmsConfigApiFormIdReq) (code int32, result *rexTypes.ModelEmsConfig, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelEmsConfig]{}
	code, err = m.do(ctx, fmt.Sprintf("/mas/emsConfig/query?id=%d", params.Id), http.MethodGet, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultEmsConfigService) QueryListWhereIds(ctx context.Context, params *rexTypes.EmsConfigApiJsonIdsReq) (code int32, result *rexTypes.EmsConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.EmsConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/mas/emsConfig/queryListWhereIds", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultEmsConfigService) QueryList(ctx context.Context, params *rexTypes.EmsConfigCommonSearchParams) (code int32, result *rexTypes.EmsConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.EmsConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/mas/emsConfig/queryList", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}
