package docgo

import (
	"context"
	"strings"

	"github.com/shurcooL/github_flavored_markdown"
	"github.com/theplant/htmlgo"
)

func Markdown(body string) htmlgo.HTMLComponent {
	return htmlgo.ComponentFunc(func(c context.Context) (r []byte, err error) {
		body = strings.Replace(body, "~", "`", -1)
		root := htmlgo.RawHTML(github_flavored_markdown.Markdown([]byte(body)))
		return root.MarshalHTML(c)
	})
}
