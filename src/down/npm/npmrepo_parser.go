package npm

import (
	"encoding/json"
	"fmt"
	"github.com/ZhangHang-z/vane/src/errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
)

func NPMRegistryInit(url string) (*NpmRepo, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, errors.ERR_HTTP_NOT_FOUND
	}

	rawJson, _ := ioutil.ReadAll(resp.Body)

	var npmRepo NpmRepo
	err = json.Unmarshal(rawJson, &npmRepo)
	if err != nil {
		return nil, err
	}

	return &npmRepo, nil
}

type NpmRepo struct {
	Name     string                 `json:"name"`
	Homepage string                 `json:"homepage"`
	DistTags NpmDistTags            `json:"dist-tags"`
	License  string                 `json:"license"`
	Versions map[string]NpmVersions `json:"versions"`
}

type NpmDistTags struct {
	Latest string
}

type NpmVersions struct {
	Dist NpmDist `json:"dist"`
}

type NpmDist struct {
	Tarball string
	Shasum  string
}

func (npm *NpmRepo) ChooseDist(version string) NpmDist {
	if nmpv, ok := npm.Versions[version]; ok {
		return nmpv.Dist
	}
	panic("Package has not this version")
}

func (npm *NpmRepo) GetLatestDist() NpmDist {
	latest := npm.DistTags.Latest
	return npm.ChooseDist(latest)
}

// PrintAvailableVersions print each version and omit experiment versions.
func (npm NpmRepo) PrintAvailableVersions() {
	fmt.Println("Available Versions:")
	if npm.Versions != nil {
		var list []string
		for name, _ := range npm.Versions {
			list = append(list, name)
		}
		sort.Strings(list)

		for _, named := range list {
			if FilterExperimentVersion(named) {
				fmt.Printf("\t- %s\n", named)
			}
		}
	}
}

// FilterExperimentVersion omit experiment version.
func FilterExperimentVersion(s string) bool {
	reg := regexp.MustCompile(`\d+\.\d+\.\d+.+`)
	return !reg.MatchString(s)
}

// GetNPMRegistryURL generate a url of npm registry.
func GetNPMRegistryURL(pkg string) string {
	return fmt.Sprint("https://registry.npmjs.org/", pkg)
}
