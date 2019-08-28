package arxivlib

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func NewArticle(t string, aut []string, cat string) *Article {
	return &Article{Title: t, Authors: aut, Category: cat}
}

func NewArticleMeta(c, p string) ArticleMeta {
	return ArticleMeta{Code: c, PdfLink: p}
}

func (a *Article) print() {
	fmt.Printf("%s %s\n", a.Meta.Code, a.Title)
	fmt.Printf("Category; %s\n", a.Category)
	fmt.Printf("%v authors\n", len(a.Authors))
}

func (a *Article) addMeta(m ArticleMeta) {
	a.Meta = m
}

func (as *Articles) addArticle(art *Article) {
	as.articles = append(as.articles, art)
}

func (as *Articles) attachCodes(c []ArticleMeta) {
	for i := range as.articles {
		as.articles[i].addMeta(c[i])
	}
}

// Len will return the number of articles.
func (as *Articles) Len() int {
	return len(as.articles)
}

// Overview will print meta data about all the articles
// (count etc).
func (as *Articles) Overview() {
	fmt.Printf("Number of articles: %q", as.Len())
}

// Print will pretty print the articles info
func (as *Articles) Print() {
	for _, a := range as.articles {
		a.print()
	}
}

func pullArxivCodes(document *goquery.Document) []ArticleMeta {
	var arxivs []ArticleMeta
	document.Find("dl").Find("dt").Each(func(index int, element *goquery.Selection) {
		eee := strings.Split(element.Text(), "  ")[1]
		arxivCode := strings.Split(eee, " ")[0]

		element.Find("a").Each(func(index int, ch *goquery.Selection) {
			a, _ := ch.Attr("title")
			switch a {
			case "Download PDF":
				pdfLink, _ := ch.Attr("href")
				arxivs = append(arxivs, NewArticleMeta(arxivCode, pdfLink))
				break
			}
		})
	})
	return arxivs
}

func pullArticles(document *goquery.Document) (Articles, error) {
	var ScrapedArticles Articles
	document.Find("dl").Find("dd").Each(func(index int, element *goquery.Selection) {
		element.Each(func(index int, element *goquery.Selection) {

			title := element.Find(".list-title.mathjax")
			category := element.Find(".list-subjects").Find(".primary-subject")
			var authors []string
			element.Find(".list-authors").Find("a").Each(func(index int, element *goquery.Selection) { authors = append(authors, element.Text()) })
			cleantitle := strings.TrimSpace(strings.Replace(title.Text(), "Title: ", "", -1))
			art := NewArticle(cleantitle, authors, category.Text())
			ScrapedArticles.addArticle(art)
		})
	})
	return ScrapedArticles, nil
}

// ScrapeForArticles will return the provided URL for
// Articles.
func ScrapeForArticles(url string) (Articles, error) {
	response, err := http.Get(url)
	if err != nil {
		return Articles{}, fmt.Errorf(err.Error())
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return Articles{}, fmt.Errorf("Error loading HTTP response body: %s ", err)
	}

	ScrapedArticles, err := pullArticles(document)
	if err != nil {
		return Articles{}, fmt.Errorf("Error scraping articles from body: %s", err)
	}
	arxivs := pullArxivCodes(document)

	ScrapedArticles.attachCodes(arxivs)
	return ScrapedArticles, nil
}
