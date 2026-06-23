package rexMas

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	sdkCtx "github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	CaptchaConfigService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelCaptchaConfig) (code int32, result *rexTypes.CaptchaConfigApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.CaptchaConfigApiFormIdReq) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.CaptchaConfigApiJsonIdsReq) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelCaptchaConfig) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelCaptchaConfig) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error)
		UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelCaptchaConfig) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.CaptchaConfigApiFormIdReq) (code int32, result *rexTypes.ModelCaptchaConfig, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.CaptchaConfigApiJsonIdsReq) (code int32, result *rexTypes.CaptchaConfigCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.CaptchaConfigCommonSearchParams) (code int32, result *rexTypes.CaptchaConfigCommonQueryListResp, err error)
	}

	defaultCaptchaConfigService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewCaptchaConfigService(SdkCtx *sdkCtx.SdkCtx) CaptchaConfigService {
	return &defaultCaptchaConfigService{Svc: "mas", SdkCtx: SdkCtx}
}

func (m *defaultCaptchaConfigService) do(ctx context.Context, path string, method string, body interface{}, out interface{}) (int32, error) {
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, path, method, body)
	if err != nil {
		logx.Errorf("rex sdk: request mas:CaptchaConfigService error: %v", err)
		return rexCodes.FAIL, err
	}
	if err = json.Unmarshal(res, out); err != nil {
		return rexCodes.FAIL, err
	}
	return rexCodes.OK, nil
}

func (m *defaultCaptchaConfigService) Create(ctx context.Context, params *rexTypes.AllowCreateModelCaptchaConfig) (code int32, result *rexTypes.CaptchaConfigApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CaptchaConfigApiCreateResp]{}
	code, err = m.do(ctx, "/mas/captchaConfig/create", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCaptchaConfigService) Delete(ctx context.Context, params *rexTypes.CaptchaConfigApiFormIdReq) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CaptchaConfigApiOKResp]{}
	code, err = m.do(ctx, fmt.Sprintf("/mas/captchaConfig/delete?id=%d", params.Id), http.MethodPost, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCaptchaConfigService) DeleteMany(ctx context.Context, params *rexTypes.CaptchaConfigApiJsonIdsReq) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CaptchaConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/captchaConfig/deleteMany", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCaptchaConfigService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelCaptchaConfig) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CaptchaConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/captchaConfig/update", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCaptchaConfigService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelCaptchaConfig) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CaptchaConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/captchaConfig/updateStatus", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCaptchaConfigService) UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelCaptchaConfig) (code int32, result *rexTypes.CaptchaConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CaptchaConfigApiOKResp]{}
	code, err = m.do(ctx, "/mas/captchaConfig/updateDefault", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCaptchaConfigService) Query(ctx context.Context, params *rexTypes.CaptchaConfigApiFormIdReq) (code int32, result *rexTypes.ModelCaptchaConfig, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelCaptchaConfig]{}
	code, err = m.do(ctx, fmt.Sprintf("/mas/captchaConfig/query?id=%d", params.Id), http.MethodGet, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCaptchaConfigService) QueryListWhereIds(ctx context.Context, params *rexTypes.CaptchaConfigApiJsonIdsReq) (code int32, result *rexTypes.CaptchaConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CaptchaConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/mas/captchaConfig/queryListWhereIds", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCaptchaConfigService) QueryList(ctx context.Context, params *rexTypes.CaptchaConfigCommonSearchParams) (code int32, result *rexTypes.CaptchaConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CaptchaConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/mas/captchaConfig/queryList", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return resultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func resultWithErr[T any](code int32, data *T, responseCode int32, msg string, err error) (int32, *T, error) {
	if err != nil {
		return code, nil, err
	}
	return responseCode, data, errors.New(msg)
}
