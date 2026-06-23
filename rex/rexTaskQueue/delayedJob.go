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
	DelayedJobService interface {
		Add(ctx context.Context, params *rexTypes.TaskQueueDelayedJobAddReq) (code int32, result *rexTypes.TaskQueueDelayedJobAddResp, err error)
		Cancel(ctx context.Context, params *rexTypes.TaskQueueDelayedJobCancelReq) (code int32, result *rexTypes.TaskQueueDelayedJobCancelResp, err error)
	}

	defaultDelayedJobService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewDelayedJobService(SdkCtx *sdkCtx.SdkCtx) DelayedJobService {
	return &defaultDelayedJobService{
		Svc:    "taskQueue",
		SdkCtx: SdkCtx,
	}
}

func (m *defaultDelayedJobService) Add(ctx context.Context, params *rexTypes.TaskQueueDelayedJobAddReq) (code int32, result *rexTypes.TaskQueueDelayedJobAddResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TaskQueueDelayedJobAddResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/taskQueue/delayedJob/add", http.MethodPost, &params)
	if err != nil {
		logx.Errorf("rex sdk: request taskQueue:DelayedJobService:Add error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request taskQueue:DelayedJobService:Add fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultDelayedJobService) Cancel(ctx context.Context, params *rexTypes.TaskQueueDelayedJobCancelReq) (code int32, result *rexTypes.TaskQueueDelayedJobCancelResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TaskQueueDelayedJobCancelResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/taskQueue/delayedJob/cancel", http.MethodPost, &params)
	if err != nil {
		logx.Errorf("rex sdk: request taskQueue:DelayedJobService:Cancel error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request taskQueue:DelayedJobService:Cancel fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
