package docgo

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type SectionBuilder struct {
	title string
	tag   *h.HTMLTagBuilder
}

func Section(vs ...h.HTMLComponent) (r *SectionBuilder) {
	r = &SectionBuilder{
		tag: h.Div(vs...),
	}
	return
}

func (b *SectionBuilder) Title(v string) (r *SectionBuilder) {
	b.title = v
	return b
}

func (b *SectionBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	return b.tag.MarshalHTML(ctx)
}
