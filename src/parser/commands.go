package parser

import (
	"github.com/ZhangHang-z/vane/src/down"
)

var (
	AllCommands = map[string]interface{}{
		"install": down.GitHubDownloader,
		"help":    HelpString,
		"info":    "info...",
	}

	AllOptions = map[string]interface{}{
		"-v":        "version",
		"--version": "version",
	}
)

func IsValidCommand(command string) bool {
	_, hasAttr := AllCommands[command]
	if hasAttr {
		return true
	}
	return false
}
