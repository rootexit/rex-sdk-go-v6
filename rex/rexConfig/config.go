package rexConfig

import (
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	defaultTimeout  = 2000
	defaultProtocol = "https"
	defaultRegion   = "cn-shanghai"
	defaultEndpoint = "api.rootexit.com"
)

type Config struct {
	// Access key ID
	AccessKeyID string `json:"AccessKeyId,optional,inherit" mapstructure:"AccessKeyId" yaml:"AccessKeyId"`
	// Secret Access Key
	AccessKeySecret string `json:"AccessKeySecret,optional,inherit" mapstructure:"AccessKeySecret" yaml:"AccessKeySecret"`
	Endpoint        string `json:"Endpoint,default=api.rootexit.com,optional" mapstructure:"Endpoint" yaml:"Endpoint"`
	Protocol        string `json:"Protocol,default=https,options=http|https,optional" mapstructure:"Protocol" yaml:"Protocol"`
	Region          string `json:"Region,default=cn-shanghai,optional" mapstructure:"Region" yaml:"Region"`
	Timeout         int    `json:"Timeout,default=2000,optional" mapstructure:"Timeout" yaml:"Timeout"`
	Debug           bool   `json:"Debug,default=false,optional" mapstructure:"Debug" yaml:"Debug"`
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
