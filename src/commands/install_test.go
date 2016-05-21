package commands

import (
	"fmt"
	"testing"
)

func TestInstallFromJsonFile(t *testing.T) {
	err := Install.InstallFromJsonFile()
	fmt.Println(err)
}
