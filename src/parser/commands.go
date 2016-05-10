package parser

import (
	"github.com/ZhangHang-z/vane/src/down"
)

var (
	AllCommands = map[string]interface{}{
		"install": down.GitHubDownloader,
		"help":    HelpString,
	}

	AllOptions = map[string]interface{}{
		"-v":        "version",
		"--version": "version",
	}
)
