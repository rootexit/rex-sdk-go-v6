package akc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	AkcService interface {
		GetPublicKey(ctx context.Context, params *rexTypes.KmsAkcGetKeychainPublicKeyReq) (code int32, result *rexTypes.KmsAkcGetKeychainPublicKeyResp, err error)
		CreateKeychain(ctx context.Context, params *rexTypes.KmsAkcCreateKeychainReq) (code int32, result *rexTypes.KmsAkcCreateKeychainResp, err error)
		Sign(ctx context.Context, params *rexTypes.KmsAkcSignReq) (code int32, result *rexTypes.KmsAkcSignResp, err error)
		Verify(ctx context.Context, params *rexTypes.KmsAkcVerifyReq) (code int32, result *rexTypes.KmsAkcVerifyResp, err error)
	}

	defaultAkcService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewAkcService(SdkCtx *sdkCtx.SdkCtx) AkcService {
	// note: 初始化Kms系统
	return &defaultAkcService{
		Svc:    "kms",
		SdkCtx: SdkCtx,
	}
}

func (m *defaultAkcService) GetPublicKey(ctx context.Context, params *rexTypes.KmsAkcGetKeychainPublicKeyReq) (code int32, result *rexTypes.KmsAkcGetKeychainPublicKeyResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsAkcGetKeychainPublicKeyResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/getKeychainPublicKey", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:AkcService:CreateKeychain error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:AkcService:CreateKeychain fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultAkcService) CreateKeychain(ctx context.Context, params *rexTypes.KmsAkcCreateKeychainReq) (code int32, result *rexTypes.KmsAkcCreateKeychainResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsAkcCreateKeychainResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/createKeychain", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:AkcService:CreateKeychain error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:AkcService:CreateKeychain fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultAkcService) Sign(ctx context.Context, params *rexTypes.KmsAkcSignReq) (code int32, result *rexTypes.KmsAkcSignResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsAkcSignResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/sign", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:AkcService:Sign error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:AkcService:Sign fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultAkcService) Verify(ctx context.Context, params *rexTypes.KmsAkcVerifyReq) (code int32, result *rexTypes.KmsAkcVerifyResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsAkcVerifyResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:AkcService:Verify  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:AkcService:Verify fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
