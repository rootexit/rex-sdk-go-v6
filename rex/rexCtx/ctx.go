package rexCtx

import (
	"context"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexClient"
)

type QxCtx struct {
	Context context.Context
	Cancel  context.CancelFunc
	Cli     *rexClient.QxClient
}

func NewQxCtx(client *rexClient.QxClient) *QxCtx {
	ctx, cancel := context.WithCancel(context.Background())
	return &QxCtx{
		Context: ctx,
		Cancel:  cancel,
		Cli:     client,
	}
}
