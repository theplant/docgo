package docgo

import (
	"context"

	. "github.com/theplant/htmlgo"
)

type DocLinkBuilder struct {
	doc *DocBuilder
	tag *HTMLTagBuilder
}

func DocLink(doc *DocBuilder) (r *DocLinkBuilder) {
	return &DocLinkBuilder{
		doc: doc,
		tag: A(),
	}
}

func (b *DocLinkBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.doc == nil || b.doc.node == nil {
		return
	}
	b.tag.Text(b.doc.node.Title).Href(b.doc.node.GetPageURL())
	return b.tag.MarshalHTML(ctx)
}
