package main

import (
	"fmt"
	"net/http"

	"github.com/theplant/docgo"
	"github.com/theplant/docgo/x/docgo/template/docsrc"
	"github.com/theplant/osenv"
)

var port = osenv.Get("PORT", "The port to serve on", "8800")

func main() {
	fmt.Println("Starting docs at :" + port)
	http.Handle("/", docgo.New().
		MainPageTitle("My Document").
		Assets("/assets/", docsrc.Assets).
		DocTree(docsrc.DocTree...).
		Build(),
	)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
