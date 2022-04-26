package docsrc

import "github.com/theplant/docgo"

var DocTree = []interface{}{
	Home,
	&docgo.DocsGroup{
		Title: "Essentials",
		Docs: []*docgo.DocBuilder{
			GithubPagesIntegration,
			UseWithHtmlGo,
			MarkdownDifference,
		},
	},
	&docgo.DocsGroup{
		Title: "Samples",
		Docs: []*docgo.DocBuilder{
			HelloWorld,
		},
	},
}
