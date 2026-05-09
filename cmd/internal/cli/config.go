package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rootexit/rex-sdk-go-v6/cmd/global"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
	"github.com/zeromicro/go-zero/core/conf"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

func HandleConfig(args []string) error {
	if len(args) == 0 {
		return configWizard()
	}
	switch args[0] {
	case "-g", "--get", "get":
		return configGet()
	case "-p", "--path", "path":
		return pathGet()
	case "-h", "--help", "help":
		printConfigHelp()
		return nil
	default:
		return fmt.Errorf("unknown config action: %s", args[0])
	}

}

func configSetting() error {
	log.Println("现在开始配置初始化流程")
	return nil
}

func configGet() error {
	defaultPath, _ := global.GetConfigPath()
	if global.GlobalConfigPath != defaultPath {
		fmt.Println("Using config file:", global.GlobalConfigPath)
	}
	if err := conf.Load(global.GlobalConfigPath, &global.GlobalConfig); err != nil {
		// note: 如果报错，则输出，cli尚未初始化配置，请调用rex-cli config
		fmt.Println("rex-cli error: The CLI has not been initialized with configuration. Please run [ rex-cli config ].")
		return err
	}
	newModel := global.GlobalConfig
	newModel.AccessKeySecret = "******"
	bt, err := yaml.Marshal(newModel)
	if err != nil {
		return err
	}
	fmt.Println(string(bt))
	return nil
}

func pathGet() error {
	fmt.Println("Using config file:", global.GlobalConfigPath)
	return nil
}

func printConfigHelp() {

	fmt.Println(`
rex-cli config - rex-cli configuration

Usage:
  rex-cli config
  rex-cli config [command]

Available Commands:
  config  	   交互式输入配置
  config get   获取当前的配置文件
  config path  获取当前的配置文件的路径

Flags:
  -g, --get, get    获取当前的配置文件
  -p, --path, path   获取当前的配置文件的路径

`)

}

func configWizard() error {
	oldCfg := global.GlobalConfig
	reader := bufio.NewReader(os.Stdin)
	accessKey, err := promptInput(reader, "Access Key", oldCfg.AccessKeyID)
	if err != nil {
		return err
	}
	accessSecret, err := promptPassword("Access Secret", oldCfg.AccessKeySecret != "")
	if err != nil {
		return err
	}
	if accessSecret == "" {
		accessSecret = oldCfg.AccessKeySecret
	}
	endpoint, err := promptInput(reader, "Endpoint", oldCfg.Endpoint)
	if err != nil {
		return err
	}
	//region, err := promptInput(reader, "Region", defaultString(oldCfg.Region, "cn-shanghai"))
	//if err != nil {
	//	return err
	//}
	//protocol, err := promptInput(reader, "Protocol", defaultString(oldCfg.Protocol, "https"))
	//if err != nil {
	//	return err
	//}
	//timeout, err := promptInt(reader, "Timeout", defaultInt(oldCfg.Timeout, 2000))
	//if err != nil {
	//	return err
	//}
	//debug, err := promptBool(reader, "Debug", oldCfg.Debug)
	//if err != nil {
	//	return err
	//}
	cfg := rexConfig.Config{
		AccessKeyID:     accessKey,
		AccessKeySecret: accessSecret,
		Endpoint:        endpoint,
		Protocol:        rexConfig.DefaultProtocol,
		Region:          rexConfig.DefaultRegion,
		Timeout:         rexConfig.DefaultTimeout,
		Debug:           false,
	}
	if err = saveConfig(cfg); err != nil {
		return err
	}
	path, _ := global.GetConfigPath()
	fmt.Println("config saved:", path)
	return nil

}

func promptInput(reader *bufio.Reader, label string, defaultValue string) (string, error) {
	if defaultValue != "" {
		fmt.Printf("%s [%s]: ", label, defaultValue)
	} else {
		fmt.Printf("%s: ", label)
	}
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSpace(text)
	if text == "" {
		return defaultValue, nil
	}
	return text, nil

}

func promptPassword(label string, hasOld bool) (string, error) {
	if hasOld {
		fmt.Printf("%s [******]: ", label)
	} else {
		fmt.Printf("%s: ", label)
	}
	bt, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println("[******]")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bt)), nil

}

func promptInt(reader *bufio.Reader, label string, defaultValue int) (int, error) {

	text, err := promptInput(reader, label, fmt.Sprintf("%d", defaultValue))
	if err != nil {
		return 0, err
	}
	var v int
	_, err = fmt.Sscanf(text, "%d", &v)
	if err != nil {
		return 0, fmt.Errorf("%s must be int", label)
	}
	return v, nil

}

func promptBool(reader *bufio.Reader, label string, defaultValue bool) (bool, error) {

	def := "false"
	if defaultValue {
		def = "true"
	}
	text, err := promptInput(reader, label, def)
	if err != nil {
		return false, err
	}
	switch strings.ToLower(text) {
	case "true", "t", "yes", "y", "1":
		return true, nil
	case "false", "f", "no", "n", "0":
		return false, nil
	default:
		return false, fmt.Errorf("%s must be bool", label)
	}

}

func saveConfig(cfg rexConfig.Config) error {
	path, err := global.GetConfigPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	bt, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, bt, 0600)
}

func defaultString(v string, def string) string {

	if v == "" {
		return def
	}
	return v

}

func defaultInt(v int, def int) int {

	if v == 0 {
		return def
	}
	return v

}
