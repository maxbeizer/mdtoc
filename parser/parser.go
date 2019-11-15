package parser

import (
	"bytes"
	"regexp"
	"strings"
)

func IsHeading(line string) bool {
	return strings.HasPrefix(line, "#")
}

func WriteDepth(b bytes.Buffer, h string) bytes.Buffer {
	var depth []string
	hashes := strings.Split(h, "")
	if len(hashes) < 2 {
		return b
	}
	hs := hashes[2:]

	for index := 0; index < len(hs); index++ {
		depth = append(depth, "  ")
	}
	b.WriteString(strings.Join(depth, ""))
	return b
}

func WriteLinkText(b bytes.Buffer, t string) bytes.Buffer {
	b.WriteString("* [")
	b.WriteString(t)
	b.WriteString("]")
	return b
}

func WriteLink(b bytes.Buffer, s []string) bytes.Buffer {
	re := regexp.MustCompile(`[\/:\(\),]`) // GitHub Markdown removes `/`, `()`, `,`, and `:`
	b.WriteString("(#")

	for i, w := range s {
		w = re.ReplaceAllString(w, "") // Remove special characters
		b.WriteString(strings.ToLower(w))

		if i < len(s)-1 {
			b.WriteString("-")
		}
	}

	b.WriteString(")")
	return b
}
