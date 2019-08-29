package arxivlib

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func pullArxivCodes(document *goquery.Document) []ArticleMeta {
	var arxivs []ArticleMeta
	document.Find("dl").Find("dt").Each(func(index int, element *goquery.Selection) {
		eee := strings.Split(element.Text(), "  ")[1]
		arxivCode := strings.Split(eee, " ")[0]
		ac := strings.Split(arxivCode, ":")[1]
		element.Find("a").Each(func(index int, ch *goquery.Selection) {
			a, _ := ch.Attr("title")
			switch a {
			case "Download PDF":
				pdfLink, _ := ch.Attr("href")

				arxivs = append(arxivs, NewArticleMeta(ac, pdfLink))
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
