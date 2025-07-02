package rexSas

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
	SasBaseService interface {
		QueryBucket(ctx context.Context, params *rexTypes.SasQueryBucketReq) (result *rexTypes.SasQueryBucketResp, err error)
		PresignerGet(ctx context.Context, params *rexTypes.SasPresignerGetObjectReq) (result *rexTypes.SasPresignerGetObjectResp, err error)
		PresignerUpload(ctx context.Context, params *rexTypes.SasPresignerUploadReq) (result *rexTypes.SasPresignerUploadResp, err error)
		PresignerHeadObject(ctx context.Context, params *rexTypes.SasPresignerHeadObjectReq) (result *rexTypes.SasPresignerHeadObjectResp, err error)
		CreateBucketAndConfig(ctx context.Context, params *rexTypes.CreateExistBucketAndConfigReq) (result *rexTypes.CreateExistBucketAndConfigResp, err error)
		CreateBucketNoConfig(ctx context.Context, params *rexTypes.CreateExistBucketNoConfigReq) (result *rexTypes.CreateExistBucketNoConfigResp, err error)
	}
	defaultSasBaseService struct {
		Svc    string
		rexCtx *rexCtx.QxCtx
	}
)

func NewSasBaseService(rexCtx *rexCtx.QxCtx) SasBaseService {
	return &defaultSasBaseService{
		Svc:    "sas",
		rexCtx: rexCtx,
	}
}

// @doc "预签名访问"
func (m *defaultSasBaseService) PresignerGet(ctx context.Context, params *rexTypes.SasPresignerGetObjectReq) (result *rexTypes.SasPresignerGetObjectResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SasPresignerGetObjectResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerGetObject", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) QueryBucket(ctx context.Context, params *rexTypes.SasQueryBucketReq) (result *rexTypes.SasQueryBucketResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SasQueryBucketResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/queryBucket", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) PresignerUpload(ctx context.Context, params *rexTypes.SasPresignerUploadReq) (result *rexTypes.SasPresignerUploadResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SasPresignerUploadResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerUpload", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) PresignerHeadObject(ctx context.Context, params *rexTypes.SasPresignerHeadObjectReq) (result *rexTypes.SasPresignerHeadObjectResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SasPresignerHeadObjectResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerHeadObject", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) CreateBucketAndConfig(ctx context.Context, params *rexTypes.CreateExistBucketAndConfigReq) (result *rexTypes.CreateExistBucketAndConfigResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CreateExistBucketAndConfigResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/createExistBucketAndConfig", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) CreateBucketNoConfig(ctx context.Context, params *rexTypes.CreateExistBucketNoConfigReq) (result *rexTypes.CreateExistBucketNoConfigResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CreateExistBucketNoConfigResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/createExistBucketNoConfig", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.QxEngineStatusOK {
		logx.Errorf("rex sdk: request fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
