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
