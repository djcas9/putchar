package putchar

import (
	"github.com/codegangsta/cli"
)

func ReadCommand() cli.Command {
	return cli.Command{
		Name:        "add",
		ShortName:   "a",
		Usage:       "Add a git repository to your gill store.",
		Description: "Example: gill add https://github.com/mephux/gill.git",
		Action:      Read,
	}
}
