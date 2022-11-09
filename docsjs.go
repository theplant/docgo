package docgo

import (
	"embed"

	"github.com/qor5/web"
)

//go:embed docjs/dist/*.*
var box embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := box.ReadFile("docjs/dist/docgo.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := box.ReadFile("docjs/dist/docgo.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
