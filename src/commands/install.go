package commands

import (
	"archive/tar"
	"github.com/ZhangHang-z/vane/src/down"
	fp "github.com/ZhangHang-z/vane/src/fileparser"
)

func Install(args ...string) error {
	if len(args) == 0 {
		InstallFromJsonFile()
	}
}

func InstallFromJsonFile() {
	fj := fp.ParseJSONFile()
}
