package parser

import (
	"encoding/json"
	"errors"
	"github.com/ZhangHang-z/vane/src/dir"
	"io/ioutil"
	"os"
	"path"
)

const (
	RCFileName   = ".vanerc"
	JSONFileName = "vane.json"
)

var (
	ERR_RC_FILE_NOT_FOUND   = errors.New(".vanerc file not found.")
	ERR_JSON_FILE_NOT_FOUND = errors.New("vane.json file not found.")
	ERR_RC_FILE_CONF        = errors.New(".vanerc file configuration error.")
	ERR_JSON_FILE_CONF      = errors.New("vane.json file configuration error.")
)

type VaneRC struct {
	Directory string
	TimeOut   int32
	Scripts   interface{}
}

func OpenRCFile(name string) ([]byte, error) {
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func RsvJSONFromRCFile(rawJSONs []byte) (*VaneRC, error) {
	var vanerc VaneRC
	err := json.Unmarshal(rawJSONs, &vanerc)
	if err != nil {
		return nil, ERR_RC_FILE_CONF
	}
	return &vanerc, nil
}

func RsvRCFile(rcfpath string) (*VaneRC, error) {
	rcfpath = path.Join(rcfpath, RCFileName)
	contents, err := OpenRCFile(rcfpath)
	if err != nil {
		return &VaneRC{}, ERR_RC_FILE_NOT_FOUND
	}
	return RsvJSONFromRCFile(contents)
}

func MkSavedDirAndIn() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	vanerc, err := RsvRCFile(cwd)
	if err != nil {
		if err == ERR_RC_FILE_NOT_FOUND {
			vanerc.Directory = dir.DefaultDirName
		}
	}

	// default package saved directory.
	if vanerc.Directory == dir.DefaultDirName {
		if !dir.DirIsExist(dir.DefaultDirName) {
			err := dir.MkSavedDir(dir.DefaultDirName)
			if err != nil {
				return errors.New("make default package saved directory <vane_components> failed.")
			}
		}
	}

	// user defined package saved directory from .vanerc file.
	if !dir.DirIsExist(vanerc.Directory) {
		err := dir.MkSavedDir(vanerc.Directory)
		if err != nil {
			return err
		}
	}

	dir.GotoComponentsDir(dir.DefaultDirName)
	return nil
}
