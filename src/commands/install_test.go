package commands

import (
	"fmt"
	"testing"
)

// install package from json file.
func TestInstallFromJsonFile(t *testing.T) {
	err := Install.InstallFromJSONFile()
	fmt.Println(err)
}
