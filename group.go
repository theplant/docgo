package docgo

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type GroupBuilder struct {
	title string
	tag   *h.HTMLTagBuilder
}

func Group(vs ...h.HTMLComponent) (r *SectionBuilder) {
	r = &SectionBuilder{
		tag: h.Div(vs...),
	}
	return
}

func (b *GroupBuilder) Title(v string) (r *GroupBuilder) {
	b.title = v
	return b
}

func (b *GroupBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	return b.tag.MarshalHTML(ctx)
}
