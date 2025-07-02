package rexKms

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexCtx"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexKms/akc"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexKms/skc"
)

type (
	KmsService struct {
		Akc akc.KmsAkcService
		Skc skc.KmsSkcService
	}
)

func NewKmsService(rexCtx *rexCtx.QxCtx) KmsService {
	return KmsService{
		Akc: akc.NewKmsAkcService(rexCtx),
		Skc: skc.NewKmsSkcService(rexCtx),
	}
}
