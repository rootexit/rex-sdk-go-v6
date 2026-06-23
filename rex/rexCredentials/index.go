package rexCredentials

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
	CredentialService interface {
		Create(ctx context.Context, params *rexTypes.AllowCreateModelCredential) (code int32, result *rexTypes.CredentialApiCreateResp, err error)
		Delete(ctx context.Context, params *rexTypes.CredentialApiFormIdReq) (code int32, result *rexTypes.CredentialApiOKResp, err error)
		DeleteMany(ctx context.Context, params *rexTypes.CredentialApiJsonIdsReq) (code int32, result *rexTypes.CredentialApiOKResp, err error)
		Update(ctx context.Context, params *rexTypes.AllowUpdateModelCredential) (code int32, result *rexTypes.CredentialApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelCredential) (code int32, result *rexTypes.CredentialApiOKResp, err error)
		UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelCredential) (code int32, result *rexTypes.CredentialApiOKResp, err error)
		Query(ctx context.Context, params *rexTypes.CredentialApiFormIdReq) (code int32, result *rexTypes.ModelCredential, err error)
		QueryListWhereIds(ctx context.Context, params *rexTypes.CredentialApiJsonIdsReq) (code int32, result *rexTypes.CredentialCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *rexTypes.CredentialCommonSearchParams) (code int32, result *rexTypes.CredentialCommonQueryListResp, err error)
	}

	defaultCredentialService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewCredentialService(SdkCtx *sdkCtx.SdkCtx) CredentialService {
	return &defaultCredentialService{
		Svc:    "credentials",
		SdkCtx: SdkCtx,
	}
}

func (m *defaultCredentialService) do(ctx context.Context, path string, method string, body interface{}, out interface{}) (int32, error) {
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, path, method, body)
	if err != nil {
		logx.Errorf("rex sdk: request credentials:CredentialService error: %v", err)
		return rexCodes.FAIL, err
	}
	if err = json.Unmarshal(res, out); err != nil {
		return rexCodes.FAIL, err
	}
	return rexCodes.OK, nil
}

func (m *defaultCredentialService) Create(ctx context.Context, params *rexTypes.AllowCreateModelCredential) (code int32, result *rexTypes.CredentialApiCreateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CredentialApiCreateResp]{}
	code, err = m.do(ctx, "/credentials/credentialConfig/create", http.MethodPost, &params, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCredentialService) Delete(ctx context.Context, params *rexTypes.CredentialApiFormIdReq) (code int32, result *rexTypes.CredentialApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CredentialApiOKResp]{}
	code, err = m.do(ctx, fmt.Sprintf("/credentials/credentialConfig/delete?id=%d", params.Id), http.MethodPost, nil, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCredentialService) DeleteMany(ctx context.Context, params *rexTypes.CredentialApiJsonIdsReq) (code int32, result *rexTypes.CredentialApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CredentialApiOKResp]{}
	code, err = m.do(ctx, "/credentials/credentialConfig/deleteMany", http.MethodPost, &params, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCredentialService) Update(ctx context.Context, params *rexTypes.AllowUpdateModelCredential) (code int32, result *rexTypes.CredentialApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CredentialApiOKResp]{}
	code, err = m.do(ctx, "/credentials/credentialConfig/update", http.MethodPost, &params, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCredentialService) UpdateStatus(ctx context.Context, params *rexTypes.AllowUpdateStatusModelCredential) (code int32, result *rexTypes.CredentialApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CredentialApiOKResp]{}
	code, err = m.do(ctx, "/credentials/credentialConfig/updateStatus", http.MethodPost, &params, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCredentialService) UpdateDefault(ctx context.Context, params *rexTypes.AllowUpdateDefaultModelCredential) (code int32, result *rexTypes.CredentialApiOKResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CredentialApiOKResp]{}
	code, err = m.do(ctx, "/credentials/credentialConfig/updateDefault", http.MethodPost, &params, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCredentialService) Query(ctx context.Context, params *rexTypes.CredentialApiFormIdReq) (code int32, result *rexTypes.ModelCredential, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ModelCredential]{}
	code, err = m.do(ctx, fmt.Sprintf("/credentials/credentialConfig/query?id=%d", params.Id), http.MethodGet, nil, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCredentialService) QueryListWhereIds(ctx context.Context, params *rexTypes.CredentialApiJsonIdsReq) (code int32, result *rexTypes.CredentialCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CredentialCommonQueryListResp]{}
	code, err = m.do(ctx, "/credentials/credentialConfig/queryListWhereIds", http.MethodPost, &params, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultCredentialService) QueryList(ctx context.Context, params *rexTypes.CredentialCommonSearchParams) (code int32, result *rexTypes.CredentialCommonQueryListResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CredentialCommonQueryListResp]{}
	code, err = m.do(ctx, "/credentials/credentialConfig/queryList", http.MethodPost, &params, tmp)
	if err != nil {
		return code, nil, err
	}
	if tmp.Code != rexCodes.OK {
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
