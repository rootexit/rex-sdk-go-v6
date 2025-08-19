package sdkCtx

import (
	"context"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexClient"
)

type SdkCtx struct {
	Context context.Context
	Cancel  context.CancelFunc
	Cli     *rexClient.Client
}

func NewSdkCtx(client *rexClient.Client) *SdkCtx {
	ctx, cancel := context.WithCancel(context.Background())
	return &SdkCtx{
		Context: ctx,
		Cancel:  cancel,
		Cli:     client,
	}
}
