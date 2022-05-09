package docgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
	. "github.com/theplant/htmlgo"
	"golang.org/x/net/html"
)

type DocBuilder struct {
	title        string
	slug         string
	abstractText string
	children     []HTMLComponent
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

func (b *DocBuilder) GetPageURL() (r string) {
	slug := b.slug
	if slug == "" {
		slug = strcase.ToKebab(b.title)
	}
	u := path.Join("/", slug)
	if u == "/" {
		return "index.html"
	}
	return strings.TrimLeft(fmt.Sprintf("%s.html", u), "/")
}

func (b *DocBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	return Div(
		Div(
			Div(
				Button("").Children(
					Div(
						menuIcon,
					).Class("w-4 h-4 fill-current text-gray-300"),
				).Class("w-12 h-12 p-4").
					Attr("@click", "vars.hideAside = !vars.hideAside"),
			).Class("flex flex-row"),
			Div(
				H1(b.title).Class("mb-8"),
				If(len(b.abstractText) > 0,
					Div(
						Text(b.abstractText),
					).Class("mb-8 text-xl font-normal"),
				),
				Div(
					b.children...,
				).Class("border-t"),
			).Class("px-16 pb-12 pt-4 overflow-auto").
				Id("docMainBox"),
		).Class("flex flex-grow flex-col w-2/3"),
		Div(
			Div(
				Text("On This Page"),
				RawHTML("<toc></toc>"),
			).Class("sticky top-4 w-52"),
		).Class("font-medium text-base hidden xl:block text-gray-600 pt-4"),
	).
		Class("flex flex-row w-full").
		Id("docContentBox").
		MarshalHTML(ctx)
}

func (b *DocBuilder) plainText() string {
	hb, err := Div(
		Div(Text(b.abstractText)),
		Div(b.children...),
	).MarshalHTML(context.Background())
	if err != nil {
		panic(err)
	}

	return html2text(hb)
}

var docIcon = RawHTML(`
<svg aria-hidden="true" class=""
viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
<path
d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"></path>
</svg>
`)

func html2text(in []byte) string {
	r := &bytes.Buffer{}

	tokenizer := html.NewTokenizer(bytes.NewReader(in))
	startToken := tokenizer.Token()

	for tt := tokenizer.Next(); tt != html.ErrorToken; tt = tokenizer.Next() {
		switch tt {
		case html.StartTagToken:
			startToken = tokenizer.Token()
			if startToken.Data == "highlightjs" {
				for _, v := range startToken.Attr {
					if v.Key == ":code" {
						code := v.Val
						if code != "" {
							if err := json.Unmarshal([]byte(code), &code); err != nil {
								panic(err)
							}
							code = strings.NewReplacer("\n", " ", "\t", " ").Replace(code)
							r.WriteString(code)
							r.WriteString(" ")
						}
					}
				}
			}
		case html.TextToken:
			if startToken.Data == "script" || startToken.Data == "style" {
				continue
			}
			txt := strings.TrimSpace(string(tokenizer.Text()))
			if txt != "" {
				r.WriteString(txt)
				r.WriteString(" ")
			}
		}
	}

	return r.String()
}

var menuIcon = RawHTML(`
<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="16px" height="16px" viewBox="0 0 16 16" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 2 12 L 2 11 L 14 11 L 14 12 Z M 2 8.5 L 2 7.5 L 14 7.5 L 14 8.5 Z M 2 5 L 2 4 L 14 4 L 14 5 Z M 2 5 "/>
</g>
</svg>
`)
