package fileparser

import (
	"encoding/json"
	"github.com/ZhangHang-z/vane/src/dir"
	"github.com/ZhangHang-z/vane/src/errors"
	"io/ioutil"
	"regexp"
)

var (
	ERR_JSON_FILE_NOT_FOUND = errors.New("vane.json file not found")
	ERR_JSON_FILE_CONF      = errors.New("bad vane.json configuration")
)

const (
	JSON_FILE_NAME = "vane.json" // default saved directory named "vane.json"
	DFT_TIME_OUT   = 36000       // default timeour set to 10 hours.
)

type VaneJSON struct {
	Directory       string            `json:"directory,-"`
	Timeout         int               `json:"timeout,-"`
	Dependencies    map[string]string `json:"dependencies,omitempty"`
	DevDependencies map[string]string `json:"devDependencies,omitempty"`
}

func (v *VaneJSON) Read() error {
	// make sure it has a name, it not in json file, then using default name.
	v.Directory = dir.DefaultDirName

	stream, err := ioutil.ReadFile(JSON_FILE_NAME)
	if err != nil {
		return ERR_JSON_FILE_NOT_FOUND
	}

	err = json.Unmarshal(stream, v)
	if err != nil {
		return ERR_JSON_FILE_CONF
	}
	return nil
}

// ReadPackages read the given map that might be VaneJSON.Dependencies or VaneJSON.DebDependencies.
func (v *VaneJSON) ReadPackages(pkgs map[string]string) []PackageInfo {
	var deps []PackageInfo
	for k, v := range pkgs {
		v = GetVersionInfo(v)
		deps = append(deps, PackageInfo{k, v})
	}
	return deps
}

type PackageInfo struct {
	Name    string
	Version string
}

type writeWhat int

const (
	depend = iota
	devDepend
)

func (v *VaneJSON) Write(code writeWhat) error {
	switch c := code; c {
	case 0:
		return v.writeToDepend()
	case 1:
		return v.writeToDevDepend()
	}
	return nil
}

func (v *VaneJSON) writeToDepend() error {
	return nil
}

func (v *VaneJSON) writeToDevDepend() error {
	return nil
}

const (
	RegPackageVersion = `\d+\.\d+(\.\d+.*)?` // will match, e.g. 1.1.1, 1.11.1, 1.2, 111.0.1, 1.0.0-alpha.7
)

func GetVersionInfo(raw string) string {
	re := regexp.MustCompile(RegPackageVersion)
	return re.FindString(raw)
}
