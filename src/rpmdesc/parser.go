package local

import "github.com/PuerkitoBio/goquery"

// Parser is a common interface
type Parser interface {
	GetSearchUrl(name string, os string, arch string) string
	GetObjectsFromRpm(doc *goquery.Document) string
	GetLicenseFromRpm(doc *goquery.Document) string
	GetHomepage(doc *goquery.Document) string
	GetDescUrl(url string) string
	GetPrefixURL() string
}
