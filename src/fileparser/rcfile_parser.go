package fileparser

import (
	"encoding/json"
	"github.com/ZhangHang-z/vane/src/errors"
	"io/ioutil"
	"path"
)

var (
	ERR_RC_FILE_NOT_FOUND = errors.New(".vanerc file not found.")
	ERR_RC_FILE_CONF      = errors.New(".vanerc file configuration error.")
)

const (
	RCFileName = ".vanerc"
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
	// if .vanerc file not exist. return error ERR_RC_FILE_NOT_FOUND,
	// return default &VaneRC{}, else resolve json data.
	if err != nil {
		return &VaneRC{}, ERR_RC_FILE_NOT_FOUND
	}
	return RsvJSONFromRCFile(contents)
}
