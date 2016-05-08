package vane

import (
	"fmt"
	"os"
	"testing"
)

func TestMkdftDir(t *testing.T) {
	e := MkDftDir()
	if e != nil {
		fmt.Println(e)
	}
}

func TestDftDirIsExist(t *testing.T) {
	d, e := os.Getwd()
	fmt.Println(d)
	if e != nil {
		t.Error(e)
	}

}
