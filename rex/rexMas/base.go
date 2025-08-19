package rexMas

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
	BaseService interface {
		// note: 生成验证码
		CaptchaGenerate(ctx context.Context, params *rexTypes.ApiCaptchaGenerateReq) (code int32, result *rexTypes.ApiCaptchaGenerateResp, err error)
		SmsSend(ctx context.Context, params *rexTypes.ApiSmsSendReq) (code int32, result *rexTypes.ApiSmsSendResp, err error)

		BehavioralVerificationInit(ctx context.Context, params *rexTypes.BehavioralVerificationInitReq) (code int32, result *rexTypes.BehavioralVerificationInitResp, err error)
		BehavioralVerificationVerify(ctx context.Context, params *rexTypes.BehavioralVerificationVerifyReq) (code int32, result *rexTypes.BehavioralVerificationVerifyResp, err error)
		SmsVerificationInit(ctx context.Context, params *rexTypes.SmsInitReq) (code int32, result *rexTypes.SmsInitResp, err error)
		SmsVerificationVerify(ctx context.Context, params *rexTypes.SmsVerifyReq) (code int32, result *rexTypes.SmsVerifyResp, err error)
	}

	defaultBaseService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewBaseService(SdkCtx *sdkCtx.SdkCtx) BaseService {
	return &defaultBaseService{
		Svc:    "mas",
		SdkCtx: SdkCtx,
	}
}

func (m *defaultBaseService) CaptchaGenerate(ctx context.Context, params *rexTypes.ApiCaptchaGenerateReq) (code int32, result *rexTypes.ApiCaptchaGenerateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ApiCaptchaGenerateResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/captcha/generate", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request mas:BaseService:CaptchaGenerate  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request mas:BaseService:CaptchaGenerate fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) SmsSend(ctx context.Context, params *rexTypes.ApiSmsSendReq) (code int32, result *rexTypes.ApiSmsSendResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ApiSmsSendResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/send", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request mas:BaseService:SmsSend  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request mas:BaseService:SmsSend fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) BehavioralVerificationInit(ctx context.Context, params *rexTypes.BehavioralVerificationInitReq) (code int32, result *rexTypes.BehavioralVerificationInitResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BehavioralVerificationInitResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/bv/init", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request mas:BaseService:BehavioralVerificationInit  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request mas:BaseService:BehavioralVerificationInit fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) BehavioralVerificationVerify(ctx context.Context, params *rexTypes.BehavioralVerificationVerifyReq) (code int32, result *rexTypes.BehavioralVerificationVerifyResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BehavioralVerificationVerifyResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/bv/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request mas:BaseService:BehavioralVerificationVerify  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request mas:BaseService:BehavioralVerificationVerify fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) SmsVerificationInit(ctx context.Context, params *rexTypes.SmsInitReq) (code int32, result *rexTypes.SmsInitResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsInitResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/init", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request mas:BaseService:SmsVerificationInit  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request mas:BaseService:SmsVerificationInit fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) SmsVerificationVerify(ctx context.Context, params *rexTypes.SmsVerifyReq) (code int32, result *rexTypes.SmsVerifyResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsVerifyResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request mas:BaseService:SmsVerificationVerify  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request mas:BaseService:SmsVerificationVerify fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
