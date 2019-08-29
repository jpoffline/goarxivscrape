package arxivlib

type Articles struct {
	Arts []Article `json:"articles"`
}

type Article struct {
	Title    string      `json:"title"`
	Authors  []string    `json:"authors"`
	Category string      `json:"category"`
	Meta     ArticleMeta `json:"meta"`
}

type ArticleMeta struct {
	Code    string `json:"code"`
	PdfLink string `json:"pdf"`
}
