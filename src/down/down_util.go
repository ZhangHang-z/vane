package down

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/ZhangHang-z/vane/src/dir"
	"github.com/ZhangHang-z/vane/src/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// RevContentsFromLink retrieve package from a given link.
func RveContentsFromLink(link string) ([]byte, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, errors.ERR_HTTP_NOT_FOUND
	}
	return ioutil.ReadAll(resp.Body)
}

// ExtractTarArchive extract the tar file from memory.
func ExtractTarArchive(contents []byte) error {
	r := bytes.NewReader(contents)
	tr := tar.NewReader(r)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // if end, exit.
		}
		if err != nil {
			return err
		}

		// if is a directory just create.
		fi := hdr.FileInfo()
		if fi.IsDir() {
			if err := os.MkdirAll(fi.Name(), dir.ModeCommonDir); err != nil {
				log.Printf("Create directory <%s> failed", fi.Name())
			}
			continue
		}

		// copy extracted contents to file.
		toFile, err := os.Create(hdr.Name)
		defer toFile.Close()
		if err != nil {
			log.Printf("Create file <%s> failed", hdr.Name)
			continue
		}

		fmt.Println("extract file:", hdr.Name)
		if _, err := io.Copy(toFile, tr); err != nil {
			log.Println(err)
		}
	}

	return nil
}
