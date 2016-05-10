package dir

import (
	"fmt"
	"os"
	"testing"
)

func TestMkSavedDir(t *testing.T) {
	e := MkSavedDir(DefaultDirName)
	if e != nil {
		fmt.Println(e)
	}
}

func TestDirIsExist(t *testing.T) {
	d, e := os.Getwd()
	fmt.Println(d)
	if e != nil {
		t.Error(e)
	}

}
