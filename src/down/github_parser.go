package down

import (
	//"archive/tar"
	"archive/zip"
	"fmt"
	"github.com/ZhangHang-z/vane/src/vane"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func ResolveGitHubAddress(pkgName, pkgTag string) string {
	return fmt.Sprintf("%s%s/%s/v%s%s", GitHubArchive, pkgName, GitHubArchive, pkgTag, ZipSuffix)
}

func RetrieveGithubPKG(url string) ([]byte, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

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

func GitHubDownloader(url string) error {
	contents, err := RetrieveGithubPKG(url)
	if err != nil {
		return err
	}
	f, err := CreateTempFile(contents)
	if err != nil {
		return err
	}
	return ExtractZip(f)
}
