package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func generate(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isHeading(line) {
			tocText := convert(line)
			result = append(result, tocText)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strings.Join(result, "")
}

func isHeading(line string) bool {
	return strings.HasPrefix(line, "#")
}

func convert(line string) string {
	var b bytes.Buffer
	s := strings.Split(line, " ")[1:]
	t := strings.Join(s, " ")

	b.WriteString("* [")
	b.WriteString(t)
	b.WriteString("]")
	b.WriteString("(#")

	for i, w := range s {
		b.WriteString(strings.ToLower(w))

		if i < len(s)-1 {
			b.WriteString("-")
		}
	}

	b.WriteString(")")
	b.WriteString("\n")

	return b.String()
}

func main() {
	app := cli.NewApp()

	app.Name = "mdtoc"
	app.Usage = "Generate a TOC from your Markdown file"
	app.Author = "maxbeizer"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		toc := generate(c.Args().Get(0))
		fmt.Println(toc)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
