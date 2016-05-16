package main

import (
	vErr "github.com/ZhangHang-z/vane/src/errors"
	"github.com/ZhangHang-z/vane/src/parser"
)

func Run() error {
	err := parser.CMDParser()
	switch e := err.(type) {
	case vErr.ERR_PTR_HELP_STRING:
		return e
	case nil:
		break
	}
	parser.OPTParser()
	return nil
}

func main() {
	Run() // omit error.
}
