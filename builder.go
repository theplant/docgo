package docgo

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/goplaid/web"
	"github.com/theplant/docgo/codehighlight"
	. "github.com/theplant/htmlgo"
)

type Builder struct {
	home        *DocBuilder
	header      HTMLComponent
	footer      HTMLComponent
	articleTree *DocNode
	mux         *http.ServeMux
	mounts      []*mount
	builder     *web.Builder
}

func New() (r *Builder) {
	r = &Builder{
		builder: web.New(),
	}
	return
}

func (b *Builder) Home(v *DocBuilder) (r *Builder) {
	b.home = v
	return b
}

func (b *Builder) Header(v HTMLComponent) (r *Builder) {
	b.header = v
	return b
}

func (b *Builder) Footer(v HTMLComponent) (r *Builder) {
	b.footer = v
	return b
}

func (b *Builder) ArticleTree() (r *DocNode) {
	return b.articleTree
}

func (b *Builder) Build() (r *Builder) {
	ctx := context.TODO()
	_, err := b.home.MarshalHTML(ctx)
	if err != nil {
		panic(err)
	}
	b.articleTree = b.home.node
	b.addToMounts(b.articleTree)
	sort.Sort(uriSorter(b.mounts))
	b.initMuxWithStatic()
	for _, m := range b.mounts {
		fmt.Println("Mounting", m.path)
		b.mux.HandleFunc(m.path, m.f)
	}
	return b
}

func (b *Builder) layout(body *DocBuilder) (r HTMLComponent) {
	return HTML(
		Head(
			Link("/index.css").Rel("stylesheet").Type("text/css"),
			Script("").Attr("defer", true).Src("/index.js"),
		),
		Body(
			Div(
				b.header,
				b.navigation(body),
				body,
				b.footer,
			).Id("app").Attr("v-cloak", true),
		),
	)
}

func (b *Builder) navigation(doc *DocBuilder) (r HTMLComponent) {
	if doc.node == nil {
		return
	}
	var items = []*DocNode{doc.node}
	var current = doc.node
	for current.ParentNode != nil {
		current = current.ParentNode
		items = append(items, current)
	}

	content := Div().Attr("aria-label", "Breadcrumbs").
		Class("flex list-none lg:max-w-5xl mx-auto px-10")

	for i := len(items) - 1; i >= 0; i-- {
		content.AppendChildren(
			Div(
				If(i < (len(items)-1),
					Div(
						arrowIcon,
					).Class("w-3 m-2 flex fill-current text-gray-500"),
				),
				A().Href(items[i].AbsoluteURI).Text(items[i].Title).
					Class("text-gray-50"),
			).Class("inline-flex"),
		)
	}

	return Nav(content).Class("bg-gray-700 py-3 text-base font-normal mb-8")
}

var startTime = time.Now()

type mount struct {
	path string
	f    http.HandlerFunc
}

type uriSorter []*mount

func (s uriSorter) Len() int {
	return len(s)
}

func (s uriSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s uriSorter) Less(i, j int) bool {
	return len(s[i].path) > len(s[j].path)
}

func (b *Builder) addToMounts(node *DocNode) {
	// println(node.AbsoluteURI)
	b.mounts = append(b.mounts,
		&mount{node.AbsoluteURI, func(w http.ResponseWriter, r *http.Request) {
			err := Fprint(w, b.layout(node.Doc), context.TODO())
			if err != nil {
				panic(err)
			}
		}})

	for _, c := range node.ChildNodes {
		b.addToMounts(c)
	}
}

func (b *Builder) initMuxWithStatic() {
	if b.mux == nil {
		b.mux = http.NewServeMux()

		b.mux.Handle("/index.js",
			b.builder.PacksHandler("text/javascript",
				web.JSVueComponentsPack(),
				codehighlight.JSComponentsPack(),
				JSComponentsPack(),
				web.JSComponentsPack()))

		b.mux.Handle("/index.css", b.builder.PacksHandler("text/css",
			codehighlight.CSSComponentsPack(),
			CSSComponentsPack(),
		))
	}
}

func (b *Builder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.mux == nil {
		b.Build()
	}
	b.mux.ServeHTTP(w, r)
	return
}

var arrowIcon = RawHTML(`
<svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 14 14" data-v-134594fd="" data-v-838665fe=""><path data-v-7abeccde="" d="m4.81347656 13.1269531c.22558594 0 .41015625-.0820312.56738282-.2324219l5.31835942-5.19531245c.1845703-.19140625.2802734-.38964844.2802734-.63574219 0-.23925781-.0888672-.45117187-.2802734-.62890625l-5.31835942-5.20214843c-.15722657-.15039063-.34179688-.23242188-.56738282-.23242188-.45800781 0-.81347656.35546875-.81347656.80664062 0 .21875.09570312.43066407.24609375.58789063l4.79199219 4.67578125-4.79199219 4.6621094c-.15722656.1640625-.24609375.3623047-.24609375.5878906 0 .4511719.35546875.8066406.81347656.8066406z"></path></svg>
`)
