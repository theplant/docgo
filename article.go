package docgo

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type ArticleNode struct {
	Title      string
	Abstract   h.HTMLComponent
	URI        string
	ChildNodes []*ArticleNode
}

type ArticleBuilder struct {
	title        string
	abstractText string
	tag          *h.HTMLTagBuilder
}

func Article(vs ...h.HTMLComponent) (r *ArticleBuilder) {
	r = &ArticleBuilder{
		tag: h.Div(vs...),
	}
	return
}

func (b *ArticleBuilder) Title(v string) (r *ArticleBuilder) {
	b.title = v
	return b
}

func (b *ArticleBuilder) AbstractText(v string) (r *ArticleBuilder) {
	b.abstractText = v
	return b
}

func (b *ArticleBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	return b.tag.MarshalHTML(ctx)
}
