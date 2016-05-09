package parser

import (
	"fmt"
	"os"
)

func CMDParser() []string {
	command := os.Args[1]
	has, cmd := IsValidCommand(command)
	if has {
		fmt.Println(cmd)
	}
	return os.Args
}

func IsValidCommand(command string) (bool, interface{}) {
	cmd, hasAttr := Commands[command]
	if hasAttr {
		return true, cmd
	}
	return false, nil
}
