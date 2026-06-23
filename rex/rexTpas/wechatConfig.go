package rexTpas

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
	WechatConfigService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelWechatConfig) (code int32, result *rexTypes.WechatConfigApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.WechatConfigApiFormIdReq) (code int32, result *rexTypes.WechatConfigApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.WechatConfigApiJsonIdsReq) (code int32, result *rexTypes.WechatConfigApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelWechatConfig) (code int32, result *rexTypes.WechatConfigApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelWechatConfig) (code int32, result *rexTypes.WechatConfigApiOKResp, err error)
		UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelWechatConfig) (code int32, result *rexTypes.WechatConfigApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.WechatConfigApiFormIdReq) (code int32, result *rexTypes.ModelWechatConfig, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.WechatConfigApiJsonIdsReq) (code int32, result *rexTypes.WechatConfigCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.WechatConfigCommonSearchParams) (code int32, result *rexTypes.WechatConfigCommonQueryListResp, err error)
	}

	defaultWechatConfigService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewWechatConfigService(SdkCtx *sdkCtx.SdkCtx) WechatConfigService {
	return &defaultWechatConfigService{Svc: "tpas", SdkCtx: SdkCtx}
}

func (m *defaultWechatConfigService) do(ctx context.Context, path string, method string, body interface{}, out interface{}) (int32, error) {
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, path, method, body)
	if err != nil {
		logx.Errorf("rex sdk: request tpas:WechatConfigService error: %v", err)
		return rexCodes.FAIL, err
	}
	if err = json.Unmarshal(res, out); err != nil {
		return rexCodes.FAIL, err
	}
	return rexCodes.OK, nil
}

func wechatConfigResultWithErr[T any](code int32, data *T, responseCode int32, msg string, err error) (int32, *T, error) {
	if err != nil {
		return code, nil, err
	}
	return responseCode, data, errors.New(msg)
}

func (m *defaultWechatConfigService) Create(ctx context.Context, params *rexTypes.AllowCreateModelWechatConfig) (code int32, result *rexTypes.WechatConfigApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatConfigApiCreateResp]{}
	code, err = m.do(ctx, "/tpas/wechatConfig/create", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultWechatConfigService) Delete(ctx context.Context, params *rexTypes.WechatConfigApiFormIdReq) (code int32, result *rexTypes.WechatConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatConfigApiOKResp]{}
	code, err = m.do(ctx, fmt.Sprintf("/tpas/wechatConfig/delete?id=%d", params.Id), http.MethodPost, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultWechatConfigService) DeleteMany(ctx context.Context, params *rexTypes.WechatConfigApiJsonIdsReq) (code int32, result *rexTypes.WechatConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatConfigApiOKResp]{}
	code, err = m.do(ctx, "/tpas/wechatConfig/deleteMany", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultWechatConfigService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelWechatConfig) (code int32, result *rexTypes.WechatConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatConfigApiOKResp]{}
	code, err = m.do(ctx, "/tpas/wechatConfig/update", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultWechatConfigService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelWechatConfig) (code int32, result *rexTypes.WechatConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatConfigApiOKResp]{}
	code, err = m.do(ctx, "/tpas/wechatConfig/updateStatus", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultWechatConfigService) UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelWechatConfig) (code int32, result *rexTypes.WechatConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatConfigApiOKResp]{}
	code, err = m.do(ctx, "/tpas/wechatConfig/updateDefault", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultWechatConfigService) Query(ctx context.Context, params *rexTypes.WechatConfigApiFormIdReq) (code int32, result *rexTypes.ModelWechatConfig, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelWechatConfig]{}
	code, err = m.do(ctx, fmt.Sprintf("/tpas/wechatConfig/query?id=%d", params.Id), http.MethodGet, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultWechatConfigService) QueryListWhereIds(ctx context.Context, params *rexTypes.WechatConfigApiJsonIdsReq) (code int32, result *rexTypes.WechatConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/tpas/wechatConfig/queryListWhereIds", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultWechatConfigService) QueryList(ctx context.Context, params *rexTypes.WechatConfigCommonSearchParams) (code int32, result *rexTypes.WechatConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.WechatConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/tpas/wechatConfig/queryList", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return wechatConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}
