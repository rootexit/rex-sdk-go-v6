package rexKms

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexKms/akc"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexKms/skc"
)

type (
	KmsService struct {
		Akc akc.AkcService
		Skc skc.SkcService
	}
)

func NewKmsService(SdkCtx *sdkCtx.SdkCtx) KmsService {
	return KmsService{
		Akc: akc.NewAkcService(SdkCtx),
		Skc: skc.NewSkcService(SdkCtx),
	}
}
