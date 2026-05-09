package cli

import "fmt"

func PrintRootHelp() {
	fmt.Println(`
rex-cli - Rex Platform CLI

Usage:
  rex-cli [command]

Available Commands:
  mas  Message Management services
  kms  Key Management services

Flags:
  -f, --file string   specify an alternate config file (default: ~/.rex/config.json)
  -h, --help          help for rex-cli

Use "rex-cli [command] --help" for more information about a command.
`)
}
