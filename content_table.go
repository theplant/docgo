package docgo

import (
	"context"

	. "github.com/theplant/htmlgo"
)

type ContentTableBuilder struct {
	title           string
	children        []HTMLComponent
	collectChildren bool
}

func ChildrenTable(vs ...HTMLComponent) (r *ContentTableBuilder) {
	r = &ContentTableBuilder{
		title:           "Topics",
		children:        vs,
		collectChildren: true,
	}
	return
}

func RelatedTable(vs ...HTMLComponent) (r *ContentTableBuilder) {
	r = &ContentTableBuilder{
		title:    "See Also",
		children: vs,
	}
	return
}

func (b *ContentTableBuilder) Title(v string) (r *ContentTableBuilder) {
	b.title = v
	return b
}

func (b *ContentTableBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if len(b.children) == 0 {
		return
	}
	ctx = context.WithValue(ctx, collectChildrenKey, b.collectChildren)
	return Section(
		Div(
			H2(b.title),
			Components(b.children...),
		).Class("lg:max-w-5xl mx-auto px-6 sm:px-10 mt-5"),
	).Class("py-4").MarshalHTML(ctx)
}
