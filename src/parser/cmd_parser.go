/*
	Package parser provied command line parser,
	parse vane.json file and .vanerc file for user defined.
*/
package parser

import (
	"fmt"
	vcmd "github.com/ZhangHang-z/vane/src/commands"
	"github.com/ZhangHang-z/vane/src/down"
	"io"
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
	cd := os.Args[1]

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
	os.Args = append(pwd, os.Args...)

	cdFunc, cdInfo := GetCommand(cd)

	CommandLine.cmd = &Command{
		cmdFunc:  cdFunc,
		HelpInfo: cdInfo,
	}
	CommandLine.cmdArgs = cmdArgs
	CommandLine.option = os.Args
	CommandLine.cmdParsed = true
	//ExeCommand(cd, cmdArgs)
}

func GetCommand(name string) (interface{}, string) {

}

func ExeCommand(cmd string, args []string) {
	if len(args) != 0 {
		ExeMultiArgsCMD(cmd, args)
	} else {
		ExeNoArgsCMD(cmd)
	}
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

type infoCode int

const (
	helpInfoAll infoCode = iota
	helpInfoPart
)

var CommandLine = NewCommandSet(os.Stdout)

func NewCommandSet(out io.Writer) *CommandSet {
	cmd := &CommandSet{
		output:   out,
		HelpInfo: HelpString,
	}
	return cmd
}

type CommandSet struct {
	cmd            *Command
	cmdArgs        []string
	cmdParsed      bool
	option         string
	HelpInfo       string
	output         io.Writer
	errHandling    io.Writer
	errHandlingSet bool
}

func (cs *CommandSet) PrintHelpInfo(info infoCode) error {
	switch info {
	case 0:
		fmt.Fprintln(cs.output, cs.HelpInfo)
	case 1:
		fmt.Fprintln(cs.output, cs.cmd.HelpInfo)
	default:
		break
	}
	return nil
}

func (cs *CommandSet) SetErrHandling(w io.Writer) {
	if !cs.errHandlingSet {
		cs.errHandling = w
		cs.errHandlingSet = true
	}
}

type Command struct {
	cmdFunc  func(args ...string) error
	HelpInfo string
}
