package docgo

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	. "github.com/theplant/htmlgo"
)

type DocNode struct {
	Title      string
	Abstract   string
	Slug       string
	URL        string
	Doc        *DocBuilder `json:"-"`
	ParentNode *DocNode    `json:"-"`
	ChildNodes []*DocNode
}

func (n *DocNode) AddChild(child *DocNode) {
	n.ChildNodes = append(n.ChildNodes, child)
	child.ParentNode = n
	child.URL = filepath.Join(n.URL, child.Slug)
}

func (n *DocNode) GetPageURL() (r string) {
	if n.URL == "/" {
		return "index.html"
	}
	return strings.TrimLeft(fmt.Sprintf("%s.html", n.URL), "/")
}

type DocBuilder struct {
	title        string
	slug         string
	abstractText string
	node         *DocNode
	children     []HTMLComponent
	tables       []HTMLComponent
}

func Doc(vs ...HTMLComponent) (r *DocBuilder) {
	r = &DocBuilder{
		children: vs,
	}
	return
}

func (b *DocBuilder) Title(v string) (r *DocBuilder) {
	b.title = v
	return b
}

func (b *DocBuilder) Slug(v string) (r *DocBuilder) {
	b.slug = v
	return b
}

func (b *DocBuilder) AbstractText(v string) (r *DocBuilder) {
	b.abstractText = v
	return b
}

func (b *DocBuilder) GetPageTitle() (r string) {
	items := parentDocNodes(b)
	var segs []string
	for _, item := range items {
		segs = append(segs, item.Title)
	}
	return strings.Join(segs, " - ")
}

func (b *DocBuilder) ContentGroupItem(ctx context.Context) (r HTMLComponent) {
	collectChildren, _ := ctx.Value(collectChildrenKey).(bool)
	if collectChildren {
		b.buildArticleNode()
	}
	_, _ = b.MarshalHTML(ctx) // for build up doc node tree

	if b.node == nil {
		return
	}

	r = Div(
		A(
			Div(
				docIcon,
			).Class("w-4 h-4 mt-2 mr-4 text-gray-500 fill-current"),
			Span(b.title),
		).Class("flex").Href(b.node.GetPageURL()),
		If(len(b.abstractText) > 0, Div(
			Div().Text(b.abstractText),
		).Class("ml-8")),
	).Class("mt-4 sm:mt-8")
	return
}

func (b *DocBuilder) Tables(vs ...HTMLComponent) (r *DocBuilder) {
	b.tables = vs
	return b
}

func (b *DocBuilder) buildArticleNode() (r *DocNode) {
	if b.node != nil {
		return b.node
	}
	if len(b.slug) == 0 {
		b.slug = strcase.ToKebab(b.title)
	}
	b.node = &DocNode{
		Title:    b.title,
		Slug:     b.slug,
		URL:      filepath.Join("/", b.slug),
		Abstract: b.abstractText,
		Doc:      b,
	}
	return b.node
}

type contextKey int

const (
	articleNodeKey contextKey = iota
	collectChildrenKey
)

func (b *DocBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	collectChildren, _ := ctx.Value(collectChildrenKey).(bool)

	b.node = b.buildArticleNode()
	parent, ok := ctx.Value(articleNodeKey).(*DocNode)

	if collectChildren {
		if ok {
			parent.AddChild(b.node)
		}
	}
	ctx = context.WithValue(ctx, articleNodeKey, b.node)

	return Div(
		Main(
			H1(b.title),
			Div(
				Div(
					// abstract
					If(len(b.abstractText) > 0,
						Div(
							Text(b.abstractText),
						).Class("mb-8 text-xl font-normal"),
					),
					Div(
						b.children...,
					).Class("border-t"),
				).Class("sm:w-9/12"),
				Div(
					Div(
						Text("On This Page"),
						RawHTML("<toc></toc>"),
					).Class("sticky top-4"),
				).Class("ml-4 sm:w-3/12 font-medium text-base hidden sm:block text-gray-600"),
			).Class("sm:flex mt-8 mb-16"),
		).Class("lg:max-w-5xl mx-auto px-6 sm:px-8"),

		If(len(b.tables) > 0,
			Div(b.tables...).Class("pb-16 bg-gray-50"),
		),
	).MarshalHTML(ctx)
}

var docIcon = RawHTML(`
<svg aria-hidden="true" class=""
viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
<path
d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"></path>
</svg>
`)
