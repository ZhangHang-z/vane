package down

import (
	"fmt"
	"testing"
)

var murl string = "https://github.com/koajs/koa/archive/2.0.0.zip"
var murl2 string = "git://github.com/jquery/jquery^2.2.3"

func _TestGitHubDownloader(t *testing.T) {
	err := GitHubDownloader(murl)
	if err != nil {
		t.Error(err)
	}
}

func _TestRsvGitHubAddr(t *testing.T) {
	githubURL := RsvGitHubAddr(murl2)
	fmt.Println(githubURL)
	err := GitHubDownloader(githubURL)
	if err != nil {
		t.Error(err)
	}
}
