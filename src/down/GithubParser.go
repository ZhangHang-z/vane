package down

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	GitHubPrefix  = "https://github.com/"
	GitHubArchive = "archive"
	ZipSuffix     = ".zip"
	TargzSuffix   = ".tar.gz"
	TagMaster     = "master"
)

type Package struct {
	name    string
	version string
	source  string
}

/*
git://github.com/bower/bower
git://github.com/bower/bower^1.7.5


https://github.com/bower/bower/archive/master.zip
https://github.com/bower/bower/archive/v1.7.5.zip
*/

func RetrieveGithubPKG(url string) {
	res, err := http.Get(url)
	buf := bytes.NewBuffer(res.Body)

}

func ResolveGitHubAddress(pkgName, pkgTag string) string {
	urls := fmt.Sprintf("%s%s/%s/v%s%s", GitHubArchive, pkgName, GitHubArchive, pkgTag, ZipSuffix)
	return urls
}
