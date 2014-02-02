package putchar

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func hasArgs(args []string, expected int) bool {
	switch expected {
	case -1:
		if len(args) > 0 {
			return true
		} else {
			return false
		}
	default:
		if len(args) == expected {
			return true
		} else {
			return false
		}
	}
}

func ReadFile(c *cli.Context) {

	if !hasArgs(c.Args(), 1) {
		log.Fatal("Incorrect number of arguments. You must pass a file to read from.")
	}

	fmt.Println("Word?", c.Args()[0])

	file, err := os.Open(c.Args()[0])

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		b := scanner.Bytes()
		r, _ := utf8.DecodeRune(b)

		if utf8.ValidRune(r) {
			test := unicode.IsPrint(r)

			if test {
				fmt.Print("DEBUG: ", r, " ", string(r), " \n")
			} else {
				fmt.Println(r)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadIn(c *cli.Context) {
	fmt.Println("Word?")
}
