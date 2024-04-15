package docgo

import (
	"bytes"
	"context"
	"github.com/theplant/htmlgo"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"strings"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithXHTML(),
	),
)

func Markdown(body string) htmlgo.HTMLComponent {
	return htmlgo.ComponentFunc(func(c context.Context) (r []byte, err error) {
		body = strings.Replace(body, "~", "`", -1)
		var buf bytes.Buffer
		err = md.Convert([]byte(body), &buf)
		if err != nil {
			return
		}
		root := htmlgo.RawHTML(buf.String())
		return root.MarshalHTML(c)
	})
}
