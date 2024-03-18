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
- Documents exists inside go code, means it can be distributed as go packages, so it won't be restricted with directory layout.
- Make developer focus on writing documentation, instead of worrying about document styles.

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
).Title("Introduction").
	Slug("/").
	AbstractText(`Write documentation with go code in a declarative way to create beautiful documentation`)

//go:embed assets/**.*
var Assets embed.FS
