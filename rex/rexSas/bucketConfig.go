package rexSas

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
	BucketConfigService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelBucketConfig) (code int32, result *rexTypes.BucketConfigApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.BucketConfigApiFormIdReq) (code int32, result *rexTypes.BucketConfigApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.BucketConfigApiJsonIdsReq) (code int32, result *rexTypes.BucketConfigApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelBucketConfig) (code int32, result *rexTypes.BucketConfigApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelBucketConfig) (code int32, result *rexTypes.BucketConfigApiOKResp, err error)
		UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelBucketConfig) (code int32, result *rexTypes.BucketConfigApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.BucketConfigApiFormIdReq) (code int32, result *rexTypes.ModelBucketConfig, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.BucketConfigApiJsonIdsReq) (code int32, result *rexTypes.BucketConfigCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.BucketConfigCommonSearchParams) (code int32, result *rexTypes.BucketConfigCommonQueryListResp, err error)
	}

	defaultBucketConfigService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewBucketConfigService(SdkCtx *sdkCtx.SdkCtx) BucketConfigService {
	return &defaultBucketConfigService{Svc: "sas", SdkCtx: SdkCtx}
}

func (m *defaultBucketConfigService) do(ctx context.Context, path string, method string, body interface{}, out interface{}) (int32, error) {
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, path, method, body)
	if err != nil {
		logx.Errorf("rex sdk: request sas:BucketConfigService error: %v", err)
		return rexCodes.FAIL, err
	}
	if err = json.Unmarshal(res, out); err != nil {
		return rexCodes.FAIL, err
	}
	return rexCodes.OK, nil
}

func bucketConfigResultWithErr[T any](code int32, data *T, responseCode int32, msg string, err error) (int32, *T, error) {
	if err != nil {
		return code, nil, err
	}
	return responseCode, data, errors.New(msg)
}

func (m *defaultBucketConfigService) Create(ctx context.Context, params *rexTypes.AllowCreateModelBucketConfig) (code int32, result *rexTypes.BucketConfigApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BucketConfigApiCreateResp]{}
	code, err = m.do(ctx, "/sas/bucketConfig/create", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBucketConfigService) Delete(ctx context.Context, params *rexTypes.BucketConfigApiFormIdReq) (code int32, result *rexTypes.BucketConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BucketConfigApiOKResp]{}
	code, err = m.do(ctx, fmt.Sprintf("/sas/bucketConfig/delete?id=%d", params.Id), http.MethodPost, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBucketConfigService) DeleteMany(ctx context.Context, params *rexTypes.BucketConfigApiJsonIdsReq) (code int32, result *rexTypes.BucketConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BucketConfigApiOKResp]{}
	code, err = m.do(ctx, "/sas/bucketConfig/deleteMany", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBucketConfigService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelBucketConfig) (code int32, result *rexTypes.BucketConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BucketConfigApiOKResp]{}
	code, err = m.do(ctx, "/sas/bucketConfig/update", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBucketConfigService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelBucketConfig) (code int32, result *rexTypes.BucketConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BucketConfigApiOKResp]{}
	code, err = m.do(ctx, "/sas/bucketConfig/updateStatus", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBucketConfigService) UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelBucketConfig) (code int32, result *rexTypes.BucketConfigApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BucketConfigApiOKResp]{}
	code, err = m.do(ctx, "/sas/bucketConfig/updateDefault", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBucketConfigService) Query(ctx context.Context, params *rexTypes.BucketConfigApiFormIdReq) (code int32, result *rexTypes.ModelBucketConfig, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelBucketConfig]{}
	code, err = m.do(ctx, fmt.Sprintf("/sas/bucketConfig/query?id=%d", params.Id), http.MethodGet, nil, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBucketConfigService) QueryListWhereIds(ctx context.Context, params *rexTypes.BucketConfigApiJsonIdsReq) (code int32, result *rexTypes.BucketConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BucketConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/sas/bucketConfig/queryListWhereIds", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBucketConfigService) QueryList(ctx context.Context, params *rexTypes.BucketConfigCommonSearchParams) (code int32, result *rexTypes.BucketConfigCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BucketConfigCommonQueryListResp]{}
	code, err = m.do(ctx, "/sas/bucketConfig/queryList", http.MethodPost, &params, tmp)
	if err != nil || tmp.Code != rexCodes.OK {
		return bucketConfigResultWithErr(code, &tmp.Data, tmp.Code, tmp.Msg, err)
	}
	return rexCodes.OK, &tmp.Data, nil
}
