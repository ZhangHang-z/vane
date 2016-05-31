package commands

import (
	"fmt"
	"github.com/ZhangHang-z/vane/src/down/npm"
)

func Info(args ...string) error {
	if n := len(args); n == 0 && n > 1 {
		fmt.Println(GetHelpInfo())
		return nil
	}

	url := npm.GetNPMRegistryURL(args[0])
	repo, err := npm.NPMRegistryInit(url)
	if err != nil {
		return err
	}
	repo.PrintAvailableVersions()
	return nil
}
