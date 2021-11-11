package main

import (
	"github.com/theplant/docgo"
	"github.com/theplant/docgo/docs"
)

func main() {
	docgo.New().
		Assets("/assets/", docs.Assets).
		Home(docs.Home).
		SitePrefix("/docgo/").
		BuildStaticSite("dist")
}
