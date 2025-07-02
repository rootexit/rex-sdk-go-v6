package rexConfig

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	defaultTimeout  = 2000
	defaultProtocol = "https"
	defaultRegion   = "cn-shanghai"
	defaultEndpoint = "core.csvw88.com"
)

type Config struct {
	// Access key ID
	AccessKeyID string `json:",optional,inherit"`
	// Secret Access Key
	AccessKeySecret string `json:",optional,inherit"`
	Endpoint        string `json:",default=core.csvw88.com,optional"`
	Protocol        string `json:",default=https,options=http|https,optional"`
	Region          string `json:",default=cn-shanghai,optional"`
	Timeout         int    `json:",default=2000,optional"`
	Debug           bool   `json:",default=false,optional"`
}

func DefaultConfig(AccessKeyID, AccessKeySecret string) (config *Config) {
	config = &Config{
		AccessKeyID:     AccessKeyID,
		AccessKeySecret: AccessKeySecret,
		Endpoint:        defaultEndpoint,
		Protocol:        defaultProtocol,
		Region:          defaultRegion,
		Timeout:         defaultTimeout,
		Debug:           false,
	}
	return config
}

func NewConfig(c Config) *Config {
	return &c
}

func (c *Config) Check() error {
	if len(c.AccessKeyID) == 0 || len(c.AccessKeySecret) == 0 {
		logx.Errorf("AccessKeyId or AccessKeySecret is empty")
		return errors.New("AccessKeyId or AccessKeySecret is empty")
	}
	return nil
}
