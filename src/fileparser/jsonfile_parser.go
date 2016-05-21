package fileparser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var (
	ERR_JSON_FILE_NOT_FOUND = errors.New("vane.json file not found.")
	ERR_JSON_FILE_CONF      = errors.New("vane.json file configuration error.")
)

const (
	JSON_FILE_NAME = "vane.json" // default saved directory named "vane.json"
	DFT_TIME_OUT   = 36000       // default timeour set to 10 hours.
)

type VaneJSON struct {
	Directory     string            `json:"directory,omitempty"`
	Timeout       int               `json:"timeout,omitempty"`
	Dependencies  map[string]string `json:"dependencies,omitempty"`
	DevDependcies map[string]string `json:"devDependencies,omitempty"`
}

func (v *VaneJSON) Read() error {
	stream, err := ioutil.ReadFile(JSON_FILE_NAME)
	if err != nil {
		return err
	}
	err = json.Unmarshal(stream, v)
	if err != nil {
		return err
	}
	return nil
}

func (v *VaneJSON) ReadPackages(pkgs map[string]string) []PackageInfo {
	var deps []PackageInfo
	for k, v := range pkgs {
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
