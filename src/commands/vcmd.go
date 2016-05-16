package commands

type CmdFunc func(args ...string) error

var VaneCommmansMap = map[string]CmdFunc{
	"install":   Install,
	"info":      Info,
	"uninstall": uninstall,
}

var VaneCommandUsage = []VaneCommand{
	{"install", "install the packages."},
	{"info", "get package infomations."},
	{"uninstall", "uninstall saved package from vane package directory."},
}

type VaneCommand struct {
	Name        string
	Description string
}

func IsValidCommand(command string) (string, bool) {
	cmd, hasAttr := VaneCommandUsage[command]
	if hasAttr {
		return cmd.Description, true
	}
	return "", false
}

var HelpString = `
Usage:
	vane <command> [, agrs] [, options]

Examples:
	vane install
	vant install pkgname1 pkgname2 --save

Commands:
	install		install the packages.

Options:
	--save		save the package name into vane.json file.
`
