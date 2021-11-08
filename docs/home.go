package docs

import . "github.com/theplant/docgo"

var Home = Doc(

	Markdown(`
## Why use docgo

The developer community often use markdown or in Python world use [Sphinx](https://www.sphinx-doc.org) and [reStructuredText](https://docutils.sourceforge.io/rst.html), We created this to solve these problems:

- Work together with [snippetgo](https://github.com/sunfmin/snippetgo) to use real executable code as examples in documentation, So that it won't be invalid or obsolete after code change, Since Go compiler will pick out the errors and you will have to fix the examples to make compiler happy, and the documentation is automatically updated.
- Write documentation in Go code not only you can still write markdown, But also you can access the flexibility of a programming language, to use components, and reuse parts that are duplicated.

`),
).
	Tables(

		ContentTable(
			ContentGroup(
				article1,
				article2,
			).Title("Essentials"),

			ContentGroup().Title("Structure and Formatting"),
		).Title("Topics"),

		ContentTable(
			ContentGroup(
				article3,
			).Title("Basics"),
		).Title("See Also"),
	).
	Title("docgo Documentation").
	URI("/").
	AbstractText(`Write documentation with go code in a declarative way to create beautiful documentation`)

var article1 = Doc().Tables(
	ContentTable(
		ContentGroup(
			article11,
		),
	),
).Title("Doc 1").
	AbstractText("this is an sample abstract for doc 1")

var article2 = Doc().Title("Doc 2").AbstractText("this is an sample abstract for doc 2")
var article3 = Doc().Title("Doc 3").AbstractText("this is an sample abstract for doc 3")
var article11 = Doc().Title("Doc 1.1").AbstractText("this is an sample abstract for doc 1.1")
