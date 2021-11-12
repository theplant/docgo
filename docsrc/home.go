package docsrc

import (
	"embed"

	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
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
	Markdown(`
Run this script to install ~docgo~ binary

~~~
$ go install github.com/theplant/docgo/x/docgo@main 
~~~

Go to a go package that use go modules with ~go.mod~ in the directory. and run:

~~~
$ docgo
~~~

It will output something like this:

~~~
docsrc/assets/logo.png generated
docsrc/build/main.go generated
docsrc/dev/main.go generated
docsrc/dev.sh generated
docsrc/examples-generated.go generated
docsrc/home.go generated
Done

cd docsrc && ./dev.sh to start your doc
~~~

Run ~cd docsrc && ./dev.sh~ and access http://localhost:8800 to see generated first doc.

The ~./dev.sh~ script is using [entr](https://eradman.com/entrproject/) to do auto restart server after you edit any go file. So make sure to have that installed if you haven't

~~~
$ brew install entr
~~~

Then you can go back to the ~docsrc~ directory to edit and create go files to update docs.

`),
	P(
		Text("The following code is used to build this doc"),
		DocLink(HelloWorld),
	),
	ch.Code(HelloWorldSample).Language("go"),
	P(
		Markdown("Use the following code to boot up your doc app, Suppose you have already created a `Home` Doc in docsrc package."),
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
				GithubPagesIntegration,
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

//go:embed assets/**.*
var Assets embed.FS
