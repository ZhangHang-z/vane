package commands

import (
	"fmt"
	"github.com/ZhangHang-z/vane/src/down/npm"
	//"github.com/ZhangHang-z/vane/src/errors"
	fp "github.com/ZhangHang-z/vane/src/fileparser"
)

var Install *tpInstall = newInstall()

// newInstall return a pointer receiver of type tpInstall.
func newInstall() *tpInstall {
	return &tpInstall{
		Name:  "install",
		Usage: "Install a pakcage into pkg-directory by given name.",
	}
}

// tpInstall used for store command name and usage infomation.
type tpInstall struct {
	Name  string
	Usage string
}

// Execute execute install command.
func (i *tpInstall) Execute(args ...string) error {
	if len(args) == 0 {
		i.InstallFromJSONFile()
	}
	return nil
}

func (i *tpInstall) RollBack() error {
	return nil
}

// InstallFromJSONFile just install packages from vane.json file.
func (i *tpInstall) InstallFromJSONFile() error {
	vj := new(fp.VaneJSON)

	err := vj.Read()
	if err != nil {
		return err
	}

	if vj.Dependencies != nil {
		deps := vj.ReadPackages(vj.Dependencies)
		for _, dep := range deps {
			StaringInstall(dep.Name, dev.Version)
		}
	}

	if vj.DevDependencies != nil {
		devDeps := vj.ReadPackages(vj.DevDependencies)
		for _, dev := range devDeps {
			StaringInstall(dev.Name, dev.Version)
		}
	}

	return nil
}

func StaringInstall(name, version string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	url := npm.GetNPMRegistryURL(name)
	npmRepo := npm.NPMRegistryInit(url)
	npmDist := npmRepo.ChooseOneDist(version)

	fmt.Println(npmDist.Tarball, npmDist.Shasum)
	return nil
}
