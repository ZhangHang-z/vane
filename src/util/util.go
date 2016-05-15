package util

import (
	"net/url"
	"os"
)

const (
	ModeCommonDir  os.FileMode = 0775 // unix common user's directory default permission.
	ModeCommonFile os.FileMode = 0664 // unix common user's file default permission.
)

// IsDomainName inspect url is a domain name or a name of package.
func IsDomainName(args string) (bool, error) {
	u, err := url.Parse(args)
	if err != nil {
		return false, err
	}
	return !(u.Scheme == ""), nil
}
