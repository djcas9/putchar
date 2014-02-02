package putchar

import (
	"github.com/codegangsta/cli"
)

func ReadFileCommand() cli.Command {
	return cli.Command{
		Name:        "file",
		ShortName:   "f",
		Usage:       "Read strings from file.",
		Description: "Example: putchar file /path/to/bin",
		Action:      ReadFile,
	}
}

func ReadInCommand() cli.Command {
	return cli.Command{
		Name:        "stdin",
		ShortName:   "-",
		Usage:       "Read strings from standard input.",
		Description: "Example: cat /path/to/bin | putchar -",
		Action:      ReadIn,
	}
}
