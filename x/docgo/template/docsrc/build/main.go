package main

import (
	"github.com/theplant/docgo"
	"github.com/theplant/docgo/x/docgo/template/docsrc"
)

func main() {
	docgo.New().
		Assets("/assets/", docsrc.Assets).
		MainPageTitle("My Document").
		SitePrefix("/docgoPackageName/").
		DocTree(docsrc.DocTree...).
		BuildStaticSite("../docs")
}
