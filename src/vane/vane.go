package vane

import (
	"log"
	"os"
	"path/filepath"
)

// DftDirIsExist inspect the installation directory exist or not exist.
// The default name is "vane_components".
func DftDirIsExist() bool {
	_, err := os.Stat(DefaultDirName)
	return err == nil || os.IsExist(err)
}

// MkDftDir create the default direatory which name is "vane_components".
func MkDftDir() bool {
	err := os.Mkdir(DefaultDirName, ModeCommonDir)
	if err != nil {
		return false
	}
	return true
}

func GotoComponentsDir(dirComponents string) {
	absDir, err := filepath.Abs(dirComponents)
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(absDir)
}
