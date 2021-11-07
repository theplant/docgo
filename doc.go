package docgo

import (
	"context"
	"net/http"

	h "github.com/theplant/htmlgo"
)

type DocBuilder struct {
	home        *ArticleBuilder
	header      h.HTMLComponent
	footer      h.HTMLComponent
	articleTree *ArticleNode
}

func Doc() (r *DocBuilder) {
	r = &DocBuilder{}
	return
}

func (b *DocBuilder) Home(v *ArticleBuilder) (r *DocBuilder) {
	b.home = v
	return b
}

func (b *DocBuilder) Header(v h.HTMLComponent) (r *DocBuilder) {
	b.header = v
	return b
}

func (b *DocBuilder) Footer(v h.HTMLComponent) (r *DocBuilder) {
	b.footer = v
	return b
}

func (b *DocBuilder) ArticleTree() (r *ArticleNode) {
	return b.articleTree
}

type contextKey int

const articleNodeKey contextKey = iota

func (b *DocBuilder) Build() (r *DocBuilder) {
	ctx := context.WithValue(context.TODO(), articleNodeKey, &ArticleNode{})
	_, err := b.home.MarshalHTML(ctx)
	if err != nil {
		panic(err)
	}
	b.articleTree = ctx.Value(articleNodeKey).(*ArticleNode)
	return b
}

func (b *DocBuilder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}
