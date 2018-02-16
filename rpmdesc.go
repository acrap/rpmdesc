package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
	"strings"
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {
	rpmName := ""
	arch :=""
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "name",
			Value: "",
			Usage: "name of rpm packet",
		},
		cli.StringFlag{
			Name: "arch",
			Value: "",
			Usage: "architecture of rpm packet",
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.String("name")) == 0 {
			rpmName = enterName()
		} else {
			rpmName = c.String("name")
		}
		arch = c.String("arch")
		return nil
	}

	app.Run(os.Args)

	res := getSearchUrl(rpmName, arch)
	descUrl := getDescUrl(res)
	if len(descUrl) == 0{
		log.Fatal("No rpm's found")
		return
	}
	fmt.Println("Found packet:", descUrl[strings.LastIndex(descUrl,"/")+1:
		strings.LastIndex(descUrl,".")])

	doc, err := goquery.NewDocument("https://rpmfind.net/" + descUrl)

	if err != nil {
		log.Fatal(err)
	}

	result := getObjectsFromRpm(doc)
	getHomepage(doc)

	fmt.Printf("%s\n", result)
}

func getSearchUrl(rpmnName, arch string) string {
	urlPart1 := "https://rpmfind.net/linux/rpm2html/search.php?query="
	urlPart2 := rpmnName + "&submit=Search+...&system=&arch=" + arch
	return urlPart1 + urlPart2
}

func enterName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rpm name: ")
	rpmName, _ := reader.ReadString('\n')
	rpmName = strings.Replace(rpmName, "\n","",1)
	fmt.Println(rpmName)
	return rpmName
}

func getDescUrl(url string) string {
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
		if attr.Key=="href" {
			if strings.Contains(attr.Val, "-"){
				if !strings.HasSuffix(attr.Val, ".rpm"){
					return attr.Val
				}
			}
		}
	}
	return result
}

func getObjectsFromRpm(doc *goquery.Document) string {
	result:=""
	doc.Find("h3").Each(func(_ int, selection *goquery.Selection) {
		text:=selection.Text()
		if strings.Contains(text,"Files"){
			result = selection.Next().Text()
			return
		}
	})

	return result
}

func getHomepage(doc *goquery.Document) {
	doc.Find("td").Each(func(_ int, selection *goquery.Selection) {
		text := selection.Text()
		if strings.HasPrefix(text, "Url"){
			fmt.Println(text[strings.Index(text, "htt"):],"")
		}
	})
}
