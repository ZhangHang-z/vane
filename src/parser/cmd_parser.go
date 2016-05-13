/*
	Package parser provied command line parser,
	parse vane.json file and .vanerc file for user defined.
*/
package parser

import (
	"fmt"
	"os"
)

// CMDParser parse the command line arguments and execute vane's commands.
// this function must be executed before vane option parser.
// lefting os.Args slice (the flag starting with "-" or "--" prefix)
// will parsed by opt_parser.go file.
func CMDParser() {
	if len(os.Args) <= 1 || os.Args[1] == "help" {
		fmt.Println(HelpString)
		return
	}
	command := os.Args[1]

	if !IsValidCommand(command) {
		fmt.Println(HelpString)
	}
	os.Args = os.Args[2 : len(os.Args)-1]

	var cmdArgs []string
	for i, v := range os.Args {
		if []byte(v)[0] != '-' {
			cmdArgs = append(cmdArgs, v)
			continue
		}
		os.Args = os.Args[i:] // lefting os.Args
		break
	}

	ExeCommand(command, cmdArgs)
}

func ExeCommand(cmd string, args []string) {
	if len(args) == 0 {
		ExeNoArgsCMD(cmd)
		return
	}
	ExeMultiArgsCMD(cmd, args)
}

func ExeNoArgsCMD(cmd string) {
	return
}

func ExeMultiArgsCMD(cmd string, args []string) {

}

func IsValidCommand(command string) bool {
	_, hasAttr := AllCommands[command]
	if hasAttr {
		return true
	}
	return false
}
