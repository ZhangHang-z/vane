package main

import (
	"github.com/ZhangHang-z/vane/src/parser"
)

func Run() error {
	err := parser.CMDParser()
	if err != nil {
		return err
	}
	parser.OPTParser()
	return nil
}

func main() {
	Run() // omit error.
}
