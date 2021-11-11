package docsrc

import (
	. "github.com/theplant/docgo"
	ch "github.com/theplant/docgo/codehighlight"
	. "github.com/theplant/htmlgo"
)

// @snippet_begin(MarkdownDifferenceSample)
var MarkdownDifference = Doc(

	H2("Be aware the Go source code limitations"),

	P(
		Text("Since it's not possible to write"),
		Code("`"),
		Text("if you are writing inside a"),
		A().Text("go raw string literals").Href("https://golang.org/ref/spec#String_literals"),
		Text("which normally used to pass in "),
		Code("Markdown"),
		Text("func, So we have to replace it with"),
		Code("~"),
	),
	ch.Code(HowToSayHelloWithCodeBlockSample).Language("go"),
	Text("Which render a page like this"),
	DocLink(HowToSayHelloWithCodeBlock),
).Title("The different with Github Flavored Markdown").
	Tables(
		ChildrenTable(
			ContentGroup(
				HowToSayHelloWithCodeBlock,
			).Title("Samples"),
		),

		RelatedTable(
			ContentGroup(
				UseWithHtmlGo,
			),
		),
	)

	// @snippet_end
