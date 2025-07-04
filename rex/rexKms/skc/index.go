package skc

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
	KmsSkcService interface {
		KmsSkcCreateKeychain(ctx context.Context, params *rexTypes.KmsSkcCreateKeychainReq) (result *rexTypes.KmsSkcCreateKeychainResp, err error)
		KmsSkcEncrypt(ctx context.Context, params *rexTypes.KmsSkcEncryptReq) (result *rexTypes.KmsSkcEncryptResp, err error)
		KmsSkcDecrypt(ctx context.Context, params *rexTypes.KmsSkcDecryptReq) (result *rexTypes.KmsSkcDecryptResp, err error)
		KmsSkcBatchEncrypt(ctx context.Context, params *rexTypes.KmsSkcBatchEncryptReq) (result *rexTypes.KmsSkcBatchEncryptResp, err error)
		KmsSkcBatchDecrypt(ctx context.Context, params *rexTypes.KmsSkcBatchDecryptReq) (result *rexTypes.KmsSkcBatchDecryptResp, err error)
		KmsSkcCompare(ctx context.Context, params *rexTypes.KmsSkcCompareReq) (result *rexTypes.KmsSkcCompareResp, err error)
	}

	defaultKmsSkcService struct {
		Svc    string
		rexCtx *rexCtx.QxCtx
	}
)

func NewKmsSkcService(rexCtx *rexCtx.QxCtx) KmsSkcService {
	// note: 初始化Kms系统
	return &defaultKmsSkcService{
		Svc:    "kms",
		rexCtx: rexCtx,
	}
}

func (m *defaultKmsSkcService) KmsSkcCreateKeychain(ctx context.Context, params *rexTypes.KmsSkcCreateKeychainReq) (result *rexTypes.KmsSkcCreateKeychainResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcCreateKeychainResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/createKeychain", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcEncrypt(ctx context.Context, params *rexTypes.KmsSkcEncryptReq) (result *rexTypes.KmsSkcEncryptResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcEncryptResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/encrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcDecrypt(ctx context.Context, params *rexTypes.KmsSkcDecryptReq) (result *rexTypes.KmsSkcDecryptResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcDecryptResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/decrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: KmsSkcDecrypt fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcBatchEncrypt(ctx context.Context, params *rexTypes.KmsSkcBatchEncryptReq) (result *rexTypes.KmsSkcBatchEncryptResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcBatchEncryptResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/batchEncrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcBatchDecrypt(ctx context.Context, params *rexTypes.KmsSkcBatchDecryptReq) (result *rexTypes.KmsSkcBatchDecryptResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcBatchDecryptResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/batchDecrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcCompare(ctx context.Context, params *rexTypes.KmsSkcCompareReq) (result *rexTypes.KmsSkcCompareResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcCompareResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/compare", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
