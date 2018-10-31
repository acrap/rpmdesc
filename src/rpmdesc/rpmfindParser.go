package local

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// RpmFind is a parser for rpmfind.net
type RpmFind struct{}

// GetSearchUrl create full search url
func (parser RpmFind) GetSearchUrl(rpmnName, os_arg, arch string) string {
	url := "https://rpmfind.net/linux/rpm2html/search.php?query=%s&submit=Search+...&system=%s&arch=%s"
	req := fmt.Sprintf(url, rpmnName, os_arg, arch)
	return req
}

// GetPrefixURL get site url
func (parser RpmFind) GetPrefixURL() string {
	return "https://rpmfind.net/"
}

// GetDescUrl get description for package
func (parser RpmFind) GetDescUrl(url string) string {
	result := ""
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return result
	}
	table := doc.Find("table")
	// Find the review items
	hrefs := table.Find("a")

	for _, element := range hrefs.Nodes {
		attr := element.Attr[0]
		if attr.Key == "href" {
			if strings.Contains(attr.Val, "-") {
				if !strings.HasSuffix(attr.Val, ".rpm") {
					return attr.Val
				}
			}
		}
	}
	return result
}

func getTextByCategory(doc *goquery.Document, category string) string {
	result := ""
	doc.Find("h3").Each(func(_ int, selection *goquery.Selection) {
		text := selection.Text()
		if strings.Contains(text, category) {
			result = selection.Next().Text()
			return
		}
	})
	return result
}

// GetLicenseFromRpm is a method to get license info
func (parser RpmFind) GetLicenseFromRpm(doc *goquery.Document) string {
	return getTextByCategory(doc, "License")
}

// GetObjectsFromRpm get list of files from rpm
func (parser RpmFind) GetObjectsFromRpm(doc *goquery.Document) string {
	return getTextByCategory(doc, "Files")
}

// GetHomepage get homepage of the project
func (parser RpmFind) GetHomepage(doc *goquery.Document) string {
	result := ""
	doc.Find("td").Each(func(_ int, selection *goquery.Selection) {
		text := selection.Text()
		if strings.HasPrefix(text, "Url") {
			result = text[strings.Index(text, "htt"):]

		}
	})
	return result
}
