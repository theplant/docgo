package docs

import . "github.com/theplant/docgo"

var Home = Doc(

	Markdown(`
## Overview

A common characteristic of a well-crafted API is that it’s easy to read and practically self-documenting. However, an API alone can’t convey important information like clear documentation does, such as:

- The overall architecture of a framework or package
- Relationships and dependencies between components in the API
- Boundary conditions, side effects, and errors that occur when using the API


`),
	Program("123").
		Language("swift"),
	Note("You must use the symbol’s absolute path for the page title of an extension file and include the name of the framework or package. DocC doesn’t support relative symbol paths in this context.").
		Title("Important"),

	Tip("Use a symbol’s full path to include it from elsewhere in the documentation hierarchy.").
		Title("Tip"),
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
).Title("Doc 1").AbstractText("this is an sample abstract for doc 1")
var article2 = Doc().Title("Doc 2").AbstractText("this is an sample abstract for doc 2")
var article3 = Doc().Title("Doc 3").AbstractText("this is an sample abstract for doc 3")
var article11 = Doc().Title("Doc 1.1").AbstractText("this is an sample abstract for doc 1.1")
