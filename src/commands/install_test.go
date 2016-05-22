package commands

import (
	"fmt"
	"testing"
)

func TestInstallFromJsonFile(t *testing.T) {
	err := Install.InstallFromJSONFile()
	fmt.Println(err)
}
