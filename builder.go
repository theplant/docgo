package docgo

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/goplaid/web"
	"github.com/theplant/docgo/ch"
	. "github.com/theplant/htmlgo"
)

type Builder struct {
	docTree       []interface{}
	mux           *http.ServeMux
	mounts        []*mount
	searchIndexes []*searchIndex
	builder       *web.Builder
	assets        embed.FS
	assetsPrefix  string
	sitePrefix    string
	mainPageTitle string
}

func New() (r *Builder) {
	r = &Builder{
		builder:      web.New(),
		sitePrefix:   "/",
		assetsPrefix: "/assets/",
	}
	return
}

func (b *Builder) Assets(prefix string, v embed.FS) (r *Builder) {
	b.assets = v
	b.assetsPrefix = prefix
	return b
}

func (b *Builder) SitePrefix(v string) (r *Builder) {
	b.sitePrefix = v
	return b
}

func (b *Builder) MainPageTitle(v string) (r *Builder) {
	b.mainPageTitle = v
	return b
}

type DocsGroup struct {
	Title string
	Docs  []*DocBuilder
}

// *DocBuilder, *DocsGroup
func (b *Builder) DocTree(vs ...interface{}) (r *Builder) {
	b.docTree = vs
	return b
}

type searchIndex struct {
	URL   string
	Title string
	Body  string
}

func (b *Builder) Build() (r *Builder) {
	for _, di := range b.docTree {
		switch v := di.(type) {
		case *DocBuilder:
			b.addToMounts(v)
			b.addToSearchIndexes(v)
		case *DocsGroup:
			for _, sd := range v.Docs {
				b.addToMounts(sd)
				b.addToSearchIndexes(sd)
			}
		default:
			panic("only support *DocBuilder,*DocsGroup")
		}
	}
	sort.Sort(uriSorter(b.mounts))
	b.initMuxWithStatic()
	for _, m := range b.mounts {
		fmt.Println("Mounting", m.path)
		b.mux.HandleFunc(m.path, m.f)
	}

	assetsFs := http.FileServer(http.FS(b.assets))
	b.mux.Handle(b.assetsPrefix, assetsFs)

	b.mux.HandleFunc("/search_indexes.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b.searchIndexes)
	})

	return b
}

func (b *Builder) BuildStaticSite(dir string) {
	dir = strings.NewReplacer(" ", "_").Replace(dir)
	if len(dir) == 0 || dir == "/" {
		dir = "dist"
	}

	handler := b.Build()
	err1 := os.RemoveAll(dir)
	if err1 != nil {
		fmt.Println("removing ", dir, err1)
	}

	var paths = []string{
		"/index.css",
		"/index.js",
		"/search_indexes.json",
	}

	for _, m := range b.mounts {
		paths = append(paths, m.path)
	}

	fs.WalkDir(b.assets, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		paths = append(paths, filepath.Join("/", path))
		return nil
	})

	for _, p := range paths {
		r := httptest.NewRequest("GET", p, nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		func() {
			path := filepath.Join(dir, p)
			fmt.Println("Generating ", path)
			err := os.MkdirAll(filepath.Dir(path), 0755)
			if err != nil {
				panic(err)
			}
			var f *os.File
			f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			_, err = io.Copy(f, w.Body)
			if err != nil {
				panic(err)
			}
		}()

	}
}

func (b *Builder) layout(body *DocBuilder) (r HTMLComponent) {
	pageTitle := body.title
	for _, di := range b.docTree {
		switch v := di.(type) {
		case *DocsGroup:
			for _, sd := range v.Docs {
				if sd == body {
					pageTitle = fmt.Sprintf("%s - %s", v.Title, pageTitle)
					break
				}
			}
		}
	}
	if b.mainPageTitle != "" {
		pageTitle = fmt.Sprintf("%s - %s", pageTitle, b.mainPageTitle)
	}

	return HTML(
		Head(
			Title(pageTitle),
			Meta().Name("description").Content(body.abstractText),
			RawHTML(`<meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">`),
			If(len(b.sitePrefix) > 0, Base().Href(b.sitePrefix)),
			Link("index.css").Rel("stylesheet").Type("text/css"),
			Script("").Attr("defer", true).Src("index.js"),
		),
		Body(
			Div(
				Div(
					Div(
						Div(
							b.aside(body),
							Main(
								body,
								RawHTML("<search-result></search-result>"),
							).Class("flex flex-col w-full bg-white overflow-x-hidden overflow-y-auto"),
						).Class("flex h-full"),
					).Class("flex-1 flex flex-col overflow-hidden"),
				).Class("flex h-screen").
					Attr(web.InitContextVars, `{hideAside: false}`),
			).Id("app").
				Attr("v-cloak", true),
		),
	)
}

func (b *Builder) aside(doc *DocBuilder) (r HTMLComponent) {
	content := Ul().Class("px-0 mx-0 text-base font-normal list-none text-gray-700")
	for _, di := range b.docTree {
		switch v := di.(type) {
		case *DocBuilder:
			link := A().Href(v.GetPageURL()).Text(v.title).Class("inline-block px-4 truncate break-words w-64")
			if v == doc {
				link.Class("text-blue-500")
			} else {
				link.Class("text-gray-700")
			}
			content.AppendChildren(
				Li(link),
			)
		case *DocsGroup:
			content.AppendChildren(
				Li().Text(v.Title).Class("cursor-default px-4 truncate break-words w-64"),
			)
			for _, sd := range v.Docs {
				link := A().Href(sd.GetPageURL()).Text(sd.title).Class("inline-block pl-10 pr-4 truncate break-words w-64")
				if sd == doc {
					link.Class("text-blue-500")
				} else {
					link.Class("text-gray-700")
				}
				content.AppendChildren(
					Li(link),
				)
			}
		}
	}

	return Aside(
		Div(
			RawHTML("<search></search>"),
		).Class("h-12"),
		content,
	).Class("flex flex-col w-80 h-full bg-gray-50 border-r border-gray-200").
		Attr("v-show", "!vars.hideAside")
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

func (b *Builder) addToMounts(d *DocBuilder) {
	b.mounts = append(b.mounts,
		&mount{filepath.Join("/", d.GetPageURL()), func(w http.ResponseWriter, r *http.Request) {
			err := Fprint(w, b.layout(d), context.TODO())
			if err != nil {
				panic(err)
			}
		}})
}

func (b *Builder) addToSearchIndexes(d *DocBuilder) {
	b.searchIndexes = append(b.searchIndexes, &searchIndex{
		URL:   d.GetPageURL(),
		Title: d.title,
		Body:  d.plainText(),
	})
}

func (b *Builder) initMuxWithStatic() {
	if b.mux == nil {
		b.mux = http.NewServeMux()

		b.mux.Handle("/index.js",
			b.builder.PacksHandler("text/javascript",
				web.JSVueComponentsPack(),
				ch.JSComponentsPack(),
				JSComponentsPack(),
				web.JSComponentsPack()))

		b.mux.Handle("/index.css", b.builder.PacksHandler("text/css",
			ch.CSSComponentsPack(),
			CSSComponentsPack(),
		))
	}
}

func (b *Builder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.mux == nil {
		b.Build()
	}

	if r.URL.String() == "/" {
		http.Redirect(w, r, "/index.html", http.StatusPermanentRedirect)
	}

	b.mux.ServeHTTP(w, r)
	return
}

var arrowIcon = RawHTML(`
<svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 14 14" data-v-134594fd="" data-v-838665fe=""><path data-v-7abeccde="" d="m4.81347656 13.1269531c.22558594 0 .41015625-.0820312.56738282-.2324219l5.31835942-5.19531245c.1845703-.19140625.2802734-.38964844.2802734-.63574219 0-.23925781-.0888672-.45117187-.2802734-.62890625l-5.31835942-5.20214843c-.15722657-.15039063-.34179688-.23242188-.56738282-.23242188-.45800781 0-.81347656.35546875-.81347656.80664062 0 .21875.09570312.43066407.24609375.58789063l4.79199219 4.67578125-4.79199219 4.6621094c-.15722656.1640625-.24609375.3623047-.24609375.5878906 0 .4511719.35546875.8066406.81347656.8066406z"></path></svg>
`)
