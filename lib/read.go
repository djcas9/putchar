package putchar

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"unicode"
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
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		b := scanner.Bytes()[0]
		//         fmt.Println(string(b))
		data := []rune(string(b))

		fmt.Println(data)

		for i := 0; i < len(data); i++ {
			test := unicode.IsPrint(data[i])

			if test {
				fmt.Print("DEBUG: ", data[i], " ", string(data[i]), " \n")
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
