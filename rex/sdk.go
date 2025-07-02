package rex

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexBase"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexClient"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexKms"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexMas"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexSas"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexUps"
)

type QxSdk struct {
	config        *rexConfig.Config
	client        *rexClient.QxClient
	rexCtx        *rexCtx.QxCtx
	QxBaseService rexBase.QxBaseService
	KmsService    rexKms.KmsService
	MasService    rexMas.MasService
	SasService    rexSas.SasService
	UpsService    rexUps.UpsService
}

func NewQxSdk(c *rexConfig.Config) (*QxSdk, error) {
	// note: sdk初始化检查
	if err := c.Check(); err != nil {
		return nil, err
	}
	// note: sdk初始化
	client := rexClient.NewQxClient(c)
	rexC := rexCtx.NewQxCtx(client)

	sdk := &QxSdk{
		config:        c,
		client:        client,
		rexCtx:        rexC,
		QxBaseService: rexBase.NewQxBaseService(rexC),
		KmsService:    rexKms.NewKmsService(rexC),
		MasService:    rexMas.NewMasService(rexC),
		SasService:    rexSas.NewSasService(rexC),
		UpsService:    rexUps.NewUpsService(rexC),
	}
	return sdk, nil
}
