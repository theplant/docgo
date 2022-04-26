package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/theplant/docgo"
	"github.com/theplant/docgo/x/docgo/template/docsrc"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8800"
	}
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
