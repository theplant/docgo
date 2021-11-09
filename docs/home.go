package docs

import (
	. "github.com/theplant/docgo"
	ch "github.com/theplant/docgo/codehighlight"
	. "github.com/theplant/htmlgo"
)

var Home = Doc(

	Markdown(`
## Why use docgo

The developer community often use markdown or in Python world use [Sphinx](https://www.sphinx-doc.org) and [reStructuredText](https://docutils.sourceforge.io/rst.html) to create tech documentations, We think they are nice, But because of these reasons we created this package for the Go community:

- Work together with [snippetgo](https://github.com/sunfmin/snippetgo) to use real executable code as example code blocks in documentation, So that it won't be invalid or obsolete after code change, Since Go compiler will pick out the errors and you will have to fix the examples to make compiler happy, and the documentation is automatically updated.
- Write documentation in Go code not only you can still write markdown, But also you can access the flexibility of a programming language, to use components, and reuse parts that are duplicated.
- Documents exists inside go code, means it can be distributed as go packages, so it wont' be restricted with directory layout.
- Make developer focus on writing documentation, instead of worrying about document styles, the default document styles is following the [Swift DocC](https://github.com/apple/swift-docc) with minor changes. So we think it is good enough for most projects

`),
	P(
		Text("The following code is used to build this doc"),
		DocLink(HelloWorld),
	),
	ch.Code(HelloWorldSample).Language("go"),
	P(
		Text("Use the following code to boot up your doc app"),
		ch.Code(BootUpDevSample).Language("go"),
	),
).Title("docgo Documentation").
	URI("/").
	AbstractText(`Write documentation with go code in a declarative way to create beautiful documentation`).
	Tables(
		ContentTable(
			ContentGroup(
				article1,
				article2,
			).Title("Samples"),

			ContentGroup().Title("Structure and Formatting"),
		).Title("Topics"),

		ContentTable(
			ContentGroup(
				article3,
			).Title("Basics"),
		).Title("See Also"),
	)

var article1 = Doc().Tables(
	ContentTable(
		ContentGroup(
			article11,
		),
	),
).Title("Doc 1").
	AbstractText("this is an sample abstract for doc 1")

var article2 = Doc().Title("Doc 2").
	AbstractText("this is an sample abstract for doc 2").
	Tables(
		ContentTable(
			ContentGroup(
				HelloWorld,
			),
		),
	)

var article3 = Doc().Title("Doc 3").AbstractText("this is an sample abstract for doc 3")
var article11 = Doc().Title("Doc 1.1").AbstractText("this is an sample abstract for doc 1.1")

// @snippet_begin(HelloWorldSample)
var HelloWorld = Doc(
	Markdown(`
## Overview

Write some beautiful docs
`),
	Tip("This is quite important to learn"),
).
	Title("Hello World").
	AbstractText(
		"Hello world doc to describe how easy it is to create a doc",
	)

	// @snippet_end
