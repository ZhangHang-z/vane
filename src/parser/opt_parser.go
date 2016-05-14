package parser

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ZhangHang-z/vane/src/vflag"
	"os"
	"strings"
)

const (
	defaultInstall = "all"
	defaultSave    = false
)

type Pkgs []string

func (ps *Pkgs) String() string {
	return fmt.Sprint(*ps)
}

func (ps *Pkgs) Set(value string) error {
	fmt.Println(*ps)
	if len(*ps) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		*ps = append(*ps, dt)
	}
	return nil
}

var ArgIsOpt bool
var ArgPkgs Pkgs

var Save bool
var SaveDev bool

func printHelpString() {
	fmt.Println(HelpString)
}

func OPTParser() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", "vane")
		printHelpString()
	}

	flag.BoolVar(&ArgIsOpt, "save-dev", defaultSave, "--save....")
	flag.Var(&ArgPkgs, "p", "print....")

	vflag.VaneBoolValue(&Save, "save", false, "--save..")
	flag.Parse()
}
