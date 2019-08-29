package arxivlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func NewArticle(t string, aut []string, cat string) Article {
	return Article{Title: t, Authors: aut, Category: cat}
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

func (as *Articles) addArticle(art Article) {
	as.Arts = append(as.Arts, art)
}

func (as *Articles) attachCodes(c []ArticleMeta) {
	for i := range as.Arts {
		as.Arts[i].addMeta(c[i])
	}
}

// Len will return the number of articles.
func (as *Articles) Len() int {
	return len(as.Arts)
}

// Overview will print meta data about all the articles
// (count etc).
func (as *Articles) Overview() {
	fmt.Printf("Number of articles: %v", as.Len())
}

// Print will pretty print the articles info
func (as *Articles) Print() {
	for _, a := range as.Arts {
		a.print()
	}
}

func (as *Articles) ToJSON(filename string) {
	file, _ := json.MarshalIndent(as, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}
