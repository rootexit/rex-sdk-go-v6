package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rootexit/rex-sdk-go-v6/cmd/global"
	"github.com/rootexit/rex-sdk-go-v6/cmd/internal/cli"
	"github.com/rootexit/rex-sdk-go-v6/rex"
	"github.com/zeromicro/go-zero/core/conf"
)

func main() {
	flag.StringVar(
		&global.GlobalConfigPath,
		"f",
		"",
		"config file path",
	)

	flag.StringVar(
		&global.GlobalConfigPath,
		"file",
		"",
		"config file path",
	)
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		cli.PrintRootHelp()
		return
	}
	log.Println(args)
	var err error

	if global.GlobalConfigPath == "" {
		global.GlobalConfigPath, err = global.GetConfigPath()
		if err != nil {
			fmt.Printf("rex-cli error: %s\n", err)
		}
		//fmt.Printf("rex-cli config path: %s\n", global.GlobalConfigPath)
	}
	if args[0] != "config" {
		if err = conf.Load(global.GlobalConfigPath, &global.GlobalConfig); err != nil {
			if args[0] != "config" {
				// note: 如果报错，则输出，cli尚未初始化配置，请调用rex-cli config
				fmt.Println("rex-cli error: The CLI has not been initialized with configuration. Please run [ rex-cli config ].")
				return
			}
		}
		// note: init sdk
		global.GlobalRexSdk, err = rex.NewSdk(&global.GlobalConfig)
		if err != nil {
			fmt.Println("rex-cli error: Failed to initialize rex sdk.")
			return
		}
	}
	switch args[0] {
	case "mas":
		cli.HandleMas(args[1:])
	case "config":
		if err := cli.HandleConfig(args[1:]); err != nil {
			fmt.Printf("rex-cli error: %s\n", err)
			return
		}
	case "-f", "--file":
		// note: init sdk
		if err = conf.Load(global.GlobalConfigPath, &global.GlobalConfig); err != nil {
			if args[0] != "config" {
				// note: 如果报错，则输出，cli尚未初始化配置，请调用rex-cli config
				fmt.Println("rex-cli error: The CLI has not been initialized with configuration. Please run [ rex-cli config ].")
				return
			}
		}
		global.GlobalRexSdk, err = rex.NewSdk(&global.GlobalConfig)
		if err != nil {
			fmt.Println("rex-cli error: Failed to initialize rex sdk.")
			return
		}
	case "--help", "-h":
		cli.PrintRootHelp()
	default:
		fmt.Printf("rex-cli error: unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
