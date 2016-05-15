package down

import (
	"net/url"
)

// IsDomainName inspect url is a domain name or a name of package.
func IsDomainName(args string) (bool, error) {
	u, err := url.Parse(args)
	if err != nil {
		return false, err
	}
	return !(u.Scheme == ""), nil
}
