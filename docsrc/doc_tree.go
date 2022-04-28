package docsrc

import "github.com/theplant/docgo"

var DocTree = []interface{}{
	Home,
	UseWithHtmlGo,
	GithubPagesIntegration,
	MarkdownDifference,
	&docgo.DocsGroup{
		Title: "Samples",
		Docs: []*docgo.DocBuilder{
			HelloWorld,
		},
	},
}
