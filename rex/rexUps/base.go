package rexUps

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
		Bootstrap(ctx context.Context, params *rexTypes.UpsBaseBootstrapReq) (result *rexTypes.UpsBaseBootstrapResp, err error)
	}
	defaultBaseService struct {
		Svc    string
		rexCtx *rexCtx.EngineCtx
	}
)

func NewBaseService(rexCtx *rexCtx.EngineCtx) BaseService {
	return &defaultBaseService{
		Svc:    "ups",
		rexCtx: rexCtx,
	}
}

func (m *defaultBaseService) Bootstrap(ctx context.Context, params *rexTypes.UpsBaseBootstrapReq) (result *rexTypes.UpsBaseBootstrapResp, err error) {
	tmp := &rexRes.BaseResponse[rexTypes.UpsBaseBootstrapResp]{}
	res, err := m.rexCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/base/bootstrap", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("rex sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != rexCodes.EngineStatusOK {
		logx.Errorf("rex sdk: request ups:BaseService:Bootstrap fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
