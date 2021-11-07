package docgo

import (
	"context"

	. "github.com/theplant/htmlgo"
)

type ContentTableBuilder struct {
	title    string
	children []HTMLComponent
}

func ContentTable(vs ...HTMLComponent) (r *ContentTableBuilder) {
	r = &ContentTableBuilder{
		children: vs,
	}
	return
}

func (b *ContentTableBuilder) Title(v string) (r *ContentTableBuilder) {
	b.title = v
	return b
}

func (b *ContentTableBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	return Section(
		Div(
			H2(b.title),
			Components(b.children...),
		).Class("lg:max-w-5xl mx-auto px-10 mt-5"),
	).Class("bg-gray-50 pt-8").MarshalHTML(ctx)
}
