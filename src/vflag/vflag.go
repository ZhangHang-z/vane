package vflag

import (
	"flag"
	"fmt"
	"strconv"
)

const (
	ExitGracefully = iota
)

type vaneBoolValue bool

func (p *vaneBoolValue) String() string { return fmt.Sprint("%v", *p) }

// If a Value has an IsBoolFlag method returning true, then -flag eqivalent to -flag=true.
func (p *vaneBoolValue) IsBoolFlag() bool { return true }

func (p *vaneBoolValue) Get(value string) interface{} { return bool(*p) }

func (p *vaneBoolValue) Set(value string) error {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	*p = vaneBoolValue(v)
	return nil
}

func newVaneBoolValue(value bool, p *bool) *vaneBoolValue {
	*p = value
	return (*vaneBoolValue)(p)
}

func VaneBoolValue(p *bool, name string, value bool, usage string) {
	flag.CommandLine.Var(newVaneBoolValue(value, p), name, usage)
}
