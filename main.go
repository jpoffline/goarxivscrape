// find_links_in_page.go
package main

import (
	"arxiv/arxivlib"
)

func main() {
	// Make HTTP request
	arts := arxivlib.ScrapeForArticles("https://arxiv.org/list/hep-th/pastweek?skip=0&show=50")
	arts.Print()
}
