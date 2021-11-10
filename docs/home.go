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
	H2("Getting Started"),
	P(
		Text("The following code is used to build this doc"),
		DocLink(HelloWorld),
	),
	ch.Code(HelloWorldSample).Language("go"),
	P(
		Markdown("Use the following code to boot up your doc app, Suppose you have already created a `Home` Doc in docs package."),
		ch.Code(BootUpDevSample).Language("go"),
	),

	Markdown(`
## Children Docs

Use ~ChildrenTable(...)~ to put other docs into current doc page children, The doc url will reflect the hierarchy, children docs url will contain parent doc slug
`),
	ch.Code(HelloWorldWithChildrenSample).Language("go"),
	Text("Check out this link to see how it works"),
	DocLink(HelloWorldWithChildren),
).Title("docgo Documentation").
	Slug("/").
	AbstractText(`Write documentation with go code in a declarative way to create beautiful documentation`).
	Tables(
		ChildrenTable(
			ContentGroup(
				UseWithHtmlGo,
				MarkdownDifference,
			).Title("Essentials"),

			ContentGroup(
				HelloWorld,
				HelloWorldWithChildren,
			).Title("Samples"),

			ContentGroup().Title("Structure and Formatting"),
		).Title("Topics"),

		RelatedTable().Title("See Also"),
	)

var related1 = []*DocBuilder{
	MarkdownDifference,
	UseWithHtmlGo,
}

func relatedDocsWithout(current *DocBuilder, group []*DocBuilder) (r []*DocBuilder) {
	for _, d := range group {
		if d == current {
			continue
		}
		r = append(r, d)
	}
	return
}
