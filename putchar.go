package main

import (
	"github.com/codegangsta/cli"
	"github.com/mephux/putchar/lib"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "putchar"
	app.Usage = "Print Characters From Anything."
	app.Version = putchar.VERSION
	app.Commands = []cli.Command{
		putchar.ReadFileCommand(),
		putchar.ReadInCommand(),
	}

	app.Run(os.Args)
}
