package vane

import (
	"os"
	"syscall"
)

const (
	ModeCommonDir  os.FileMode = 0775              // unix common user's directory default permission.
	ModeCommonFile os.FileMode = 0664              // unix common user's file default permission.
	DefaultDirName string      = "vane_components" // default hold directory of installed package.
)

// DefaultDirIsExist inspect the installation directory exist or not exist.
// The default name is "vane_components".
func DefaultDirIsExist() bool {
	_, err := os.Stat(DefaultDirName)
	return err == nil || os.IsExist(err)
}

// MkDefDir create the default direatory which name is "vane_components".
func MkDefDir() error {
	return os.Mkdir(DefaultDirName, ModeCommonDir)
}
