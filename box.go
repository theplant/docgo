package docgo

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type BoxBuilder struct {
	tag      *h.HTMLTagBuilder
	text     string
	titleTag *h.HTMLTagBuilder
}

func Note(text string) (r *BoxBuilder) {
	r = &BoxBuilder{
		tag: h.Tag("div").Class("mt-6 p-4 rounded-2xl bg-gray-50 border border-gray-400 shadow"),
	}
	r.titleTag = h.Tag("label").Text("Note").Class("text-gray-700 font-normal")
	return
}

func Important(text string) (r *BoxBuilder) {
	r = &BoxBuilder{
		tag: h.Tag("div").Class("mt-6 p-4 rounded-2xl bg-yellow-50 border border-yellow-700 shadow"),
	}
	r.titleTag = h.Tag("label").Text("Important").Class("text-yellow-700 font-normal")
	return
}

func Deprecated(text string) (r *BoxBuilder) {
	r = &BoxBuilder{
		tag: h.Tag("div").Class("mt-6 p-4 rounded-2xl bg-pink-50 border border-yellow-700 shadow"),
	}
	r.titleTag = h.Tag("label").Text("Important").Class("text-yellow-700 font-normal")
	return
}

func Experiment(text string) (r *BoxBuilder) {
	r = &BoxBuilder{
		tag: h.Tag("div").Class("mt-6 p-4 rounded-2xl bg-purple-50 border border-purple-800 shadow"),
	}
	r.titleTag = h.Tag("label").Text("Experiment").Class("text-purple-800 font-normal")
	return
}

func Tip(text string) (r *BoxBuilder) {
	r = &BoxBuilder{
		tag: h.Tag("div").Class("mt-6 p-4 rounded-2xl bg-green-50 border border-green-700 shadow"),
	}
	r.titleTag = h.Tag("label").Text("Tip").Class("text-green-700 font-normal")
	return
}

func (b *BoxBuilder) Title(v string) (r *BoxBuilder) {
	b.titleTag.Text(v)
	return b
}

func (b *BoxBuilder) Children(vs ...h.HTMLComponent) (r *BoxBuilder) {
	b.tag.Children(vs...)
	return b
}

func (b *BoxBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	b.tag.Children(
		b.titleTag,
		h.P().Text(b.text).Class("mt-2"),
	)
	return b.tag.MarshalHTML(ctx)
}
