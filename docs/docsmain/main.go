package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/theplant/docgo"
	"github.com/theplant/docgo/docs"
)

// @snippet_begin(BootUpDevSample)
func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8800"
	}
	fmt.Println("Starting docs at :" + port)
	http.Handle("/", docgo.New().
		Home(docs.Home).
		Build(),
	)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// @snippet_end
