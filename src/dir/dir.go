package dir

import (
	"errors"
	"github.com/ZhangHang-z/vane/src/vane"
	"log"
	"os"
	"path/filepath"
)

const (
	DefaultDirName string = "vane_components" // default hold directory of installed package.
)

var (
	ERR_MK_SAVE_DIR = errors.New("make package saved directory failed.")
)

// DftDirIsExist inspect the installation directory exist or not exist.
// The default name is "vane_components".
func DirIsExist(dirName string) bool {
	_, err := os.Stat(dirName)
	return err == nil || os.IsExist(err)
}

// MkDftDir create the default direatory which name is "vane_components".
func MkSavedDir(dirName string) error {
	err := os.Mkdir(dirName, vane.ModeCommonDir)
	if err != nil {
		return ERR_MK_SAVE_DIR
	}
	return nil
}

func GotoComponentsDir(dirComponents string) {
	absDir, err := filepath.Abs(dirComponents)
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(absDir)
	return
}
