package vane

import (
	"os"
)

const (
	ModeCommonDir  os.FileMode = 0775              // unix common user's directory default permission.
	ModeCommonFile os.FileMode = 0664              // unix common user's file default permission.
	DefaultDirName string      = "vane_components" // default hold directory of installed package.
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
