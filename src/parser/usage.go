package parser

import (
	"fmt"
	"net/url"
)

var Usage = map[string]string{
	"install": "install packages from configuration file or by name that behind command",
	"info":    "info get pakcage infomation by name behind command",
}

var raw = `
Usage:
	vane <command> [, agrs] [, options]

Examples:
	vane install
	vant install pkgname1 pkgname2 --save

%s

Options:
	--save		save the package name into vane.json file
`

func GenHelpInfo() string {
	var s = "Commands:\n"
	for name, usage := range Usage {
		s = s + fmt.Sprintf("\t%s\t%s\n", name, usage)
	}
	return fmt.Sprintf(raw, s)
}

// IsDomainName inspect url is a domain name or a name of package.
func IsDomainName(args string) (bool, error) {
	u, err := url.Parse(args)
	if err != nil {
		return false, err
	}
	return !(u.Scheme == ""), nil
}
