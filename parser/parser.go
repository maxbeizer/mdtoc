package parser

import (
	"bytes"
	"strings"
)

func IsHeading(line string) bool {
	return strings.HasPrefix(line, "#")
}

func WriteDepth(b bytes.Buffer, h string) bytes.Buffer {
	return b
}

func WriteLinkText(b bytes.Buffer, t string) bytes.Buffer {
	b.WriteString("* [")
	b.WriteString(t)
	b.WriteString("]")
	return b
}

func WriteLink(b bytes.Buffer, s []string) bytes.Buffer {
	b.WriteString("(#")

	for i, w := range s {
		b.WriteString(strings.ToLower(w))

		if i < len(s)-1 {
			b.WriteString("-")
		}
	}

	b.WriteString(")")
	return b
}
