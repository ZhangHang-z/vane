package commands

import (
	"fmt"
	"net/url"
)

type CmdFunc func(args ...string) error

type Commander interface {
	Execute(args ...string) error
}

var VaneCommands = map[string]Commander{
	"install": Install,
}

var HelpString = `
Usage:
	vane <command> [, agrs] [, options]

Examples:
	vane install
	vant install pkgname1 pkgname2 --save

Commands:
	install		install the packages.

Options:
	--save		save the package name into vane.json file.
`

// IsDomainName inspect url is a domain name or a name of package.
func IsDomainName(args string) (bool, error) {
	u, err := url.Parse(args)
	if err != nil {
		return false, err
	}
	return !(u.Scheme == ""), nil
}

func PrintHelpStringAll() {
	fmt.Println(HelpString)
}
