package parser

import (
	"fmt"
	"testing"
)

func TestOpenRCFile(t *testing.T) {
	fmt.Println("-----------------")
	vanerc, err := RsvRCFile("../../")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vanerc)
}

func TestMkSavedDirAndIn(t *testing.T) {
	err := MkSavedDirAndIn()
	if err != nil {
		t.Error(err)
	}
}
