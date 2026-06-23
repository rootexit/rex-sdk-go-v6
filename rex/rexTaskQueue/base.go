package rexTaskQueue

import sdkCtx "github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"

type (
	BaseService interface {
	}

	defaultBaseService struct {
		Svc    string
		SdkCtx *sdkCtx.SdkCtx
	}
)

func NewBaseService(SdkCtx *sdkCtx.SdkCtx) BaseService {
	return &defaultBaseService{
		Svc:    "taskQueue",
		SdkCtx: SdkCtx,
	}
}
