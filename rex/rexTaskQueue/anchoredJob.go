package rexTaskQueue

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	sdkCtx "github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
	"github.com/rootexit/rexLib/rexCodes"
	"github.com/rootexit/rexLib/rexRes"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	AnchoredJobService interface {
		Add(ctx context.Context, params *rexTypes.TaskQueueAnchoredJobAddReq) (code int32, result *rexTypes.TaskQueueAnchoredJobAddResp, err error)
		Remove(ctx context.Context, params *rexTypes.TaskQueueAnchoredJobRemoveReq) (code int32, result *rexTypes.TaskQueueAnchoredJobRemoveResp, err error)
	}

	defaultAnchoredJobService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewAnchoredJobService(SdkCtx *sdkCtx.SdkCtx) AnchoredJobService {
	return &defaultAnchoredJobService{
		Svc:    "taskQueue",
		SdkCtx: SdkCtx,
	}
}

func (m *defaultAnchoredJobService) Add(ctx context.Context, params *rexTypes.TaskQueueAnchoredJobAddReq) (code int32, result *rexTypes.TaskQueueAnchoredJobAddResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TaskQueueAnchoredJobAddResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/taskQueue/anchoredJob/add", http.MethodPost, &params)
	if err != nil {
		logx.Errorf("rex sdk: request taskQueue:AnchoredJobService:Add error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request taskQueue:AnchoredJobService:Add fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultAnchoredJobService) Remove(ctx context.Context, params *rexTypes.TaskQueueAnchoredJobRemoveReq) (code int32, result *rexTypes.TaskQueueAnchoredJobRemoveResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TaskQueueAnchoredJobRemoveResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/taskQueue/anchoredJob/remove", http.MethodPost, &params)
	if err != nil {
		logx.Errorf("rex sdk: request taskQueue:AnchoredJobService:Remove error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request taskQueue:AnchoredJobService:Remove fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
