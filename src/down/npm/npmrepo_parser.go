package npm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NpmRepo struct {
	Name     string
	Homepage string
	DistTags NpmDistTags `json:"dist-tags"`
	License  string
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

func GetNPMRegistry(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rawJson, _ := ioutil.ReadAll(resp.Body)
	var npmRepo NpmRepo
	err = json.Unmarshal(rawJson, &npmRepo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(npmRepo)
}
