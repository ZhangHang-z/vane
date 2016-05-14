/*
	Package parser provied command line parser,
	parse vane.json file and .vanerc file for user defined.
*/
package parser

import (
	"fmt"
	"github.com/ZhangHang-z/vane/src/down"
	"os"
)

// CMDParser parse the command line arguments and execute vane's commands.
// this function must be executed before vane option parser.
// lefting os.Args slice (the flag starting with "-" or "--" prefix)
// will parsed by opt_parser.go file.
func CMDParser() {
	var pwd []string = make([]string, 1, 10)
	pwd[0] = os.Args[0]

	if len(os.Args) <= 1 || os.Args[1] == "help" && len(os.Args) == 2 {
		fmt.Println(HelpString)
		return
	}
	command := os.Args[1]

	if !IsValidCommand(command) {
		fmt.Println(HelpString)
	}
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
	ExeCommand(command, cmdArgs)
	os.Args = append(pwd, os.Args...)
}

func ExeCommand(cmd string, args []string) {
	if len(args) != 0 {
		ExeMultiArgsCMD(cmd, args)
	} else {
		ExeNoArgsCMD(cmd)
	}
	return
}

func ExeNoArgsCMD(cmd string) {
	switch cmd {
	case "install":
		fmt.Println("One argument install...")
	case "info":
		fmt.Println(HelpString)
	default:
		fmt.Println(HelpString)
	}
	return
}

func ExeMultiArgsCMD(cmd string, args []string) {
	switch cmd {
	case "install":
		MkSavedDirAndIn()
		for _, v := range args {
			if isDomain, _ := down.IsDomainName(v); isDomain {
				url := down.RsvGitHubAddr(v)
				err := down.GitHubDownloader(url)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				continue
			}
		}
	case "info":
		fmt.Println("info command.....")
	default:
		fmt.Println("default multiple arguments....")
	}
	return
}

// ------
type DomainPackage struct {
	name    string
	version string
	source  string
}
