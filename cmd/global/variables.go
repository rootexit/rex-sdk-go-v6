package global

import (
	"github.com/rootexit/rex-sdk-go-v6/rex"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
)

var (
	GlobalConfigPath string
	GlobalConfig     rexConfig.Config
	GlobalRexSdk     *rex.Sdk
)
