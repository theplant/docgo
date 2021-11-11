package main

import (
	"github.com/theplant/docgo"
	"github.com/theplant/docgo/docsrc"
)

func main() {
	docgo.New().
		Assets("/assets/", docsrc.Assets).
		Home(docsrc.Home).
		SitePrefix("/docgo/").
		BuildStaticSite("../docs")
}
