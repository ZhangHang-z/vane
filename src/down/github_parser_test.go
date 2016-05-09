package down

import (
	"testing"
)

var url string = "https://github.com/ZhangHang-z/vane/archive/master.zip"

func TestGitHubDownloader(t *testing.T) {
	err := GitHubDownloader(url)
	if err != nil {
		t.Error(err)
	}

}
