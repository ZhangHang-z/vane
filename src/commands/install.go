package commands

import (
	//"archive/tar"
	"fmt"
	fp "github.com/ZhangHang-z/vane/src/fileparser"
	"github.com/ZhangHang-z/vane/src/verrors"
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
		i.InstallFromJsonFile()
	}
	return nil
}

func (i *tInstall) RollBack() error {
	return nil
}

func (i *tInstall) InstallFromJsonFile() error {
	vj := new(fp.VaneJSON)

	err := vj.Read()
	if err != nil {
		return err
	}

	if vj.Dependencies != nil {
		deps := vj.ReadPackages(vj.Dependencies)
		for i, v := range deps {
			fmt.Println(i, v.Name)
		}
	}

	if vj.Dependencies != nil {
		devDeps := vj.ReadPackages(vj.DevDependcies)
		for i, v := range devDeps {
			fmt.Println(i, v.Name)
		}
	}

	return nil
}
