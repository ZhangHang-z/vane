package vflag

import (
	"flag"
	"fmt"
	"testing"
)

func TestVaneBoolValue(t *testing.T) {
	var save bool
	VaneBoolValue(&save, "save", false, "--save will save the command.")
	flag.Parse()
}
