package rex

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexBase"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexClient"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCredentials"
	sdkCtx "github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexKms"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexMas"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexSas"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTaskQueue"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTpas"
)

type Sdk struct {
	config            *rexConfig.Config
	client            *rexClient.Client
	rexCtx            *sdkCtx.SdkCtx
	BaseService       rexBase.BaseService
	CredentialService rexCredentials.CredentialService
	KmsService        rexKms.KmsService
	MasService        rexMas.MasService
	SasService        rexSas.SasService
	TpasService       rexTpas.TpasService
	TaskQueueService  rexTaskQueue.TaskQueueService
}

func NewSdk(c *rexConfig.Config) (*Sdk, error) {
	// note: sdk初始化检查
	if err := c.Check(); err != nil {
		return nil, err
	}
	// note: sdk初始化
	client := rexClient.NewClient(c)
	rexC := sdkCtx.NewSdkCtx(client)

	sdk := &Sdk{
		config:            c,
		client:            client,
		rexCtx:            rexC,
		BaseService:       rexBase.NewBaseService(rexC),
		CredentialService: rexCredentials.NewCredentialService(rexC),
		KmsService:        rexKms.NewKmsService(rexC),
		MasService:        rexMas.NewMasService(rexC),
		SasService:        rexSas.NewSasService(rexC),
		TpasService:       rexTpas.NewTpasService(rexC),
		TaskQueueService:  rexTaskQueue.NewTaskQueueService(rexC),
	}
	return sdk, nil
}
