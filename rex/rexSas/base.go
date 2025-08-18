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
	BaseService interface {
		QueryBucket(ctx context.Context, params *rexTypes.SasQueryBucketReq) (code int32, result *rexTypes.SasQueryBucketResp, err error)
		PresignerGet(ctx context.Context, params *rexTypes.SasPresignerGetObjectReq) (code int32, result *rexTypes.SasPresignerGetObjectResp, err error)
		PresignerUpload(ctx context.Context, params *rexTypes.SasPresignerUploadReq) (code int32, result *rexTypes.SasPresignerUploadResp, err error)
		PresignerHeadObject(ctx context.Context, params *rexTypes.SasPresignerHeadObjectReq) (code int32, result *rexTypes.SasPresignerHeadObjectResp, err error)
		CreateBucketAndConfig(ctx context.Context, params *rexTypes.CreateExistBucketAndConfigReq) (code int32, result *rexTypes.CreateExistBucketAndConfigResp, err error)
		CreateBucketNoConfig(ctx context.Context, params *rexTypes.CreateExistBucketNoConfigReq) (code int32, result *rexTypes.CreateExistBucketNoConfigResp, err error)
	}
	defaultBaseService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewBaseService(rexCtx *rexCtx.EngineCtx) BaseService {
	return &defaultBaseService{
		Svc:    "sas",
		rexCtx: rexCtx,
	}
}

// @doc "预签名访问"
func (m *defaultBaseService) PresignerGet(ctx context.Context, params *rexTypes.SasPresignerGetObjectReq) (code int32, result *rexTypes.SasPresignerGetObjectResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SasPresignerGetObjectResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerGetObject", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request sas:BaseService:PresignerGet error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request sas:BaseService:PresignerGet fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) QueryBucket(ctx context.Context, params *rexTypes.SasQueryBucketReq) (code int32, result *rexTypes.SasQueryBucketResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SasQueryBucketResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/queryBucket", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request sas:BaseService:QueryBucket error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request sas:BaseService:QueryBucket fail: %s", res)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) PresignerUpload(ctx context.Context, params *rexTypes.SasPresignerUploadReq) (code int32, result *rexTypes.SasPresignerUploadResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SasPresignerUploadResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerUpload", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request sas:BaseService:PresignerUpload error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request sas:BaseService:PresignerUpload fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) PresignerHeadObject(ctx context.Context, params *rexTypes.SasPresignerHeadObjectReq) (code int32, result *rexTypes.SasPresignerHeadObjectResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.SasPresignerHeadObjectResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerHeadObject", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request sas:BaseService:PresignerHeadObject  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request sas:BaseService:PresignerHeadObject fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) CreateBucketAndConfig(ctx context.Context, params *rexTypes.CreateExistBucketAndConfigReq) (code int32, result *rexTypes.CreateExistBucketAndConfigResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CreateExistBucketAndConfigResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/createExistBucketAndConfig", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: requestsas:BaseService:CreateBucketAndConfig  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request sas:BaseService:CreateBucketAndConfig  fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultBaseService) CreateBucketNoConfig(ctx context.Context, params *rexTypes.CreateExistBucketNoConfigReq) (code int32, result *rexTypes.CreateExistBucketNoConfigResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CreateExistBucketNoConfigResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/createExistBucketNoConfig", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request sas:BaseService:CreateBucketNoConfig  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request sas:BaseService:CreateBucketNoConfig  fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
