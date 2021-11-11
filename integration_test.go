package docgo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/docs"
	h "github.com/theplant/htmlgo"
	"github.com/theplant/testingutils"
)

func TestAll(t *testing.T) {
	var s = New().
		Home(docs.Home).
		Header(h.Div()).
		Footer(h.Div()).
		Build()

	var tree = s.ArticleTree()

	expected := &DocNode{
		Title:    "docgo Documentation",
		Slug:     "/",
		URL:      "/",
		Abstract: "Use go code in declarative way to create beautiful documentation",
		ChildNodes: []*DocNode{
			{
				Title: "Doc 1",
				Slug:  "doc-1",
				URL:   "/doc-1",
				ChildNodes: []*DocNode{
					{
						Title: "Doc 1.1",
						Slug:  "doc-1-1",
						URL:   "/doc-1/doc-1-1",
					},
				},
			},
			{
				Title: "Doc 2",
				Slug:  "doc-2",
				URL:   "/doc-2",
			},
			{
				Title: "Doc 3",
				Slug:  "doc-3",
				URL:   "/doc-3",
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
