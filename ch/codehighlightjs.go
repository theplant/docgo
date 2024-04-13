package ch

import (
	"embed"

	"github.com/qor5/web/v3"
)

//go:embed codehighlightjs/dist
var assetsbox embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("codehighlightjs/dist/index.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("codehighlightjs/dist/style.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
