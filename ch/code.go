package ch

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type CodeBuilder struct {
	tag *h.HTMLTagBuilder
}

func Code(code string) (r *CodeBuilder) {
	r = &CodeBuilder{
		tag: h.Tag("highlightjs").Attr(":language", h.JSONString("go")),
	}
	r.Code(code)
	return
}

func (b *CodeBuilder) Code(v string) (r *CodeBuilder) {
	b.tag.Attr(":code", h.JSONString(v))
	return b
}

func (b *CodeBuilder) Language(v string) (r *CodeBuilder) {
	b.tag.Attr(":language", h.JSONString(v))
	return b
}

func (b *CodeBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
