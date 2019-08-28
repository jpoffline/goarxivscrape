package arxivlib

type Articles struct {
	articles []*Article
}

type Article struct {
	Title    string
	Authors  []string
	Category string
	Meta     ArticleMeta
}

type ArticleMeta struct {
	Code    string
	PdfLink string
}
