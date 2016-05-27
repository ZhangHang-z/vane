package commands

import (
	"fmt"
	"github.com/ZhangHang-z/vane/src/down"
	"github.com/ZhangHang-z/vane/src/down/npm"
	fp "github.com/ZhangHang-z/vane/src/fileparser"
	"log"
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
		return i.InstallFromJSONFile()
	}
	return nil
}

// InstallFromJSONFile just install packages from vane.json file.
func (i *tpInstall) InstallFromJSONFile() error {
	vj := new(fp.VaneJSON)

	// read the json file.
	err := vj.Read()
	if err != nil {
		return err
	}

	if err = fp.MkSavedDirAndIn(); err != nil {
		return err
	}

	if vj.Dependencies != nil {
		deps := vj.ReadPackages(vj.Dependencies)
		for _, dep := range deps {
			InstallByVersion(dep.Name, dep.Version)
		}
	}

	if vj.DevDependencies != nil {
		devDeps := vj.ReadPackages(vj.DevDependencies)
		for _, dev := range devDeps {
			InstallByVersion(dev.Name, dev.Version)
		}
	}

	return nil
}

func InstallByVersion(name, version string) {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println(info)
		}
	}()
	url := npm.GetNPMRegistryURL(name)
	npmRepo := npm.NPMRegistryInit(url)
	npmDist := npmRepo.ChooseDist(version)
	StaringInstallAuto(npmDist.Tarball)
}

func InstallLatestVersion(name string) {
	url := npm.GetNPMRegistryURL(name)
	npmRepo := npm.NPMRegistryInit(url)
	npmDist := npmRepo.GetLatestVersion()
	StaringInstallAuto(npmDist.Tarball)
}

func StaringInstallAuto(tarBall string) error {
	contents, err := down.RveContentsFromLink(tarBall)
	if err != nil {
		return err
	}

	err = down.ExtractTarArchive(name, contents)
	if err != nil {
		log.Println(err)
	}

	return nil
}
