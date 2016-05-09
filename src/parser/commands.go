package parser

import (
	"github.com/ZhangHang-z/vane/src/down"
)

var (
	Commands = map[string]interface{}{
		"install": down.GitHubDownloader,
	}
)
