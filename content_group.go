package docgo

import (
	"context"

	. "github.com/theplant/htmlgo"
)

type ContentGroupBuilder struct {
	title string
	docs  []*DocBuilder
}

func ContentGroup(vs ...*DocBuilder) (r *ContentGroupBuilder) {
	r = &ContentGroupBuilder{
		docs: vs,
	}
	return
}

func (b *ContentGroupBuilder) Title(v string) (r *ContentGroupBuilder) {
	b.title = v
	return b
}

func (b *ContentGroupBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	if len(b.docs) == 0 {
		return nil, nil
	}
	var vs []HTMLComponent
	for _, d := range b.docs {
		vs = append(vs, d.ContentGroupItem(ctx))
	}

	return Div(
		Div(
			H3(b.title),
		).Class("sm:w-1/4 pb-4 border-b sm:border-none"),
		Div(
			vs...,
		).Class("sm:w-3/4"),
	).Class("sm:flex sm:border-t mt-8").MarshalHTML(ctx)
}
