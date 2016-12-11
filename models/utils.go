package models

import (
	"strings"

	"github.com/russross/blackfriday"
)

func processMarkdown(text string) string {
	html := blackfriday.MarkdownCommon([]byte(text))
	strHTML := string(html)
	return strings.Replace(strHTML, "\n", "", -1)
}
