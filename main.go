// find_links_in_page.go
package main

import (
	"arxiv/arxivlib"
	"flag"
)

func main() {
	// Make HTTP request

	arxivType := flag.String("arxiv", "hep-th", "which arxiv")
	flag.Parse()

	arts, _ := arxivlib.ScrapeForArticles("https://arxiv.org/list/" + *arxivType + "/pastweek?skip=0&show=50")
	arts.Print()
	arts.Overview()
	arts.ToJSON("arts.json")
}
