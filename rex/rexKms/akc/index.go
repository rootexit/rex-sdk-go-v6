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
	KmsAkcService interface {
		KmsAkcCreateKeychain(ctx context.Context, params *rexTypes.KmsAkcCreateKeychainReq) (result *rexTypes.KmsAkcCreateKeychainResp, err error)
		KmsAkcSign(ctx context.Context, params *rexTypes.KmsAkcSignReq) (result *rexTypes.KmsAkcSignResp, err error)
		KmsAkcVerify(ctx context.Context, params *rexTypes.KmsAkcVerifyReq) (result *rexTypes.KmsAkcVerifyResp, err error)
	}

	defaultKmsAkcService struct {
		Svc    string
		rexCtx *rexCtx.QxCtx
	}
)

func NewKmsAkcService(rexCtx *rexCtx.QxCtx) KmsAkcService {
	// note: 初始化Kms系统
	return &defaultKmsAkcService{
		Svc:    "kms",
		rexCtx: rexCtx,
	}
}

func (m *defaultKmsAkcService) KmsAkcCreateKeychain(ctx context.Context, params *rexTypes.KmsAkcCreateKeychainReq) (result *rexTypes.KmsAkcCreateKeychainResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsAkcCreateKeychainResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/createKeychain", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: KmsAkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultKmsAkcService) KmsAkcSign(ctx context.Context, params *rexTypes.KmsAkcSignReq) (result *rexTypes.KmsAkcSignResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsAkcSignResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/sign", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: KmsAkcSign fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultKmsAkcService) KmsAkcVerify(ctx context.Context, params *rexTypes.KmsAkcVerifyReq) (result *rexTypes.KmsAkcVerifyResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsAkcVerifyResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: KmsAkcVerify fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
