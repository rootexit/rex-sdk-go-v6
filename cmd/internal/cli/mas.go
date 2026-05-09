package cli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rootexit/rex-sdk-go-v6/cmd/global"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexTypes"
)

func HandleMas(args []string) {
	if len(args) == 0 {
		printMasHelp()
		return
	}
	switch args[0] {
	case "--help", "-h":
		printMasHelp()
	case "captcha_config":
		handleMasCaptchaConfig(args[1:])
	default:
		fmt.Printf("unknown mas module: %s\n", args[0])
	}
}

func handleMasCaptchaConfig(args []string) {
	switch args[0] {
	case "test":
		test()
	}
}

func test() error {
	code, result, err := global.GlobalRexSdk.MasService.BehavioralVerificationInit(context.Background(), &rexTypes.BehavioralVerificationInitReq{
		Key:     "default",
		Service: "aa.lilsite.com",
		Type:    "login",
	})
	if err != nil {
		fmt.Printf("rex-cli error: [code: %d] [error: %s]\n", code, err)
		return err
	}
	tempJson, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("rex-cli result: %s \n", tempJson)
	return nil
}

func printMasHelp() {

	fmt.Println(`

rexCli mas

Usage:

  rexCli mas <service> <action>

Resources:

  captcha_config

Examples:

  rexCli mas captcha_config queryList

`)

}
