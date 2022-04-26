package main

import (
	"github.com/theplant/docgo"
	"github.com/theplant/docgo/docsrc"
)

func main() {
	docgo.New().
		Assets("/assets/", docsrc.Assets).
		MainPageTitle("docgo Document").
		SitePrefix("/docgo/").
		DocTree(docsrc.DocTree...).
		BuildStaticSite("../docs")
}
