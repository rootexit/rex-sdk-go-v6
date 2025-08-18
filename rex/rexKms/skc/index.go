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
	SkcService interface {
		CreateKeychain(ctx context.Context, params *rexTypes.KmsSkcCreateKeychainReq) (code int32, result *rexTypes.KmsSkcCreateKeychainResp, err error)
		Encrypt(ctx context.Context, params *rexTypes.KmsSkcEncryptReq) (code int32, result *rexTypes.KmsSkcEncryptResp, err error)
		Decrypt(ctx context.Context, params *rexTypes.KmsSkcDecryptReq) (code int32, result *rexTypes.KmsSkcDecryptResp, err error)
		BatchEncrypt(ctx context.Context, params *rexTypes.KmsSkcBatchEncryptReq) (code int32, result *rexTypes.KmsSkcBatchEncryptResp, err error)
		BatchDecrypt(ctx context.Context, params *rexTypes.KmsSkcBatchDecryptReq) (code int32, result *rexTypes.KmsSkcBatchDecryptResp, err error)
		Compare(ctx context.Context, params *rexTypes.KmsSkcCompareReq) (code int32, result *rexTypes.KmsSkcCompareResp, err error)
	}

	defaultSkcService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewSkcService(rexCtx *rexCtx.EngineCtx) SkcService {
	// note: 初始化Kms系统
	return &defaultSkcService{
		Svc:    "kms",
		rexCtx: rexCtx,
	}
}

func (m *defaultSkcService) CreateKeychain(ctx context.Context, params *rexTypes.KmsSkcCreateKeychainReq) (code int32, result *rexTypes.KmsSkcCreateKeychainResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcCreateKeychainResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/createKeychain", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:SkcService:CreateKeychain error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:SkcService:CreateKeychain fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSkcService) Encrypt(ctx context.Context, params *rexTypes.KmsSkcEncryptReq) (code int32, result *rexTypes.KmsSkcEncryptResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcEncryptResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/encrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:SkcService:Encrypt error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:SkcService:Encrypt fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSkcService) Decrypt(ctx context.Context, params *rexTypes.KmsSkcDecryptReq) (code int32, result *rexTypes.KmsSkcDecryptResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcDecryptResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/decrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:SkcService:Decrypt error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:SkcService:Decrypt fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSkcService) BatchEncrypt(ctx context.Context, params *rexTypes.KmsSkcBatchEncryptReq) (code int32, result *rexTypes.KmsSkcBatchEncryptResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcBatchEncryptResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/batchEncrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:SkcService:BatchEncrypt error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:SkcService:BatchEncrypt fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSkcService) BatchDecrypt(ctx context.Context, params *rexTypes.KmsSkcBatchDecryptReq) (code int32, result *rexTypes.KmsSkcBatchDecryptResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcBatchDecryptResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/batchDecrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:SkcService:BatchDecrypt error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:SkcService:BatchDecrypt fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultSkcService) Compare(ctx context.Context, params *rexTypes.KmsSkcCompareReq) (code int32, result *rexTypes.KmsSkcCompareResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.KmsSkcCompareResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/compare", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request kms:SkcService:Compare error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request kms:SkcService:Compare fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
