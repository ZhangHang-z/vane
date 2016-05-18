package commands

type CmdFunc func(args ...string) error

type Commander interface {
	Execute(args ...string) error
	RollBack() error
}

var VaneCommmans = map[string]Commander{
	"install":   Install,
	"info":      Info,
	"uninstall": uninstall,
}

func IsValidCommand(cmd string) (string, bool) {
	command, ok := VaneCommands[cmd]
	if ok {
		return command.Usage, true
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
