package rexTaskQueue

import (
	sdkCtx "github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
)

type (
	TaskQueueService struct {
		BaseService
		PeriodicJobService PeriodicJobService
		AnchoredJobService AnchoredJobService
		DelayedJobService  DelayedJobService
	}
)

func NewTaskQueueService(SdkCtx *sdkCtx.SdkCtx) TaskQueueService {
	return TaskQueueService{
		BaseService:        NewBaseService(SdkCtx),
		PeriodicJobService: NewPeriodicJobService(SdkCtx),
		AnchoredJobService: NewAnchoredJobService(SdkCtx),
		DelayedJobService:  NewDelayedJobService(SdkCtx),
	}
}
