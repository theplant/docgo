package main

import (
	"github.com/theplant/docgo"
	"github.com/theplant/docgo/x/docgo/template/docsrc"
)

func main() {
	docgo.New().
		Assets("/assets/", docsrc.Assets).
		Home(docsrc.Home).
		SitePrefix("/docgoPackageName/").
		BuildStaticSite("../docs")
}
