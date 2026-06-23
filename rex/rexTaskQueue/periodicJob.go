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
	PeriodicJobService interface {
		Add(ctx context.Context, params *rexTypes.TaskQueuePeriodicJobAddReq) (code int32, result *rexTypes.TaskQueuePeriodicJobAddResp, err error)
		Remove(ctx context.Context, params *rexTypes.TaskQueuePeriodicJobRemoveReq) (code int32, result *rexTypes.TaskQueuePeriodicJobRemoveResp, err error)
	}

	defaultPeriodicJobService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewPeriodicJobService(SdkCtx *sdkCtx.SdkCtx) PeriodicJobService {
	return &defaultPeriodicJobService{
		Svc:    "taskQueue",
		SdkCtx: SdkCtx,
	}
}

func (m *defaultPeriodicJobService) Add(ctx context.Context, params *rexTypes.TaskQueuePeriodicJobAddReq) (code int32, result *rexTypes.TaskQueuePeriodicJobAddResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TaskQueuePeriodicJobAddResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/taskQueue/periodicJob/add", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request taskQueue:PeriodicJobService:Add error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request taskQueue:PeriodicJobService:Add fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultPeriodicJobService) Remove(ctx context.Context, params *rexTypes.TaskQueuePeriodicJobRemoveReq) (code int32, result *rexTypes.TaskQueuePeriodicJobRemoveResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.TaskQueuePeriodicJobRemoveResp]{}
	res, err := m.SdkCtx.Cli.EasyNewRequest(ctx, m.Svc, "/taskQueue/periodicJob/remove", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request taskQueue:PeriodicJobService:Remove error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request taskQueue:PeriodicJobService:Remove fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
