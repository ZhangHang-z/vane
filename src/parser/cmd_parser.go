/*
Package parser provied command line parser,
parse vane.json file and .vanerc file for user defined.
*/
package parser

import (
	vcmd "github.com/ZhangHang-z/vane/src/commands"
	"github.com/ZhangHang-z/vane/src/errors"
	"log"
	"os"
	"strings"
)

// CMDParser parse the command line arguments and execute vane's commands.
// this function must be executed before vane option parser.
// lefting os.Args slice (the flag starting with "-" or "--" prefix)
// will parsed by opt_parser.go file.
func CMDParser() error {
	if len(os.Args) <= 1 || os.Args[1] == "help" && len(os.Args) == 2 {
		log.Println(GenHelpInfo())
		return errors.ERR_PTR_HELP_STRING
	}
	cdName := strings.ToLower(os.Args[1])

	if _, ok := Usage[cdName]; !ok {
		log.Println(GenHelpInfo())
		return errors.ERR_PTR_HELP_STRING
	}

	var pwd []string = make([]string, 1, 10)
	pwd[0] = os.Args[0]

	os.Args = os.Args[2:len(os.Args)]
	var cmdArgs []string
	for i, v := range os.Args {
		if []byte(v)[0] != '-' {
			cmdArgs = append(cmdArgs, v)
			continue
		}
		os.Args = os.Args[i:] // lefting os.Args
		break
	}

	// recover command line path name self
	os.Args = append(pwd, os.Args...)
	return executeCommand(cdName, cmdArgs)
}

func executeCommand(cmd string, args []string) error {
	switch cmd {
	case "install":
		return vcmd.Install.Execute(args...)
	case "info":
		log.Println("commmand info missed.")
	}
	return nil
}
