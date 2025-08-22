package rexSigner

import (
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
	"github.com/rootexit/rexLib/rexCustomAwsSign"
	"github.com/rootexit/rexLib/rexHeaders"
	"net/http"
	"time"
)

type Signer struct {
	// Access key ID
	AccessKeyID string `json:",optional,inherit"`
	// Secret Access Key
	AccessKeySecret string `json:",optional,inherit"`
	Region          string `json:",default=cn-shanghai,optional"`
	signer          rexCustomAwsSign.CustomSigner
}

func NewSigner(conf *rexConfig.Config) *Signer {
	region := "cn-shanghai"
	if conf.Region == "" {
		region = "cn-shanghai"
	} else {
		region = conf.Region
	}
	s := rexCustomAwsSign.NewCustomSigner("REx", 1)
	if conf.Debug {
		s.WithDebug(true)
	}
	return &Signer{
		AccessKeyID:     conf.AccessKeyID,
		AccessKeySecret: conf.AccessKeySecret,
		Region:          region,
		signer:          s,
	}
}

func (s *Signer) Sign(req *http.Request, bodyBytes []byte, svc string, reqTime time.Time) error {
	hexContentSha256 := s.signer.Sha256Content(bodyBytes)
	req.Header.Add(s.signer.GetHeaderContentSha256(), hexContentSha256)
	timeStr := s.signer.FormatDate(reqTime)
	req.Header.Add(s.signer.GetHeaderDate(), timeStr)
	canonicalHeaders, signedHeadersStr := s.signer.BuildCanonicalHeaders(req)
	canonicalString := s.signer.BuildCanonicalString(req, canonicalHeaders, signedHeadersStr, hexContentSha256)
	credentialString := s.signer.BuildCredentialString(s.Region, svc, reqTime)
	stringToSign := s.signer.BuildStringToSign(timeStr, credentialString, canonicalString)
	signature := s.signer.BuildSignature(s.Region, svc, s.AccessKeySecret, stringToSign, reqTime)
	auth := s.signer.SignAuth(s.AccessKeyID, credentialString, signedHeadersStr, signature)
	req.Header.Set(rexHeaders.HeaderAuthorization, auth)
	return nil
}
