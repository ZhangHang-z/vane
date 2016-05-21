package dir

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

const (
	ModeCommonDir  os.FileMode = 0775 // unix common user's directory default permission.
	ModeCommonFile os.FileMode = 0664 // unix common user's file default permission.
)

const (
	DefaultDirName string = "vane_components" // default hold directory of installed package.
)

var (
	ERR_MK_SAVE_DIR = errors.New("make package saved directory failed.")
)

// DirIsExist inspect the installation directory exist or not exist.
func DirIsExist(dirName string) bool {
	_, err := os.Stat(dirName)
	return err == nil || os.IsExist(err)
}

// MkSavedDir create the default direatory which name is "vane_components".
func MkSavedDir(dirName string) error {
	err := os.MkdirAll(dirName, ModeCommonDir)
	if err != nil {
		return ERR_MK_SAVE_DIR
	}
	return nil
}

// GotoComponentsDir go to the package saved directory.
func GotoComponentsDir(dirComponents string) {
	absDir, err := filepath.Abs(dirComponents)
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(absDir)
	return
}
