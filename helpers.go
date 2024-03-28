package snaparser_server

import (
	"fmt"
	"io"
	"strings"

	"github.com/vanillaiice/snaparser/parser"
)

// mediaTypeText represents a media that is text
const mediaTypeText = "TEXT"

// replaceSlash checks if a string contains '/' characters.
// If yes, they are replaced with '-'.
func replaceSlash(s *string) {
	if strings.Contains(*s, "/") {
		*s = strings.ReplaceAll(*s, "/", "-")
	}
}

// writeContent writes the content of a parsed message to a writer.
func writeContent(w io.Writer, c []*parser.Content) (err error) {
	for i := len(c) - 1; i >= 0; i-- {
		if c[i].MediaType != mediaTypeText {
			continue
		}

		s := fmt.Sprintf("%s (%s): %s\n", c[i].From, c[i].Created, c[i].Content)
		if _, err = io.WriteString(w, s); err != nil {
			return
		}
	}

	return
}
