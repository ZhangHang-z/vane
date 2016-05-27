package down

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/ZhangHang-z/vane/src/dir"
	"github.com/ZhangHang-z/vane/src/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
func ExtractTarArchive(tarName string, contents []byte) error {
	r := bytes.NewReader(contents)
	gz, _ := gzip.NewReader(r)
	tr := tar.NewReader(gz)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // if end, exit.
		}
		if err != nil {
			return err
		}

		nameList := strings.Split(hdr.Name, "/")
		nameList[0] = tarName
		hdr.Name = strings.Join(nameList, "/")

		err = dir.ParseDirAndMake(hdr.Name)
		if err != nil {
			log.Println(err)
		}

		// copy extracted contents to file.
		toFile, err := os.Create(hdr.Name)
		if err != nil {
			log.Printf("Create file <%s> failed", hdr.Name)
			continue
		}
		defer toFile.Close()
		if _, err := io.Copy(toFile, tr); err != nil {
			log.Println(err)
		}
	}

	return nil
}
