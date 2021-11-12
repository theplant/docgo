package docsrc

import (
	"embed"

	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Home = Doc(
	Markdown(`
## Home title

home content
`),
	ch.Code(MyHelloWorldSample),
).
	Title("Home").Slug("/")

// @snippet_begin(MyHelloWorldSample)
func helloworld() {
	println("hello world")
}

// @snippet_end

//go:embed assets/**.*
var Assets embed.FS
