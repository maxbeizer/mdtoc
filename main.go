package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maxbeizer/mdtoc/parser"
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
		if parser.IsHeading(line) {
			tocText := convert(line)
			result = append(result, tocText)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strings.Join(result, "")
}

func convert(line string) string {
	var b bytes.Buffer
	split := strings.Split(line, " ")
	h := split[0]
	s := split[1:]
	t := strings.Join(s, " ")

	if len(h) == 1 {
		return ""
	}

	b = parser.WriteDepth(b, h)
	b = parser.WriteLinkText(b, t)
	b = parser.WriteLink(b, s)
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
