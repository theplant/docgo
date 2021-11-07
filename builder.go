package docgo

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"sort"
	"time"

	h "github.com/theplant/htmlgo"
)

type Builder struct {
	home        *DocBuilder
	header      h.HTMLComponent
	footer      h.HTMLComponent
	articleTree *DocNode
	mux         *http.ServeMux
	mounts      []*mount
}

func New() (r *Builder) {
	r = &Builder{}
	return
}

func (b *Builder) Home(v *DocBuilder) (r *Builder) {
	b.home = v
	return b
}

func (b *Builder) Header(v h.HTMLComponent) (r *Builder) {
	b.header = v
	return b
}

func (b *Builder) Footer(v h.HTMLComponent) (r *Builder) {
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

func (b *Builder) layout(body h.HTMLComponent) (r h.HTMLComponent) {
	return h.HTML(
		h.Head(
			h.Link("/index.css").Rel("stylesheet").Type("text/css"),
			// h.Script("").Src("/index.js"),
		),
		h.Body(
			b.header,
			body,
			b.footer,
		),
	)
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
			err := h.Fprint(w, b.layout(node.Doc), context.TODO())
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
		js1, _ := box.ReadFile("docjs/dist/vendor.js")
		js2, _ := box.ReadFile("docjs/dist/index.js")
		css, _ := box.ReadFile("docjs/dist/index.css")
		js := bytes.NewBuffer(js1)
		js.WriteString(";\n")
		js.Write(js2)

		b.mux.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/javascript")
			http.ServeContent(w, r, "", startTime, bytes.NewReader(js.Bytes()))
		})

		b.mux.HandleFunc("/index.css", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/css")
			http.ServeContent(w, r, "", startTime, bytes.NewReader(css))
		})
	}
}

func (b *Builder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.mux == nil {
		b.Build()
	}
	b.mux.ServeHTTP(w, r)
	return
}
