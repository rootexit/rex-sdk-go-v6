package rexCtas

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
	PeriodicJobService interface {
		// note: 为某个公众号强制刷新凭证
		Add(ctx context.Context, params *rexTypes.CtasPeriodicJobAddReq) (code int32, result *rexTypes.CtasPeriodicJobAddResp, err error)
		// note: 获取公众号普通AccessToken
		Remove(ctx context.Context, params *rexTypes.CtasPeriodicJobRemoveReq) (code int32, result *rexTypes.CtasPeriodicJobRemoveResp, err error)
	}

	defaultPeriodicJobService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewPeriodicJobService(rexCtx *rexCtx.EngineCtx) PeriodicJobService {
	return &defaultPeriodicJobService{
		Svc:    "ctas",
		rexCtx: rexCtx,
	}
}

func (m *defaultPeriodicJobService) Add(ctx context.Context, params *rexTypes.CtasPeriodicJobAddReq) (code int32, result *rexTypes.CtasPeriodicJobAddResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CtasPeriodicJobAddResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ctas/periodicJob/add", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ctas:PeriodicJobService:Add  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ctas:PeriodicJobService:Add fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}

func (m *defaultPeriodicJobService) Remove(ctx context.Context, params *rexTypes.CtasPeriodicJobRemoveReq) (code int32, result *rexTypes.CtasPeriodicJobRemoveResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.CtasPeriodicJobRemoveResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ctas/periodicJob/remove", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request ctas:PeriodicJobService:Remove  error: %v", err)
		return rexCodes.FAIL, nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.OK {
		logx.Errorf("rex sdk: request ctas:PeriodicJobService:Remove fail: %v", tmp)
		return tmp.Code, &tmp.Data, errors.New(tmp.Msg)
	}
	return rexCodes.OK, &tmp.Data, nil
}
