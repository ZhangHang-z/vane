package commands

import (
	"archive/tar"
	"github.com/ZhangHang-z/vane/src/down"
	fp "github.com/ZhangHang-z/vane/src/fileparser"
)

var Install *tInstall = newInstall()

func newInstall() *tInstall {
	return &tInstall{
		Name:  "install",
		Usage: "Install a pakcage into pkg-directory by given name.",
	}
}

type tInstall struct {
	Name  string
	Usage string
}

func (i *tInstall) Execute(args ...string) error {
	if len(args) == 0 {
		InstallFromJsonFile()
	}
	return nil
}

func (i *tInstall) RollBack() error {
	return nil
}

func (i *tInstall) InstallFromJsonFile() {
	fj := fp.ParseJSONFile()
}
