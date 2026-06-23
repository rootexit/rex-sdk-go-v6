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
	SmsConfigService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelSmsConfig) (code int32, result *rexTypes.SmsConfigApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.SmsConfigApiFormIdReq) (code int32, result *rexTypes.SmsConfigApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.SmsConfigApiJsonIdsReq) (code int32, result *rexTypes.SmsConfigApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelSmsConfig) (code int32, result *rexTypes.SmsConfigApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelSmsConfig) (code int32, result *rexTypes.SmsConfigApiOKResp, err error)
		UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelSmsConfig) (code int32, result *rexTypes.SmsConfigApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.SmsConfigApiFormIdReq) (code int32, result *rexTypes.ModelSmsConfig, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.SmsConfigApiJsonIdsReq) (code int32, result *rexTypes.SmsConfigCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.SmsConfigCommonSearchParams) (code int32, result *rexTypes.SmsConfigCommonQueryListResp, err error)
	}

	defaultSmsConfigService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewSmsConfigService(SdkCtx *sdkCtx.SdkCtx) SmsConfigService {
	return &defaultSmsConfigService{Svc: "mas", SdkCtx: SdkCtx}
}

func (m *defaultSmsConfigService) do(ctx context.Context, path string, method string, body interface{}, out interface{}) (int32, error) {
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, path, method, body)
	if err != nil {
		logx.Errorf("rex sdk: request mas:SmsConfigService error: %v", err)
		return rexCodes.FAIL, err
	}
	if err = json.Unmarshal(res, out); err != nil {
		return rexCodes.FAIL, err
	}
	return rexCodes.OK, nil
}

func (m *defaultSmsConfigService) Create(ctx context.Context, params *rexTypes.AllowCreateModelSmsConfig) (code int32, result *rexTypes.SmsConfigApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsConfigApiCreateResp]{}
	code, err = m.do(ctx, "/mas/smsConfig/create", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSmsConfigService) Delete(ctx context.Context, params *rexTypes.SmsConfigApiFormIdReq) (code int32, result *rexTypes.SmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsConfigApiOKResp]{}
	code, err = m.do(ctx, fmt.Sprintf("/mas/smsConfig/delete?id=%d", params.Id), http.MethodPost, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSmsConfigService) DeleteMany(ctx context.Context, params *rexTypes.SmsConfigApiJsonIdsReq) (code int32, result *rexTypes.SmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/smsConfig/deleteMany", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSmsConfigService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelSmsConfig) (code int32, result *rexTypes.SmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/smsConfig/update", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSmsConfigService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelSmsConfig) (code int32, result *rexTypes.SmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/smsConfig/updateStatus", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSmsConfigService) UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelSmsConfig) (code int32, result *rexTypes.SmsConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/smsConfig/updateDefault", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSmsConfigService) Query(ctx context.Context, params *rexTypes.SmsConfigApiFormIdReq) (code int32, result *rexTypes.ModelSmsConfig, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelSmsConfig]{}
	code, err = m.do(ctx, fmt.Sprintf("/mas/smsConfig/query?id=%d", params.Id), http.MethodGet, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSmsConfigService) QueryListWhereIds(ctx context.Context, params *rexTypes.SmsConfigApiJsonIdsReq) (code int32, result *rexTypes.SmsConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/mas/smsConfig/queryListWhereIds", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSmsConfigService) QueryList(ctx context.Context, params *rexTypes.SmsConfigCommonSearchParams) (code int32, result *rexTypes.SmsConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/mas/smsConfig/queryList", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}
