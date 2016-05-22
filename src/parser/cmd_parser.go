/*
Package parser provied command line parser,
parse vane.json file and .vanerc file for user defined.
*/
package parser

import (
	"fmt"
	vcmd "github.com/ZhangHang-z/vane/src/commands"
	"github.com/ZhangHang-z/vane/src/down"
	"github.com/ZhangHang-z/vane/src/errors"
	fp "github.com/ZhangHang-z/vane/src/fileparser"
	"io"
	"os"
	"strings"
)

// CMDParser parse the command line arguments and execute vane's commands.
// this function must be executed before vane option parser.
// lefting os.Args slice (the flag starting with "-" or "--" prefix)
// will parsed by opt_parser.go file.
func CMDParser() error {
	if len(os.Args) <= 1 || os.Args[1] == "help" && len(os.Args) == 2 {
		CommandLine.PrintHelpInfo(helpInfoAll)
		return errors.ERR_PTR_HELP_STRING
	}
	cdName := strings.ToLower(os.Args[1])

	var (
		cdInfo string
		ok     bool
	)

	if cdInfo, ok = vcmd.IsValidCommand(cdName); !ok {
		CommandLine.PrintHelpInfo(helpInfoAll)
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
	os.Args = append(pwd, os.Args...) // recover command line path name self

	CommandLine.cmd = &Command{
		cmdFunc:  cdName,
		HelpInfo: cdInfo,
	}
	CommandLine.cmdArgs = cmdArgs
	CommandLine.option = os.Args
	CommandLine.cmdParsed = true
	return nil
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
		fp.MkSavedDirAndIn()
		for _, v := range args {
			if isDomain, _ := util.IsDomainName(v); isDomain {
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
		HelpInfo: vcmd.HelpString,
	}
	return cmd
}

type CommandSet struct {
	cmd            Command
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
	Name     string
	HelpInfo string
}
