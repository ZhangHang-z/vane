package parser

import (
	"fmt"
	"os"
)

func CMDParser() {
	if len(os.Args) <= 1 || os.Args[1] == "help" {
		fmt.Println(HelpString)
		return
	}
	command := os.Args[1]
	has, cmd := IsValidCommand(command)
	if !has {
		fmt.Println(HelpString)
		return
	}
	fmt.Println(cmd)
}

func IsValidCommand(command string) (bool, interface{}) {
	cmd, hasAttr := AllCommands[command]
	if hasAttr {
		return true, cmd
	}
	return false, nil
}
