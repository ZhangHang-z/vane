package parser

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func VTestCMDParser(t *testing.T) {
	CMDParser()
	fmt.Println(os.Args)
}

func TestCmdline(t *testing.T) {
	f, err := os.Open("./errors.go")
	if err != nil {
		fmt.Println(err)
	}
	CommandLine.SetErrHandling(f)
}
