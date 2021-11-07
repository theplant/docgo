package docgo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/theplant/docgo"
	h "github.com/theplant/htmlgo"
	"github.com/theplant/testingutils"
)

func TestHello(t *testing.T) {
	var s = Doc().
		Home(documentASwiftFrameworkOrPackage).
		Header(h.Div()).
		Footer(h.Div()).
		Build()

	var tree = s.ArticleTree()

	expected := &ArticleNode{
		Title:    "Formatting Your Documentation Content",
		URI:      "formatting-your-documentation-content",
		Abstract: "Enhance your content’s presentation with special formatting and styling for text, links, and other page elements.",
		ChildNodes: []*ArticleNode{
			{
				Title: "Article 1",
				URI:   "article-1",
				ChildNodes: []*ArticleNode{
					{
						Title: "Article 1.1",
						URI:   "article-1-1",
					},
				},
			},
			{
				Title: "Article 2",
				URI:   "article-2",
			},
			{
				Title: "Article 3",
				URI:   "article-3",
			},
		},
	}

	diff := testingutils.PrettyJsonDiff(expected, tree)
	if len(diff) > 0 {
		t.Error(diff)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

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
	Note("You must use the symbol’s absolute path for the page title of an extension file and include the name of the framework or package. DocC doesn’t support relative symbol paths in this context.").
		Title("Important"),

	Tip("Use a symbol’s full path to include it from elsewhere in the documentation hierarchy.").
		Title("Tip"),

	Section(
		Group(
			article1,
			article2,
		).Title("Essentials"),

		Group().Title("Structure and Formatting"),
	).Title("Topics"),

	Section(
		Group(
			article3,
		).Title("Basics"),
	).Title("See Also"),
).Title("Formatting Your Documentation Content").
	AbstractText("Enhance your content’s presentation with special formatting and styling for text, links, and other page elements.")

var article1 = Article(Section(
	Group(
		article11,
	),
)).Title("Article 1")
var article2 = Article().Title("Article 2")
var article3 = Article().Title("Article 3")
var article11 = Article().Title("Article 1.1")
