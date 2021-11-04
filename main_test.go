package docgo_test

import (
	"net/http"
	"testing"

	h "github.com/theplant/htmlgo"
)

func TestHello(t *testing.T) {
	var s = Doc().
		Home(documentASwiftFrameworkOrPackage).
		Header(h.Div()).
		Footer()

	http.Handler(s).ServeHTTP(w, r)
}

var documentASwiftFrameworkOrPackage = Article(
	Markdown(`
## Overview

A common characteristic of a well-crafted API is that it’s easy to read and practically self-documenting. However, an API alone can’t convey important information like clear documentation does, such as:

- The overall architecture of a framework or package
- Relationships and dependencies between components in the API
- Boundary conditions, side effects, and errors that occur when using the API


`),
	Code("123").
		Language("swift"),
	Warning("You must use the symbol’s absolute path for the page title of an extension file and include the name of the framework or package. DocC doesn’t support relative symbol paths in this context.").
		Title("Important"),

	Tip("Use a symbol’s full path to include it from elsewhere in the documentation hierarchy.").
		Title("Tip"),

	Section(
		Group(
			addStructureToYourDocumentationPages,
		).Title("Essentials"),

		Group().Title("Structure and Formatting"),
	).Title("Topics"),

	Section(
		Group(
			SlothCreatorBuildingDocC,
		).Title("Basics"),
	).Title("See Also"),
).Title("Formatting Your Documentation Content").
	Abstract("Enhance your content’s presentation with special formatting and styling for text, links, and other page elements.")

var addStructureToYourDocumentationPages = Article()
var SlothCreatorBuildingDocC = Article()
