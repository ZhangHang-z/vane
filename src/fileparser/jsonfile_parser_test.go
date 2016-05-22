package fileparser

import (
	"fmt"
	"testing"
)

func TestReadJSONFile(t *testing.T) {
	var vj = new(VaneJSON)
	err := vj.Read()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("unmarshaled json struct:", vj)
}

func TestGetVersionInfo(t *testing.T) {
	v1 := GetVersionInfo("~1.1.2")
	v2 := GetVersionInfo("~1.10.2")
	v3 := GetVersionInfo("~111.2")
	v4 := GetVersionInfo("~1.13.2")
	v5 := GetVersionInfo("~1.0.0-alpha.7")
	fmt.Println("package version info: ", v1, v2, v3, v4, v5)
}
