package rexCtx

import (
	"context"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexClient"
)

type EngineCtx struct {
	Context context.Context
	Cancel  context.CancelFunc
	Cli     *rexClient.QxClient
}

func NewEngineCtx(client *rexClient.QxClient) *EngineCtx {
	ctx, cancel := context.WithCancel(context.Background())
	return &EngineCtx{
		Context: ctx,
		Cancel:  cancel,
		Cli:     client,
	}
}
