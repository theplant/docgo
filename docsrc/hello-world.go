// @snippet_begin(HelloWorldSample)
package docsrc

import (
	. "github.com/theplant/docgo"
	ch "github.com/theplant/docgo/codehighlight"
	. "github.com/theplant/htmlgo"
)

var HelloWorld = Doc(
	Markdown(`
## Overview

Write some beautiful docs
`),
	Tip("This is quite important to learn"),

	H2("A new section"),
	P(Text("Doc of that section")),
	ch.Code("var Hello = 123").Language("go"),
).
	Title("Hello World").
	AbstractText(
		"Hello world doc to describe how easy it is to create a doc",
	)

// @snippet_end

// @snippet_begin(HelloWorldWithChildrenSample)
var HelloWorldWithChildren = Doc(
	Markdown(`
## Overview

Write some beautiful docs
`),
).Title("Hello World with children").
	AbstractText("Hello world with children is to describe how to add children docs").
	Tables(
		ChildrenTable(
			ContentGroup(
				HowToSayHello,
				HowToGoodBye,
			).Title("Essentials"),
		),
	)

var HowToSayHello = Doc(Markdown(`
## Say Hello

Hi, There

~~~go
var Hello = true
~~~

`)).Title("How to say hello")

var HowToGoodBye = Doc(Markdown(`
## Say Good Bye

Bye, There

`)).Title("How to say good bye")

// @snippet_end

// @snippet_begin(HowToSayHelloWithCodeBlockSample)
var HowToSayHelloWithCodeBlock = Doc(Markdown(`
## Say Hello

Hi, There

~~~
var Hello = true
~~~

`)).Title("How to say hello with code block")

// @snippet_end

var f = 1
