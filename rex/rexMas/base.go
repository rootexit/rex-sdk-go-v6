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
	MasBaseService interface {
		// note: 生成验证码
		CaptchaGenerate(ctx context.Context, params *rexTypes.ApiCaptchaGenerateReq) (result *rexTypes.ApiCaptchaGenerateResp, err error)
		SmsSend(ctx context.Context, params *rexTypes.ApiSmsSendReq) (result *rexTypes.ApiSmsSendResp, err error)

		BehavioralVerificationInit(ctx context.Context, params *rexTypes.BehavioralVerificationInitReq) (result *rexTypes.BehavioralVerificationInitResp, err error)
		BehavioralVerificationVerify(ctx context.Context, params *rexTypes.BehavioralVerificationVerifyReq) (result *rexTypes.BehavioralVerificationVerifyResp, err error)
		SmsVerificationInit(ctx context.Context, params *rexTypes.SmsInitReq) (result *rexTypes.SmsInitResp, err error)
		SmsVerificationVerify(ctx context.Context, params *rexTypes.SmsVerifyReq) (result *rexTypes.SmsVerifyResp, err error)
	}

	defaultMasBaseService struct {
		Svc    string
		rexCtx *rexCtx.QxCtx
	}
)

func NewMsgBaseService(rexCtx *rexCtx.QxCtx) MasBaseService {
	return &defaultMasBaseService{
		Svc:    "mas",
		rexCtx: rexCtx,
	}
}

func (m *defaultMasBaseService) CaptchaGenerate(ctx context.Context, params *rexTypes.ApiCaptchaGenerateReq) (result *rexTypes.ApiCaptchaGenerateResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ApiCaptchaGenerateResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/captcha/generate", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: captcha generate fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) SmsSend(ctx context.Context, params *rexTypes.ApiSmsSendReq) (result *rexTypes.ApiSmsSendResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.ApiSmsSendResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/send", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: sms send fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) BehavioralVerificationInit(ctx context.Context, params *rexTypes.BehavioralVerificationInitReq) (result *rexTypes.BehavioralVerificationInitResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BehavioralVerificationInitResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/bv/init", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk:Behavioral Verification Init fail: %v", result)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) BehavioralVerificationVerify(ctx context.Context, params *rexTypes.BehavioralVerificationVerifyReq) (result *rexTypes.BehavioralVerificationVerifyResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.BehavioralVerificationVerifyResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/bv/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: Behavioral Verification Verify fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) SmsVerificationInit(ctx context.Context, params *rexTypes.SmsInitReq) (result *rexTypes.SmsInitResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsInitResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/init", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: sms init fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) SmsVerificationVerify(ctx context.Context, params *rexTypes.SmsVerifyReq) (result *rexTypes.SmsVerifyResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SmsVerifyResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: sms verify fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
