package fileparser

import (
	"fmt"
	"os"
	"testing"
)

func TestOpenRCFile(t *testing.T) {
	vanerc, err := RsvRCFile("../../")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(".vanerc", vanerc)
}

func TestMkSavedDirAndIn(t *testing.T) {
	err := MkSavedDirAndIn()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(os.Getwd())
}
