package docgo

import (
	"context"

	"github.com/iancoleman/strcase"
	h "github.com/theplant/htmlgo"
)

type ArticleNode struct {
	Title      string
	Abstract   string
	URI        string
	Article    *ArticleBuilder `json:"-"`
	ParentNode *ArticleNode    `json:"-"`
	ChildNodes []*ArticleNode
}

type ArticleBuilder struct {
	title        string
	uri          string
	abstractText string
	node         *ArticleNode
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

func (b *ArticleBuilder) URI(v string) (r *ArticleBuilder) {
	b.uri = v
	return b
}

func (b *ArticleBuilder) AbstractText(v string) (r *ArticleBuilder) {
	b.abstractText = v
	return b
}

func (b *ArticleBuilder) buildArticleNode() (r *ArticleNode) {
	if len(b.uri) == 0 {
		b.uri = strcase.ToKebab(b.title)
	}
	r = &ArticleNode{
		Title:    b.title,
		URI:      b.uri,
		Abstract: b.abstractText,
		Article:  b,
	}
	return
}

type contextKey int

const articleNodeKey contextKey = iota

func (b *ArticleBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	b.node = b.buildArticleNode()
	parent, ok := ctx.Value(articleNodeKey).(*ArticleNode)
	if ok {
		parent.ChildNodes = append(parent.ChildNodes, b.node)
		b.node.ParentNode = parent
	}
	ctx = context.WithValue(ctx, articleNodeKey, b.node)
	return b.tag.MarshalHTML(ctx)
}
