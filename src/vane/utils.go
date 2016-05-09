package vane

import (
	"os"
)

const (
	ModeCommonDir  os.FileMode = 0775              // unix common user's directory default permission.
	ModeCommonFile os.FileMode = 0664              // unix common user's file default permission.
	DefaultDirName string      = "vane_components" // default hold directory of installed package.
)
