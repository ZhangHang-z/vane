package parser

import (
	"flag"
)

const (
	defaultInstall = "all"
)

var ArgI string
var ArgInstall string

func init() {
	flag.StringVar(&ArgI, "i", defaultInstall, "vane -i pkg")
	flag.StringVar(&ArgInstall, "install", defaultInstall, "vane -install pkg")
	flag.Parse()
}

func OPTParser() {
}
