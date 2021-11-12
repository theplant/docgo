package docsrc

import (
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var UseWithHtmlGo = Doc(
	Markdown(`
## Use any html

Use together with [htmlgo](https://github.com/htmlgo), basically means you can write any html you like inside your doc, This gives you a whole lot of flexibility when writing documentation.

Take a look at this example:
`),
	ch.Code(MarkdownDifferenceSample),
).Title("Use with htmlgo").
	AbstractText("The ability to use any html code inside your documentation is pretty good")
