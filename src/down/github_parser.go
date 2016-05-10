package down

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/ZhangHang-z/vane/src/vane"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	GitHubPrefix = "https://github.com"
	ZipSuffix    = ".zip"
	TargzSuffix  = ".tar.gz"
	TagMaster    = "master"
)

var (
	ERR_HTTP_NOT_FOUND = errors.New("Package Not Found")
)

type Package struct {
	name    string
	version string
	source  string
}

// RsvGitHubAddr resolve github's package download address from raw address.
// parameter GPAP meaning: GitHub Package Address Protocol, such as "git://github.com/user/package^1.2.1".
func RsvGitHubAddr(GPAP string) string {
	u, err := url.Parse(GPAP)
	if err != nil {
		return ""
	}
	switch strings.ToLower(u.Scheme) {
	case "git":
		pkgName, pkgTag := GetPKGNameAndTag(strings.Split(u.Path, "^"))
		return GetGitHubFullURL(pkgName, pkgTag)
	case "http", "https":
		return GPAP
	default:
		u.Scheme = "http"
		return u.String()
	}
}

// GetGitHubFillURL generate a full github download address.
func GetGitHubFullURL(pkgName, pkgTag string) string {
	if pkgTag == "" {
		return fmt.Sprintf("%s%s/archive/master%s", GitHubPrefix, pkgName, ZipSuffix)
	}
	return fmt.Sprintf("%s%s/archive/%s%s", GitHubPrefix, pkgName, pkgTag, ZipSuffix)
}

func GetPKGNameAndTag(paths []string) (pkgName, pkgTag string) {
	lens := len(paths)
	if lens == 2 {
		return paths[0], paths[1]
	}
	if lens > 2 {
		return paths[0], paths[lens-1]
	}
	return "master", ""
}

// RevGithubPKG retrieve package from a given url.
func RevGithubPKG(url string) ([]byte, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if res.StatusCode == 404 {
		return nil, ERR_HTTP_NOT_FOUND
	}
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

// ExtractZip extract a zip file from a temporary file.
func ExtractZip(f *os.File) error {
	// Clean up temporary file. This defer function must before Close(),
	// the order of defer is first in last out, so if file not closed, remove file will failed.
	defer os.Remove(f.Name())
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}
	r, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return err
	}
	for _, exf := range r.File {
		if exf.FileInfo().IsDir() {
			err := os.MkdirAll(exf.Name, vane.ModeCommonDir)
			if err != nil {
				fmt.Printf("make directory: <%s> failed", err)
			}
			continue
		}

		r, err := exf.Open()
		defer r.Close()

		fmt.Printf("extract file: %s ...\n", exf.Name)
		if err != nil {
			log.Print(err)
			continue
		}

		dstF, err := os.Create(exf.Name)
		defer dstF.Close()
		if err != nil {
			fmt.Printf("create file %s failed.\n", exf.Name)
			continue
		}
		io.Copy(dstF, r)
	}
	return nil
}

// CreateTempFile create a temporary file by byte slice, return type *os.File.
func CreateTempFile(rawContents []byte) (*os.File, error) {
	tempfile, err := ioutil.TempFile("./", "")
	if err != nil {
		return nil, err
	}
	if _, err := tempfile.Write(rawContents); err != nil {
		return nil, err
	}
	return tempfile, nil
}

// GitHubDownloader retrieve file and extract.
func GitHubDownloader(url string) error {
	contents, err := RevGithubPKG(url)
	if err != nil {
		return err
	}

	f, err := CreateTempFile(contents)
	if err != nil {
		return err
	}

	return ExtractZip(f)
}
