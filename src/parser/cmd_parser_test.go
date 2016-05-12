package parser

import (
	"fmt"
	"os"
	"testing"
)

func TestCMDParser(t *testing.T) {
	CMDParser()
	fmt.Println(os.Args)
}
