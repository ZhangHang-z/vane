package commands

import (
	"github.com/ZhangHang-z/vane/src/dir"
	"github.com/ZhangHang-z/vane/src/down"
	"github.com/ZhangHang-z/vane/src/down/npm"
	fp "github.com/ZhangHang-z/vane/src/fileparser"
	"log"
)

var Install *tpInstall = newInstall()

// newInstall return a pointer receiver of type tpInstall.
func newInstall() *tpInstall {
	return &tpInstall{
		Name:      "install",
		JSONExist: true,
	}
}

// tpInstall used for store command name and usage infomation.
type tpInstall struct {
	Name      string
	VaneJSON  *fp.VaneJSON
	JSONExist bool
}

// Execute execute install command.
func (i *tpInstall) Execute(args ...string) error {
	if err := i.ReadJSONFile(); err != nil {
		if err == fp.ERR_JSON_FILE_NOT_FOUND {
			i.JSONExist = false
		} else {
			return err
		}
	}

	dir.MkSavedDirAndIn(i.VaneJSON.Directory)

	// install from json.
	if len(args) == 0 {
		if !i.JSONExist {
			return fp.ERR_JSON_FILE_NOT_FOUND
		}
		return i.InstallFromJSONFile()
	}

	// install by name.
	for _, name := range args {
		if err := InstallLatestVersion(name); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (i *tpInstall) ReadJSONFile() error {
	i.VaneJSON = new(fp.VaneJSON)

	// read the json file.
	err := i.VaneJSON.Read()
	if err != nil {
		return err
	}
	return nil
}

// InstallFromJSONFile just install packages from vane.json file.
func (i *tpInstall) InstallFromJSONFile() error {
	vj := i.VaneJSON
	if vj.Dependencies != nil {
		deps := vj.ReadPackages(vj.Dependencies)
		for _, dep := range deps {
			if err := InstallByVersion(dep.Name, dep.Version); err != nil {
				log.Println(err)
			}
		}
	}

	if vj.DevDependencies != nil {
		devDeps := vj.ReadPackages(vj.DevDependencies)
		for _, dev := range devDeps {
			if err := InstallByVersion(dev.Name, dev.Version); err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

func InstallByVersion(name, version string) error {
	defer func() {
		if info := recover(); info != nil {
			log.Println(info)
		}
	}() // registry this recover for ChooseDist

	url := npm.GetNPMRegistryURL(name)

	npmRepo, err := npm.NPMRegistryInit(url)
	if err != nil {
		return err
	}

	npmDist := npmRepo.ChooseDist(version)
	return StaringInstallAuto(name, npmDist.Tarball)
}

func InstallLatestVersion(name string) error {
	url := npm.GetNPMRegistryURL(name)

	npmRepo, err := npm.NPMRegistryInit(url)
	if err != nil {
		return err
	}

	npmDist := npmRepo.GetLatestDist()
	return StaringInstallAuto(name, npmDist.Tarball)
}

func StaringInstallAuto(name, tarBall string) error {
	contents, err := down.RveContentsFromLink(tarBall)
	if err != nil {
		return err
	}

	err = down.ExtractTarArchive(name, contents)
	if err != nil {
		log.Println(err)
	}
	log.Println("install %s successful", name)

	return nil
}
